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
	Course     byte   `json:"course,omitempty"`
	Groudates  byte   `json:"groudates,omitempty"`
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
type OutEmployer struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`

	Enterprise    string `json:"enterprise,omitempty"`
	WorkStartDate string `json:"work_start_date,omitempty"`
	JobTitle      string `json:"job_title,omitempty"`
}

type Inf_Employer struct {
	Code     int     `json:"code"`
	Employer EInfEmp `json:"employer,omitempty"`
	Error    string  `json:"error,omitempty"`
}

// Group
type CrtOutGroupApi struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`

	Id         [2]int `json:"id,omitempty"`
	Course     byte   `json:"course,omitempty"`
	Groudates  byte   `json:"groudates,omitempty"`
	Speciality string `json:"speciality,omitempty"`
	Number     int    `json:"group_num,omitempty"`
	Semester   byte   `json:"semester,omitempty"`
}

type UpdateOutGroupApi struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`

	Id         int    `json:"id,omitempty"`
	Course     byte   `json:"course,omitempty"`
	Groudates  byte   `json:"groudates,omitempty"`
	Speciality string `json:"speciality,omitempty"`
	Number     int    `json:"group_num,omitempty"`
	Semester   byte   `json:"semester,omitempty"`
}

type Inf_AllGroup struct {
	Code   int         `json:"code"`
	Groups []EinfGroup `json:"groups,omitempty"`
	Error  string      `json:"error,omitempty"`
}

// Subject
type OutSubject struct {
	Code    int    `json:"code"`
	Error   string `json:"error,omitempty"`
	GroupId int    `json:"group_id,omitempty"`
	Subject string `json:"subject_name,omitempty"`
}

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
