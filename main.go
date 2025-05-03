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

	//test "CRM_System/Backend/internal/utils"

	"CRM_System/internal/Test"
	"CRM_System/internal/db"
	routes "CRM_System/internal/routes"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	fmt.Println("Ну типо запуск")

	db.Init()

	Test.TestAll()

	//FastTestInf_AllGroupAndSubject()

}

func FastTestInf_AllGroupAndSubject() {
	data, err := routes.Inf_AllGroupAndSubject()
	if err != nil {
		fmt.Println("Ошибка при получении данных:", err)
		return
	}

	for _, group := range data {
		fmt.Println("Группа:")
		fmt.Printf(" ID: %d\n", group.Id)
		fmt.Printf(" Курс: %d\n", group.Course)
		fmt.Printf(" Специальность: %s\n", group.Spesiality)
		fmt.Printf(" Годы обучения: %d\n", group.Groduates)
		fmt.Printf(" Номер группы: %d\n", group.Number)

		fmt.Println(" Предметы 1 семестра:")
		for _, subj := range group.Subject.FirstSemester {
			fmt.Printf("  ID: %d, Название: %s\n", subj.Id, subj.Title)
		}

		fmt.Println(" Предметы 2 семестра:")
		for _, subj := range group.Subject.SecondSemester {
			fmt.Printf("  ID: %d, Название: %s\n", subj.Id, subj.Title)
		}

		fmt.Println("---------------------------------")
	}
}
