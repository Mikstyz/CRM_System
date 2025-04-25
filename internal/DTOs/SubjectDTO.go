package dtos

type InfSubjectDTO struct {
	GroupId int
}

type AddSubjectDTO struct {
	GroupId    int
	NewSubject string
}

type UpdateSubjectDTO struct {
	SubjectId  int
	NewSubject string
}

type DeleteSubjectById struct {
	SubjectId int
}

type DeleteAddSubjectByGroup struct {
	GroupId int
}
