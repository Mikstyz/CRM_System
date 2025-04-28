package models

type EinfGroup struct {
	Id         int
	Course     byte
	Speciality string
	Groudates  byte
	Number     int
}

type InFGroupAndSubject struct {
	Id         int
	Course     byte
	Spesiality string
	Groduates  byte
	Number     int

	Subject DisciplinesBySemester
}
