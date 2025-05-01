package dtos

// InfAllGroupDTO для запроса всех групп
type InfAllGroupDTO struct{}

type InfAllGroupAndSubjectDTO struct {
	SwitchSubject bool `json:"Switch"`
}

// CreateGroupDTO для создания группы
type CreateGroupDTO struct {
	Course     byte   `json:"course"`
	Graduates  byte   `json:"groudates"` // Исправлено groduates на graduates
	Speciality string `json:"speciality"`
	Number     int    `json:"group_num"`
}

// UpdateGroupDTO для обновления группы
type UpdateGroupDTO struct {
	GroupId       int    `json:"group_id"`
	NewCourse     byte   `json:"new_course"`
	NewGraduates  byte   `json:"new_groudates"` // Исправлено groduates
	NewSpeciality string `json:"new_speciality"`
	NewNumber     int    `json:"new_group_num"`
}

// DeleteGroupDTO для удаления группы
type DeleteGroupDTO struct {
	GroupId int `json:"group_id"`
}

// GetGroupIdByInfoDTO для получения ID группы
type GetGroupIdByInfoDTO struct {
	Course     byte   `json:"course"`
	Graduates  byte   `json:"groudates"` // Исправлено groduates
	Speciality string `json:"speciality"`
	Number     int    `json:"group_num"`
	Semester   byte   `json:"semester"`
}

// DublicateGroupDTO для дублирования группы
type DublicateGroupDTO struct {
	Course     byte   `json:"course"`
	Graduates  byte   `json:"groudates"` // Исправлено groduates
	Speciality string `json:"speciality"`
	Number     int    `json:"group_num"`
}
