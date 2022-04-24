package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/slowrookie/redis-web-manager/api"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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

// export methods

// Abount .
func (a *App) AboutInfo() map[string]string {
	return a.About
}

// Connections .
func (a *App) Connections() []*api.Connection {
	api.LoadConnections()
	return api.Connections
}

func (a *App) TestConnection(con api.Connection) error {
	return con.Test()
}

func (a *App) EditConnection(con api.Connection) error {
	return con.Edit()
}

func (a *App) DeleteConnection(id string) error {
	connection, err := api.GetConnection(id)
	if err != nil {
		return err
	}
	return connection.Delete()
}

func (a *App) OpenConnection(id string) error {
	connection, err := api.GetConnection(id)
	if err != nil {
		return err
	}
	return connection.Open()
}

func (a *App) DisConnection(id string) error {
	connection, err := api.GetConnection(id)
	if err != nil {
		return err
	}
	return connection.Disconnection()
}

func (a *App) NewConnection(con api.Connection) error {
	return con.New()
}

func (a *App) CommandConnection(cmd api.Command) ([]interface{}, error) {
	connection, err := api.GetConnection(cmd.ID)
	if err != nil {
		return nil, err
	}
	return connection.Command(cmd)
}

func (a *App) ExecutionScript(lua *api.Lua) (*api.Lua, error) {
	connection, err := api.GetConnection(lua.ConnectionID)
	if err != nil {
		return nil, err
	}
	err = connection.Scripting(lua)
	return lua, nil
}

// Config .
func (a *App) Config() (api.Config, error) {
	if err := api.DefaultConfig.Get(); err != nil {
		return *api.DefaultConfig, err
	}
	return *api.DefaultConfig, nil
}

func (a *App) SetConfig(config api.Config) error {
	return config.Set()
}

// Lua
func (a *App) NewLua(lua api.Lua) error {
	return lua.New()
}

func (a *App) EditLua(lua api.Lua) error {
	return lua.Edit()
}

func (a *App) DeleteLua(lua api.Lua) error {
	return lua.Delete()
}

func (a *App) LoadLuas(connectionId string) ([]api.Lua, error) {
	luas := make([]api.Lua, 0)
	if err := api.LoadLuas(&luas); nil != err {
		return nil, err
	}
	_luas := []api.Lua{}
	for _, v := range luas {
		if v.ConnectionID == connectionId {
			_luas = append(_luas, v)
		}
	}
	return _luas, nil
}

// ReadFile .
func (a *App) ReadFile() (string, error) {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return "", err
	}
	bts, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(bts), err
}

// Quit
func (a *App) Quit() {
	runtime.Quit(a.ctx)
}
