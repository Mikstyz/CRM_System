package modeles

type Student struct {
	ID         int    `json:"id"`
	FullName   string `json:"full_name"`
	Speciality string `json:"speciality"`
	GroupNum   int    `json:"group_num"`
	Semester   int    `json:"semester"`
	Well       int    `json:"well"`
	GClass     int    `json:"g_class"`
}

type StudentInput struct {
	FullName   string `db:"FullName"`
	Speciality string `db:"Speciality"`
	GroupNum   int    `db:"GroupNum"`
	Semester   int16  `db:"Semester"`
	Well       int16  `db:"Well"`
	GClass     int16  `db:"gClass"`
}
