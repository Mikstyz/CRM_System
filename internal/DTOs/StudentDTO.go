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
	GroupId  int    `json:"GroupId"`

	Enterprise    string `json:"Enterprise"`
	WorkStartDate string `json:"WorkStartDate"`
	JobTitle      string `json:"JobTitle"`
}

type Update_StudentById struct {
	StudId      int    `json:"StudId"`
	NewFullName string `json:"NewFullName"`
	NewGroupId  int    `josn:"NewGroupId"`

	NewEnterprise    string `json:"NewEnterprise"`
	NewWorkStartDate string `json:"NewWorkStartDate"`
	NewJobTitle      string `json:"NewJobTitle"`
}

type Delete_Student struct {
	StudId int `json:"StudId"`
}
