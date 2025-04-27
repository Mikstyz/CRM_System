package dtos

type CreateGroupDTO struct {
	Course     byte
	Groudates  byte
	Speciality string
	Number     int
}

type UpdateGroupDTO struct {
	GroupId       int
	NewCourse     byte
	NewGroudates  byte
	NewSpeciality string
	NewNumber     int
	NewSemester   byte
}

type DeleteGroupDTO struct {
	Course     byte
	Graduates  byte
	Speciality string
	Number     int
}

type GetGroupIdByInfo struct {
	Course     byte
	Groudates  byte
	Speciality string
	Number     int
	Semester   byte
}
