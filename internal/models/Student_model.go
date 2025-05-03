package models

type Student struct {
	ID         int    `json:"id"`
	FullName   string `json:"full_name"`
	GroupId    int    `json:"GroupId"`
	Speciality string `json:"speciality"`
	GroupNum   int    `json:"group_num"`
	Semester   int    `json:"semester"`
	Course     int    `json:"Course"`
	Groudates  int    `json:"Groduates"`
}

type StudentInput struct {
	FullName   string `db:"FullName"`
	Speciality string `db:"Speciality"`
	GroupNum   int    `db:"GroupNum"`
	Semester   int16  `db:"Semester"`
	Course     int16  `db:"Course"`
	Groudates  int16  `db:"Groduates"`
}
