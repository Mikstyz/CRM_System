package models

type Student struct {
	ID            int     `json:"id"`
	FullName      string  `json:"full_name"`
	GroupId       int     `json:"GroupId"`
	Enterprise    *string `json:"Enterprise"`
	WorkStartDate *string `json:"WorkStartDate"`
	JobTitle      *string `json:"JobTitle"`
}
