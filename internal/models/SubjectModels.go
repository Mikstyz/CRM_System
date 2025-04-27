package models

type SemesterDiscipline struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type DisciplinesBySemester struct {
	FirstSemester  []SemesterDiscipline `json:"firstSemester"`
	SecondSemester []SemesterDiscipline `json:"secondSemester"`
}
