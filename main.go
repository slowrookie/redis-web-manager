package main

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/pkg/browser"
	"github.com/slowrookie/redis-web-manager/api"
	"go.lsp.dev/jsonrpc2"
	"golang.org/x/net/websocket"
)

var (
	version = "Development"
	commit  = "Development"
	date    = "Now"
	builtBy = "Development"
	MODE    = "Debug"
)

//go:embed web/build/*
var webFS embed.FS

func main() {
	// files
	os.MkdirAll(api.ROOT_PATH, os.ModePerm)

	api.InitializeDB(path.Join("./", api.ROOT_PATH, "rwm.db?cache=shared&mode=rwc&_journal_mode=WAL"))

	// static files
	buildFiles, err := fs.Sub(webFS, "web/build")
	if err != nil {
		panic(nil)
	}

	// static files
	http.Handle("/", http.FileServer(http.FS(buildFiles)))

	// server
	http.Handle("/ws", websocket.Handler(func(ws *websocket.Conn) {
		ctx := ws.Request().Context()
		ServerConn(ctx, ws, jsonrpc2.HandlerServer(func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
			dec := json.NewDecoder(bytes.NewReader(req.Params()))
			switch req.Method() {
			case "About":
				var about = make(map[string]string)
				about["version"] = version
				about["commit"] = commit
				about["date"] = date
				about["builtBy"] = builtBy
				about["environment"] = MODE
				return reply(ctx, about, nil)
			// Config
			case "Config":
				conf, err := api.DefaultConfig.Get()
				return reply(ctx, conf, err)
			case "Config.Set":
				var config api.Config
				if err := dec.Decode(&config); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				err := config.Set()
				return reply(ctx, nil, err)
			case "Config.CheckPort":
				var port int
				if err := dec.Decode(&port); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				err := api.DefaultConfig.CheckPort(fmt.Sprintf("%d", port))
				return reply(ctx, nil, err)
			// Connections
			case "Connections":
				connections, err := api.Connections()
				return reply(ctx, connections, err)
			case "Connection.Test":
				var connection api.Connection
				if err := dec.Decode(&connection); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				err := connection.Test()
				return reply(ctx, nil, err)
			case "Connection.Edit":
				var connection api.Connection
				if err := dec.Decode(&connection); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				err := connection.New()
				return reply(ctx, nil, err)
			case "Connection.Delete":
				var id string
				if err := dec.Decode(&id); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				connection, err := api.FindConnectionByID(id)
				if err != nil {
					return reply(ctx, nil, err)
				}
				err = connection.Delete()
				return reply(ctx, nil, err)
			case "Connection.Open":
				var id string
				if err := dec.Decode(&id); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				connection, err := api.FindConnectionByID(id)
				if err != nil {
					return reply(ctx, nil, err)
				}
				err = connection.Open()
				return reply(ctx, nil, err)
			case "Connection.Disconnection":
				var id string
				if err := dec.Decode(&id); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				connection, err := api.FindConnectionByID(id)
				if err != nil {
					return reply(ctx, nil, err)
				}
				err = connection.Disconnection()
				return reply(ctx, nil, err)
			case "Connection.Copy":
				connection := &api.Connection{}
				if err := dec.Decode(&connection); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				err := connection.New()
				return reply(ctx, nil, err)
			case "Connection.Command":
				var cmd api.Command
				if err := dec.Decode(&cmd); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				connection, err := api.FindConnectionByID(cmd.ID)
				if err != nil {
					return reply(ctx, nil, err)
				}
				ret, err := connection.Command(cmd)
				return reply(ctx, ret, err)
			case "Convert.Length":
				var convert api.Convert
				if err := dec.Decode(&convert); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				return reply(ctx, strconv.Itoa(convert.Length()), nil)
			case "Convert.ToHex":
				var convert api.Convert
				if err := dec.Decode(&convert); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				return reply(ctx, convert.ToHex(), nil)
			case "Convert.ToJson":
				var convert api.Convert
				if err := dec.Decode(&convert); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				return reply(ctx, convert.ToJson(), nil)
			case "Convert.ToBinary":
				var convert api.Convert
				if err := dec.Decode(&convert); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				return reply(ctx, convert.ToBinary(), nil)
			case "Convert.Base64ToText":
				var convert api.Convert
				if err := dec.Decode(&convert); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				return reply(ctx, convert.Base64ToText(), nil)
			case "Convert.Base64ToJson":
				var convert api.Convert
				if err := dec.Decode(&convert); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				return reply(ctx, convert.Base64ToJson(), nil)
			default:
				return jsonrpc2.MethodNotFoundHandler(ctx, reply, req)
			}
		}))
	}))

	// config
	conf, err := api.DefaultConfig.Get()
	if err != nil {
		panic(nil)
	}

	// listen and serve on 0.0.0.0:63790 (for windows "localhost:63790")
	// r.Run()
	port := strconv.FormatUint(uint64(conf.Port), 10)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatal(err)
	}
	// 服务启动之后，打开系统浏览器
	if MODE != "Debug" {
		_ = browser.OpenURL(fmt.Sprintf("http://127.0.0.1:%s", port))
	}
	log.Fatal(http.Serve(listen, nil))
}
