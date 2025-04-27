package main

// import (
// 	"embed"

// 	"github.com/wailsapp/wails/v2"
// 	"github.com/wailsapp/wails/v2/pkg/options"
// 	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
// )

// //go:embed all:frontend/dist
// var assets embed.FS

// func main() {
// 	// Create an instance of the app structure
// 	app := NewApp()

// 	// Create application with options
// 	err := wails.Run(&options.App{
// 		Title:  "CRM_System",
// 		Width:  1024,
// 		Height: 768,
// 		AssetServer: &assetserver.Options{
// 			Assets: assets,
// 		},
// 		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
// 		OnStartup:        app.startup,
// 		Bind: []interface{}{
// 			app,
// 		},
// 	})

// 	if err != nil {
// 		println("Error:", err.Error())
// 	}
// }

import (
	"CRM_System/internal/Test"
	"CRM_System/internal/db"

	//test "CRM_System/Backend/internal/utils"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	fmt.Println("Ну типо запуск")

	db.Init()

	Test.TestSubjectsALL()

	//Test.TestAll()

	//Test.Test_Employers()
	//test.Test_CreateStudentWithEmptyEmployment()

	//Test.TestGenerateFilledPDF()
}
