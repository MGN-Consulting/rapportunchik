package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Initialize MongoDB database connection
	_, err := InitDB()
	if err != nil {
		log.Printf("Warning: Failed to connect to MongoDB: %v. Please ensure mongod is running on localhost:27017.", err)
	}

	rapportService := NewRapportService()

	app := application.New(application.Options{
		Name:        "Rapportunchik",
		Description: "Collaborazione Partita IVA | Timesheet & Monthly Activity Reporter",
		Services: []application.Service{
			application.NewService(rapportService),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "Rapportunchik",
		Width:  1200,
		Height: 780,
		MinWidth: 1050,
		MinHeight: 680,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 48,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(8, 10, 20),
		URL:              "/",
	})

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
