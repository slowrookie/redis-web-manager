package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/slowrookie/redis-web-manager/api"
)

// App struct
type App struct {
	ctx   context.Context
	About map[string]string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (a *App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	fmt.Println("")
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// export methods

// Abount .
func (a *App) AboutInfo() map[string]string {
	return a.About
}

// Connections .
func (a *App) Connections() ([]api.Connection, error) {
	return api.Connections()
}

func (a *App) TestConnection(con api.Connection) error {
	return con.Test()
}

func (a *App) EditConnection(con api.Connection) error {
	return con.New()
}

func (a *App) DeleteConnection(id string) error {
	connection, err := api.FindConnectionByID(id)
	if err != nil {
		return err
	}
	return connection.Delete()
}

func (a *App) OpenConnection(id string) error {
	connection, err := api.FindConnectionByID(id)
	if err != nil {
		return err
	}
	return connection.Open()
}

func (a *App) DisConnection(id string) error {
	connection, err := api.FindConnectionByID(id)
	if err != nil {
		return err
	}
	return connection.Disconnection()
}

func (a *App) CopyConnection(con api.Connection) error {
	return con.New()
}

func (a *App) CommandConnection(cmd api.Command) ([]interface{}, error) {
	connection, err := api.FindConnectionByID(cmd.ID)
	if err != nil {
		return nil, err
	}
	return connection.Command(cmd)
}

// Convert .
func (a *App) ConvertLength(cv api.Convert) string {
	return strconv.Itoa(cv.Length())
}

func (a *App) ConvertToHex(cv api.Convert) string {
	return cv.ToHex()
}

func (a *App) ConvertToJson(cv api.Convert) string {
	return cv.ToJson()
}

func (a *App) ConvertToBinary(cv api.Convert) string {
	return cv.ToBinary()
}

func (a *App) ConvertBase64ToText(cv api.Convert) string {
	return cv.Base64ToText()
}

func (a *App) ConvertBase64ToJson(cv api.Convert) string {
	return cv.Base64ToJson()
}

// Config .
func (a *App) Config() (api.Config, error) {
	return api.DefaultConfig.Get()
}

func (a *App) SetConfig(config api.Config) error {
	return config.Set()
}
