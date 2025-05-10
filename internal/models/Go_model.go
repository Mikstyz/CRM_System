package models

type DateOnly struct {
	Year  int
	Month byte
	Day   byte
}

type GeneratePDF struct {
	//данные о студенте
	StudentName string

	//данные о работе
	Enterprise    string
	WorkStartDate string
	JobTitle      string
}
