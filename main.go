package main

import (
	routes "CRM_System/app/api/routes"
	"embed"
	"log"
	//"github.com/wailsapp/wails/v2"
	//"github.com/wailsapp/wails/v2/pkg/options"
	//"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// app := NewApp()

	// err := wails.Run(&options.App{
	// 	Title:  "CRM_System",
	// 	Width:  1024,
	// 	Height: 768,
	// 	AssetServer: &assetserver.Options{
	// 		Assets: assets,
	// 	},
	// 	Debug:     options.Debug{OpenInspectorOnStartup: true},
	// 	OnStartup: app.startup,
	// 	Bind: []interface{}{
	// 		app,
	// 	},
	// })

	// if err != nil {
	// 	println("Error:", err.Error())
	// }

	student, err := routes.Inf_StudentByGroup(81)

	if err != nil {
		log.Printf(err.Error())
	}

	print(student)

}

// import (

// 	//test "CRM_System/Backend/internal/utils"

// 	"CRM_System/internal/db"
// 	models "CRM_System/internal/models"
// 	"CRM_System/internal/utils"
// 	"fmt"
// 	"log"

// 	_ "modernc.org/sqlite"
// )

// func main() {
// 	fmt.Println("Ну типо запуск")

// 	db.Init()

// 	//Test.Test_GenerateFilledPDF()

// 	dataPdf := models.GeneratePDF{
// 		StudentName:   "Иван Иванов",
// 		Enterprise:    "ООО Ромашка",
// 		WorkStartDate: "2025-04-27",
// 		JobTitle:      "Программист",
// 		GroupId:       49,
// 		Semester:      1,
// 	}

// 	pdf, err := utils.GenerateFiledPDF(dataPdf)

// 	if err != nil {
// 		log.Print("Ошибка при генерации документа")
// 	}

// 	if len(pdf) == 0 {
// 		log.Print("Документ пуст")
// 	}

// 	utils.SavePDFToFile(pdf, "pdf")
// 	//FastTestInf_AllGroupAndSubject()

// }

// func FastTestInf_AllGroupAndSubject() {
// 	data, err := routes.Inf_AllGroupAndSubject()
// 	if err != nil {
// 		fmt.Println("Ошибка при получении данных:", err)
// 		return
// 	}

// 	for _, group := range data {
// 		fmt.Println("Группа:")
// 		fmt.Printf(" ID: %d\n", group.Id)
// 		fmt.Printf(" Курс: %d\n", group.Course)
// 		fmt.Printf(" Специальность: %s\n", group.Spesiality)
// 		fmt.Printf(" Годы обучения: %d\n", group.Groduates)
// 		fmt.Printf(" Номер группы: %d\n", group.Number)

// 		fmt.Println(" Предметы 1 семестра:")
// 		for _, subj := range group.Subject.FirstSemester {
// 			fmt.Printf("  ID: %d, Название: %s\n", subj.Id, subj.Title)
// 		}

// 		fmt.Println(" Предметы 2 семестра:")
// 		for _, subj := range group.Subject.SecondSemester {
// 			fmt.Printf("  ID: %d, Название: %s\n", subj.Id, subj.Title)
// 		}

// 		fmt.Println("---------------------------------")
// 	}
// }
