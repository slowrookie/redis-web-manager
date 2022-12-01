package main

import (
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"io"
	"log"
	"os"
	"path"

	"github.com/slowrookie/redis-web-manager/api"
	_ "github.com/slowrookie/redis-web-manager/api"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
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

//go:embed build/appicon.png
var icon []byte

func main() {
	// log
	f, err := os.OpenFile(path.Join(api.AppRoot, "rwm.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	log.SetOutput(io.MultiWriter(os.Stdout, f))
	log.Println(fmt.Sprintf("Root Path: %s", api.AppRoot))

	app := NewApp()
	var about = make(map[string]string)
	about["version"] = version
	about["commit"] = commit
	about["date"] = date
	about["builtBy"] = builtBy
	about["environment"] = MODE
	app.About = about

	err = wails.Run(&options.App{
		Title:             "Redis Web Manager",
		Width:             1024,
		Height:            800,
		MinWidth:          720,
		MinHeight:         570,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		AssetServer:       &assetserver.Options{Assets: assets},
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnShutdown:        app.shutdown,
		Bind: []interface{}{
			app,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			DisableWindowIcon:    false,
		},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarDefault(),
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "Redis Web Manager",
				Message: "Redis Web Manager is a desktop application developed with React & Golang, used to manage Redis, and supports multi-platform operation.",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
