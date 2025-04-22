package modeles

//All
type GetId struct {
	Code  int
	Id    int
	Error string
}

type Create struct {
	Code  int
	Id    int
	Error string
}

type Update struct {
	Code  int
	Error string
}

type Remove struct {
	Code  int
	Error string
}

//Student
type Inf_AllStudents struct {
	Code     int
	Students []Student
	Error    string
}

type Inf_Student struct {
	Code    int
	Student []Student
	Error   string
}

//Employers
type Inf_Employer struct {
	Code     int
	Employer EInfEmp
	Error    string
}

//Group
type Inf_AllGroup struct {
	Code   int
	Groups []EinfGroup
	Error  string
}

//Get Id

// Subject

type Inf_Subject struct {
	Code    int
	Subject []string
	Error   string
}

type PdfDock struct {
	Code  int
	File  []byte
	Error string
}
