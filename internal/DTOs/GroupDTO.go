package dtos

type CreateGroupDTO struct {
	Course     byte
	Groudates  byte
	Speciality string
	Number     int
	Semester   byte
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
	GroupId int
}

type GetGroupIdByInfo struct {
	Course     byte
	Groudates  byte
	Speciality string
	Number     int
	Semester   byte
}
