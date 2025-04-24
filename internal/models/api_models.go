package models

// All
type GetId struct {
	Code  int    `json:"code"`
	Id    int    `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

type Create struct {
	Code  int    `json:"code"`
	Id    int    `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

type Update struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`
}

type Remove struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`
}

// Student
type OutStudentAPI struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`

	Id         int    `json:"id,omitempty"`
	FullName   string `json:"full_name,omitempty"`
	Speciality string `json:"speciality,omitempty"`
	Number     int    `json:"group_num,omitempty"`
	Semester   byte   `json:"semester,omitempty"`
	Course     byte   `json:"Course,omitempty"`
	Groudates  byte   `json:"Groduates,omitempty"`
}

type Inf_AllStudents struct {
	Code     int       `json:"code"`
	Students []Student `json:"students,omitempty"`
	Error    string    `json:"error,omitempty"`
}

type Inf_Student struct {
	Code    int       `json:"code"`
	Student []Student `json:"student,omitempty"`
	Error   string    `json:"error,omitempty"`
}

// Employers
type Inf_Employer struct {
	Code     int     `json:"code"`
	Employer EInfEmp `json:"employer,omitempty"`
	Error    string  `json:"error,omitempty"`
}

// Group
type OutGroupApi struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`

	Id         int    `json:"Id,omitempty"`
	Course     byte   `json:"Course,omitempty"`
	Groudates  byte   `json:"Groudates,omitempty"`
	Speciality string `json:"speciality,omitempty"`
	Number     int    `json:"Number,omitempty"`
	Semester   byte   `json:"Semester,omitempty"`
}

type Inf_AllGroup struct {
	Code   int         `json:"code"`
	Groups []EinfGroup `json:"groups,omitempty"`
	Error  string      `json:"error,omitempty"`
}

// Subject
type Inf_Subject struct {
	Code    int      `json:"code"`
	Subject []string `json:"subject,omitempty"`
	Error   string   `json:"error,omitempty"`
}

// Pdf
type PdfDock struct {
	Code  int    `json:"code"`
	File  []byte `json:"file,omitempty"`
	Error string `json:"error,omitempty"`
}
