package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/slowrookie/redis-web-manager/api"
)

func main() {
	r := gin.Default()
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
	r.GET("/connections", func(c *gin.Context) {
		connections, err := api.DefaultConnection.Connections()
		if nil != err {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
			return
		}
		c.JSON(http.StatusOK, connections)
	})
	connectioniGroup := r.Group("/connection")
	{
		connectioniGroup.POST("/test", func(c *gin.Context) {
			c.Bind(api.DefaultConnection)
			err := api.DefaultConnection.Test()
			if nil != err {
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
				return
			}
		})
		connectioniGroup.POST("/", func(c *gin.Context) {
			c.Bind(api.DefaultConnection)
			ret, err := api.DefaultConnection.New()
			if nil != err {
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
				return
			}
			c.JSON(http.StatusOK, ret)
		})
		connectioniGroup.DELETE("/:id", func(c *gin.Context) {
			if err := api.DefaultConnection.Delete(c.Param("id")); nil != err {
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
				return
			}
		})
		connectioniGroup.POST("/:id/open", func(c *gin.Context) {
			ret, err := api.DefaultConnection.Open(c.Param("id"))
			if nil != err {
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
				return
			}
			c.JSON(http.StatusOK, ret)
		})
		connectioniGroup.POST("/:id/disconnection", func(c *gin.Context) {
			if err := api.DefaultConnection.Disconnection(c.Param("id")); nil != err {
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
				return
			}
		})
		connectioniGroup.POST("/copy", func(c *gin.Context) {
			c.Bind(api.DefaultConnection)
			connection, err := api.DefaultConnection.Copy()
			if nil != err {
				c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
				return
			}
			c.JSON(http.StatusOK, connection)
		})
		connectioniGroup.POST("/command", func(c *gin.Context) {
			cmd := &api.Command{}
			c.Bind(cmd)
			ret, err := api.DefaultConnection.Command(*cmd)
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
			data := c.PostForm("data")
			c.String(http.StatusOK, strconv.Itoa(api.DefaultConvert.Length(data)))
		})
		convertGroup.POST("/toHex", func(c *gin.Context) {
			data := c.PostForm("data")
			c.String(http.StatusOK, api.DefaultConvert.ToHex(data))
		})
		convertGroup.POST("/toHex", func(c *gin.Context) {
			data := c.PostForm("data")
			c.String(http.StatusOK, api.DefaultConvert.ToJson(data))
		})
		convertGroup.POST("/toBinary", func(c *gin.Context) {
			data := c.PostForm("data")
			c.String(http.StatusOK, api.DefaultConvert.ToBinary(data))
		})
		convertGroup.POST("/base64ToText", func(c *gin.Context) {
			data := c.PostForm("data")
			c.String(http.StatusOK, api.DefaultConvert.Base64ToText(data))
		})
		convertGroup.POST("/base64ToJson", func(c *gin.Context) {
			data := c.PostForm("data")
			c.String(http.StatusOK, api.DefaultConvert.Base64ToJson(data))
		})
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
