package main

import (
	"CRM_System/app/paths"
	"embed"
	"log"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed app/Data/Sql/stud.db app/Data/Fonts/DejaVuSans.ttf
var embeddedFiles embed.FS

func main() {

	if err := paths.InitPaths(); err != nil {
		log.Fatalf("Ошибка получения путей")
	}

	if err := ensureDBFromTemplate(); err != nil {
		log.Fatalf("Ошибка подготовки БД: %v", err)
	}

	if err := ensureFontsFromTemplate(); err != nil {
		log.Fatalf("Ошибка подготовки шрифтов: %v", err)
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
	targetPath := paths.GetEmbedTarget()

	if _, err := os.Stat(targetPath); err == nil {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(targetPath), os.ModePerm); err != nil {
		return err
	}

	// Путь внутри embed должен совпадать с путём в //go:embed
	data, err := embeddedFiles.ReadFile("app/Data/Sql/stud.db")
	if err != nil {
		return err
	}

	return os.WriteFile(targetPath, data, 0644)
}

func ensureFontsFromTemplate() error {
	// Тоже путь как в //go:embed
	const embedFontPath = "app/Data/Fonts/DejaVuSans.ttf"

	targetFontPath := paths.GetEmbedFontPath()

	if _, err := os.Stat(targetFontPath); err == nil {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(targetFontPath), os.ModePerm); err != nil {
		return err
	}

	data, err := embeddedFiles.ReadFile(embedFontPath)
	if err != nil {
		return err
	}

	return os.WriteFile(targetFontPath, data, 0644)
}
