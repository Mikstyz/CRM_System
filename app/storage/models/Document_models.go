package models

type GeneratePDF struct {
	//данные о студенте
	StudentName string

	//Данные о группе
	GroupId  int
	Semester byte

	//данные о работе
	Enterprise    string
	WorkStartDate string
	JobTitle      string

	//Папка
	Path string
}
