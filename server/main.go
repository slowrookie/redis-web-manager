package main

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"path"

	"github.com/pkg/browser"
	"github.com/slowrookie/redis-web-manager/api"
	_ "github.com/slowrookie/redis-web-manager/api"
	"github.com/slowrookie/redis-web-manager/api/parser"
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

//go:embed web/build
var assets embed.FS

func main() {
	// log
	f, err := os.OpenFile(path.Join(api.APP_ROOT, "rwm.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, f))
	log.Println(fmt.Sprintf("Root Path: %s", api.APP_ROOT))

	// static files
	buildFiles, err := fs.Sub(assets, "web/build")
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
			case "AboutInfo":
				var about = make(map[string]string)
				about["version"] = version
				about["commit"] = commit
				about["date"] = date
				about["builtBy"] = builtBy
				about["environment"] = MODE
				return reply(ctx, about, nil)
			// Connections
			case "Connections":
				err := api.LoadConnections()
				return reply(ctx, api.Connections, err)
			case "TestConnection":
				var connection api.Connection
				if err := dec.Decode(&connection); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				err := connection.Test()
				return reply(ctx, nil, err)
			case "EditConnection":
				var connection api.Connection
				if err := dec.Decode(&connection); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				err := connection.New()
				return reply(ctx, nil, err)
			case "DeleteConnection":
				var id string
				if err := dec.Decode(&id); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				connection, err := api.GetConnection(id)
				if err != nil {
					return reply(ctx, nil, err)
				}
				err = connection.Delete()
				return reply(ctx, nil, err)
			case "OpenConnection":
				var id string
				if err := dec.Decode(&id); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				connection, err := api.GetConnection(id)
				if err != nil {
					return reply(ctx, nil, err)
				}
				err = connection.Open()
				return reply(ctx, nil, err)
			case "DisConnection":
				var id string
				if err := dec.Decode(&id); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				connection, err := api.GetConnection(id)
				if err != nil {
					return reply(ctx, nil, err)
				}
				err = connection.Disconnection()
				return reply(ctx, nil, err)
			case "NewConnection":
				connection := &api.Connection{}
				if err := dec.Decode(&connection); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				err := connection.New()
				return reply(ctx, nil, err)
			case "CommandConnection":
				var cmd api.Command
				if err := dec.Decode(&cmd); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				connection, err := api.GetConnection(cmd.ID)
				if err != nil {
					return reply(ctx, nil, err)
				}
				ret, err := connection.Command(cmd)
				return reply(ctx, ret, err)
			case "ExecutionScript":
				var lua api.Lua
				if err := dec.Decode(&lua); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				connection, err := api.GetConnection(lua.ConnectionID)
				if err != nil {
					return reply(ctx, nil, err)
				}
				err = connection.Scripting(&lua)
				return reply(ctx, lua, err)
			case "Suggestions":
				var command string
				if err := dec.Decode(&command); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				ret := parser.Suggestions(command)
				return reply(ctx, ret, err)
			// Config
			case "Config":
				if err := api.DefaultConfig.Get(); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				return reply(ctx, *api.DefaultConfig, err)
			case "SetConfig":
				var config api.Config
				if err := dec.Decode(&config); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				err := config.Set()
				return reply(ctx, nil, err)
			// Lua
			case "NewLua":
				lua := &api.Lua{}
				if err := dec.Decode(&lua); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				err := lua.New()
				return reply(ctx, nil, err)
			case "EditLua":
				lua := &api.Lua{}
				if err := dec.Decode(&lua); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				err := lua.Edit()
				return reply(ctx, nil, err)
			case "DeleteLua":
				lua := &api.Lua{}
				if err := dec.Decode(&lua); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				err := lua.Delete()
				return reply(ctx, nil, err)
			case "LoadLuas":
				var connectionId string
				if err := dec.Decode(&connectionId); err != nil {
					return jsonrpc2.ErrInvalidRequest
				}
				luas := make([]api.Lua, 0)
				if err := api.LoadLuas(&luas); nil != err {
					return reply(ctx, nil, err)
				}
				var _luas []api.Lua
				for _, v := range luas {
					if v.ConnectionID == connectionId {
						_luas = append(_luas, v)
					}
				}
				return reply(ctx, _luas, nil)
			default:
				return jsonrpc2.MethodNotFoundHandler(ctx, reply, req)
			}
		}))
	}))

	// port := strconv.FormatUint(uint64(conf.Port), 10)
	port := "63790"
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
