package dtos

type Inf_EmpByStudentId struct {
	StudentId int `json:"StudentId"`
}

type Create_Employers struct {
	StudId        int    `json:"StudId"`
	Enterprise    string `json:"Enterprise"`
	WorkStartDate string `json:"WorkStartDate"`
	JobTitle      string `json:"JobTitle"`
}

type Update_EmpbyStudentId struct {
	StudId           int    `json:"StudId"`
	NewEnterprise    string `json:"NewEnterprise"`
	NewWorkStartDate string `json:"NewWorkStartDate"`
	NewJobTitle      string `json:"NewJobTitle"`
}

type Delete_EmpbyStudentId struct {
	StudId int `json:"StudId"`
}
