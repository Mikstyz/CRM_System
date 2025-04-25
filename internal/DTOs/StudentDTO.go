package dtos

type InfStudByIdDTO struct {
	StudentID int
}

type CreateStudDTO struct {
	FullName   string
	Course     byte
	Groudates  byte
	Speciality string
	Number     int
	Semester   byte
}

type UpdateStudDTO struct {
	StudId        int
	NewFullName   string
	NewCourse     byte
	NewGroudates  byte
	NewSpeciality string
	NewNumber     int
	NewSemester   byte
}

type DeleteStudDTO struct {
	StudId int
}
