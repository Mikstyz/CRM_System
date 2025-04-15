package modeles

type EinfGroup struct {
	Id         int
	Well       int
	Speciality string
	GClass     byte
	Semester   byte
	Number     int
}

type Subject struct {
	Name []string
}
