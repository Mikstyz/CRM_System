package dtos

type InfEmployersDTO struct {
	StudInt int
}

type UpdateEmployersDTO struct {
	StudId           int
	NewEnterprise    string
	NewWorkStartDate string
	NewJobTitle      string
}
type DeleteEmployersDTO struct {
	StudId int
}
