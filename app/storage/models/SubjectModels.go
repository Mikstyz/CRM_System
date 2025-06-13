package models

type SemesterDiscipline struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type DisciplinesBySemester struct {
	OneSemester []SemesterDiscipline `json:"OneSemester"`
	TwoSemester []SemesterDiscipline `json:"TwoSemester"`
}
