package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "CRM_System",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

// import (

// 	//test "CRM_System/Backend/internal/utils"

// 	"CRM_System/internal/Test"
// 	"CRM_System/internal/db"
// 	"fmt"

// 	_ "modernc.org/sqlite"
// )

// func main() {
// 	fmt.Println("Ну типо запуск")

// 	db.Init()

// 	_, _, err, Groups := Test.Test_InfAllGroup()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	for _, i := range Groups {
// 		fmt.Printf("----------")
// 		fmt.Printf("Id: %v\n", i.Id)
// 		fmt.Printf("Course: %v\n", i.Course)
// 		fmt.Printf("Speciality: %v\n", i.Speciality)
// 		fmt.Printf("Groudates: %v\n", i.Groudates)
// 		fmt.Printf("Number: %v\n", i.Number)
// 	}
// }

//Test.TestAll()
