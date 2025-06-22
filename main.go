package main

import (
	"embed"
	"log"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	//"github.com/wailsapp/wails/v2"
	//"github.com/wailsapp/wails/v2/pkg/options"
	//"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed app/Data/Sql/stud.db
var embeddedDB embed.FS

func main() {
	if err := ensureDBFromTemplate(); err != nil {
		log.Fatalf("[main] Не удалось подготовить БД: %v", err)
	}

	app := NewApp()

	err := wails.Run(&options.App{
		Title:  "CRM_System",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Debug:     options.Debug{OpenInspectorOnStartup: true},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func ensureDBFromTemplate() error {
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	exeDir := filepath.Dir(exePath)

	targetPath := filepath.Join(exeDir, "Data", "Sql", "stud.db")

	if _, err := os.Stat(targetPath); err == nil {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(targetPath), os.ModePerm); err != nil {
		return err
	}

	data, err := embeddedDB.ReadFile("app/Data/Sql/stud.db")
	if err != nil {
		return err
	}

	// Записать файл
	return os.WriteFile(targetPath, data, 0644)
}
