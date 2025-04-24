package models

type DateOnly struct {
	Year  int
	Month byte
	Day   byte
}

type PdfDoc struct {
	Name          string
	Enterprise    string
	WorkStartDate string
	JobTitle      string
}

type GeneratePDF struct {
	//данные о студенте
	StudentName string

	//данные группы
	GroupId int

	Course     int
	Speciality string
	Groduates  byte
	Number     int

	//данные о работе
	Enterprise    string
	WorkStartDate string
	JobTitle      string
}
