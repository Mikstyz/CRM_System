package dtos

// InfSubjectDTO для запроса предметов по ID группы
type InfSubjectDTO struct {
	GroupId int `json:"group_id"`
}

// InfDisciplinesByGroupDataDTO для запроса предметов по данным группы
type InfDisciplinesByGroupDataDTO struct {
	Speciality string `json:"speciality"`
	GroupNum   int    `json:"group_num"`
	Course     int    `json:"course"`
	Groudates  int    `json:"groudates"`
}

// AddSubjectDTO для добавления предмета в группу
type AddSubjectDTO struct {
	GroupId    int    `json:"group_id"`
	NewSubject string `json:"new_subject"`
}

// UpdateSubjectDTO для обновления предмета
type UpdateSubjectDTO struct {
	SubjectId  int    `json:"subject_id"`
	NewSubject string `json:"new_subject"`
}

// DeleteSubjectDTO для удаления предмета
type DeleteSubjectDTO struct {
	SubjectId int `json:"subject_id"`
}

// DeleteAllSubjectByGroupDTO для удаления всех предметов группы
type DeleteAllSubjectByGroupDTO struct {
	GroupId int `json:"group_id"`
}
