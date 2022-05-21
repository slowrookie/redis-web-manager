package api

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var AppRoot string

func AppDataPath() string {
	const project = "com.github.slowrookie.redis-web-manager"

	switch runtime.GOOS {
	case "windows":
		return fmt.Sprintf("%s/%s", os.Getenv("APPDATA"), project)
	case "darwin":
		return fmt.Sprintf("%s/Library/Containers/%s", os.Getenv("HOME"), project)
	case "linux":
		return fmt.Sprintf("%s/.%s", os.Getenv("HOME"), project)
	}
	// default path
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func init() {
	AppRoot = AppDataPath()
	if err := os.MkdirAll(AppRoot, os.ModePerm); err != nil {
		panic(err)
	}

	// storage
	if err := GlobalStorage.Initialize(path.Join(AppRoot, "storage")); err != nil {
		panic(err)
	}
}
