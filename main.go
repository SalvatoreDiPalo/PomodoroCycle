package main

import (
	"embed"
	"fmt"
	"pomodoro/backend"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	myLog := backend.NewFileLogger(fmt.Sprintf("info-%v.log", time.Now().Format("2006-01-02")))
	// Create an instance of the app structure
	app := backend.NewApp(myLog)

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "PomodoroCycle",
		Width:         438,
		Height:        625,
		DisableResize: true,
		Fullscreen:    false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour:   &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		Logger:             myLog,
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.ERROR,
		OnStartup:          app.Startup,
		OnShutdown:         app.Shutdown,
		Bind: []interface{}{
			app,
		},
		Mac: &mac.Options{
			About: &mac.AboutInfo{
				Title:   "PomodoroCycle",
				Message: "© 2024 Salvatore Di Palo",
			},
		},
		Linux: &linux.Options{
			ProgramName: "PomodoroCycle",
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
