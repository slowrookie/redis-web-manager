package api

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var APP_ROOT string

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
	APP_ROOT = AppDataPath()
	if err := os.MkdirAll(APP_ROOT, os.ModePerm); err != nil {
		panic(err)
	}

	// storage
	GlobalStorage.Initialize(path.Join(APP_ROOT, "storage"))
}
