package dtos

type Inf_AllStudent struct{}

type Inf_StudentByID struct {
	StudentID int `json:"StudentID"`
}

type Inf_StudentByGroup struct {
	GroupId int `json:"GroupId"`
}

type Create_Student struct {
	FullName string `json:"FullName"`
	GroupId  int    `josn:"GroupId"`
}

type Update_StudentById struct {
	StudId      int    `json:"StudId"`
	NewFullName string `json:"NewFullName"`
	NewGroupId  int    `josn:"NewGroupId"`
}

type Delete_Student struct {
	StudId int `json:"StudId"`
}
