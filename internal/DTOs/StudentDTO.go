package dtos

type Inf_AllStudent struct{}

type Inf_StudentByID struct {
	StudentID int `json:"StudentID"`
}

type Inf_StudentByGroup struct {
	Course     byte   `json:"Course"`
	Speciality string `json:"Speciality"`
	GroupNum   int    `json:"Number"`
}

type Create_Student struct {
	FullName   string `json:"FullName"`
	Course     byte   `json:"Course"`
	Graduates  byte   `json:"Groduates"`
	Speciality string `json:"Speciality"`
	GroupNum   int    `json:"Number"`
}

type Update_StudentById struct {
	StudId        int    `json:"StudId"`
	NewFullName   string `json:"NewFullName"`
	NewCourse     byte   `json:"NewCourse"`
	NewGraduates  byte   `json:"NewGroudates"`
	NewSpeciality string `json:"NewSpeciality"`
	NewNumber     int    `json:"NewNumber"`
}

type Delete_Student struct {
	StudId int `json:"StudId"`
}
