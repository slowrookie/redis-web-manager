package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/browser"
	"github.com/slowrookie/redis-web-manager/api"
)

var (
	version  = "Development"
	commit   = "Development"
	date     = "Now"
	builtBy  = "Development"
	GIN_MODE = gin.DebugMode
)

//go:embed web/build/*
var webFS embed.FS

func main() {
	gin.SetMode(GIN_MODE)

	// log
	os.MkdirAll(api.ROOT_PATH, os.ModePerm)
	logfile, _ := os.Create(api.LogFilePath)
	gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)

	r := gin.Default()

	// static files
	buildFiles, err := fs.Sub(webFS, "web/build")
	if err != nil {
		panic(nil)
	}

	r.GET("/", func(c *gin.Context) {
		c.FileFromFS("/", http.FS(buildFiles))
	})

	r.GET("/manifest.json", func(c *gin.Context) {
		c.FileFromFS("/manifest.json", http.FS(buildFiles))
	})

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.FileFromFS("/favicon.ico", http.FS(buildFiles))
	})

	r.GET("/logo192.png", func(c *gin.Context) {
		c.FileFromFS("/logo192.png", http.FS(buildFiles))
	})

	r.GET("/static/*filepath", func(c *gin.Context) {
		buildFiles, err := fs.Sub(webFS, "web/build/static")
		if err != nil {
			panic(nil)
		}
		c.FileFromFS(c.Param("filepath"), http.FS(buildFiles))
	})

	r.GET("/about", func(c *gin.Context) {
		var about = make(map[string]string)
		about["version"] = version
		about["commit"] = commit
		about["date"] = date
		about["builtBy"] = builtBy
		about["environment"] = GIN_MODE
		c.JSON(http.StatusOK, about)
	})

	// groups
	configGroup := r.Group("/config")
	{
		configGroup.GET("/", func(c *gin.Context) {
			conf, err := api.DefaultConfig.Get()
			if nil != err {
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
				return
			}
			c.JSON(http.StatusOK, conf)
		})
		configGroup.POST("/", func(c *gin.Context) {
			c.Bind(api.DefaultConfig)
			if err := api.DefaultConfig.Set(); nil != err {
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
				return
			}
			c.JSON(http.StatusOK, api.DefaultConfig)
		})
	}
	// connections
	// init connections
	api.LoadConnections()
	r.GET("/connections", func(c *gin.Context) {
		connections, err := api.Connections()
		if nil != err {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
			return
		}
		c.JSON(http.StatusOK, connections)
	})
	connectioniGroup := r.Group("/connection")
	{
		connectioniGroup.POST("/test", func(c *gin.Context) {
			connection := &api.Connection{}
			c.Bind(connection)
			err := connection.Test()
			if nil != err {
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
				return
			}
		})
		connectioniGroup.POST("/", func(c *gin.Context) {
			connection := &api.Connection{}
			c.Bind(connection)
			ret, err := connection.New()
			if nil != err {
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
				return
			}
			c.JSON(http.StatusOK, ret)
		})
		connectioniGroup.DELETE("/:id", func(c *gin.Context) {
			if err := api.FindConnectionByID(c.Param("id")).Delete(); nil != err {
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
				return
			}
		})
		connectioniGroup.POST("/:id/open", func(c *gin.Context) {
			err := api.FindConnectionByID(c.Param("id")).Open()
			if nil != err {
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
				return
			}
		})
		connectioniGroup.POST("/:id/disconnection", func(c *gin.Context) {
			if err := api.FindConnectionByID(c.Param("id")).Disconnection(); nil != err {
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
				return
			}
		})
		connectioniGroup.POST("/copy", func(c *gin.Context) {
			connection := &api.Connection{}
			c.Bind(connection)
			_connection, err := connection.Copy()
			if nil != err {
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
				return
			}
			c.JSON(http.StatusOK, _connection)
		})
		connectioniGroup.POST("/command", func(c *gin.Context) {
			cmd := &api.Command{}
			c.Bind(cmd)
			ret, err := api.FindConnectionByID(cmd.ID).Command(*cmd)
			if nil != err {
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
				return
			}
			c.JSON(http.StatusOK, ret)
		})
	}
	// convert
	convertGroup := r.Group("/convert")
	{
		convertGroup.POST("/length", func(c *gin.Context) {
			convert := &api.Convert{}
			c.Bind(convert)
			c.String(http.StatusOK, strconv.Itoa(convert.Length()))
		})
		convertGroup.POST("/toHex", func(c *gin.Context) {
			convert := &api.Convert{}
			c.Bind(convert)
			c.String(http.StatusOK, convert.ToHex())
		})
		convertGroup.POST("/toJson", func(c *gin.Context) {
			convert := &api.Convert{}
			c.Bind(convert)
			c.String(http.StatusOK, convert.ToJson())
		})
		convertGroup.POST("/toBinary", func(c *gin.Context) {
			convert := &api.Convert{}
			c.Bind(convert)
			c.String(http.StatusOK, convert.ToBinary())
		})
		convertGroup.POST("/base64ToText", func(c *gin.Context) {
			convert := &api.Convert{}
			c.Bind(convert)
			c.String(http.StatusOK, convert.Base64ToText())
		})
		convertGroup.POST("/base64ToJson", func(c *gin.Context) {
			convert := &api.Convert{}
			c.Bind(convert)
			c.String(http.StatusOK, convert.Base64ToJson())
		})
	}

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// r.Run()
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
	// 服务启动之后，打开系统浏览器
	if GIN_MODE != gin.DebugMode {
		_ = browser.OpenURL("http://127.0.0.1:8080")
	}
	log.Fatal(http.Serve(listen, r))

}
