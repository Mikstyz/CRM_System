package models

type Student struct {
	//Данные студента
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	GroupId  int    `json:"GroupId"`

	//Данные работодателя
	Enterprise    *string `json:"Enterprise"`
	WorkStartDate *string `json:"WorkStartDate"`
	JobTitle      *string `json:"JobTitle"`
}
