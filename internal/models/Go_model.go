package models

type DateOnly struct {
	Year  int
	Month byte
	Day   byte
}

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
}
