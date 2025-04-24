package models

type EinfGroup struct {
	Id         int
	Course     int
	Speciality string
	Groudates  byte
	Semester   byte
	Number     int
}

type Subject struct {
	Name []string
}
