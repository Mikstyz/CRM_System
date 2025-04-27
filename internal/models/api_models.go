package models

// -------------------- [ All / Общие типы ] --------------------

// AppRemove для возврата результата удаления
type AppRemove struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`
}

// AppGetId для возврата ID группы
type AppGetId struct {
	Code  int    `json:"code"`
	Id    int    `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

// -------------------- [ Group / Группы ] --------------------

// AppInf_AllGroup для возврата списка всех групп
type AppInf_AllGroup struct {
	Code   int         `json:"code"`
	Groups []EinfGroup `json:"groups,omitempty"`
	Error  string      `json:"error,omitempty"`
}

// AppCrtOutGroupApi для возврата данных при создании группы
type AppCrtOutGroupApi struct {
	Code  int       `json:"code"`
	One   EinfGroup `json:One,omitempty`
	Two   EinfGroup `json:Two,omitempty`
	Error string    `json:"error,omitempty"`
}

// AppUpdateOutGroupApi для возврата данных при обновлении группы
type AppUpdateOutGroupApi struct {
	Code       int    `json:"code"`
	Error      string `json:"error,omitempty"`
	Id         int    `json:"id,omitempty"`
	Course     byte   `json:"course,omitempty"`
	Graduates  byte   `json:"groudates,omitempty"`
	Speciality string `json:"speciality,omitempty"`
	Number     int    `json:"group_num,omitempty"`
}

// -------------------- [ Subject / Предметы ] --------------------

// AppInf_Subject для возврата списка предметов
type AppInf_Subject struct {
	Code    int      `json:"code"`
	Subject []string `json:"subject,omitempty"`
	Error   string   `json:"error,omitempty"`
}

// AppInf_Disciplines для возврата предметов по семестрам
type AppInf_Disciplines struct {
	Code        int                   `json:"code"`
	Disciplines DisciplinesBySemester `json:"disciplines,omitempty"`
	Error       string                `json:"error,omitempty"`
}

// AppOutSubject для возврата данных о предмете при создании/обновлении
type AppOutSubject struct {
	Code    int    `json:"code"`
	Id      int    `json:"id,omitempty"`
	GroupId int    `json:"group_id,omitempty"`
	Subject string `json:"subject_name,omitempty"`
	Error   string `json:"error,omitempty"`
}

// -------------------- [ Student / Студенты ] --------------------

type AppInf_AllStudent struct {
	Code     int       `json:"code"`
	Students []Student `json:"students,omitempty"`
	Error    string    `json:"error,omitempty"`
}

type AppInf_StudentByID struct {
	Code    int     `json:"code"`
	Student Student `json:"Student,omitempty"`
	Error   string  `json:"error,omitempty"`
}

type AppInf_StudentByGroup struct {
	Code     int       `json:"code"`
	Students []Student `json:"Student,omitempty"`
	Error    string    `json:"error,omitempty"`
}

type AppCreate_Student struct {
	Code       int    `json:"code"`
	Id         int    `json:"Id"`
	FullName   string `json:"FullName"`
	Course     byte   `json:"Course"`
	Graduates  byte   `json:"Groduates"`
	Speciality string `json:"Speciality"`
	GroupNum   int    `json:"Number"`
	Semester   byte   `json:"Semester"`
	Error      string `json:"error,omitempty"`
}

type AppUpdate_StudentById struct {
	Code          int    `json:"code"`
	NewId         int    `json:"Id"`
	NewFullName   string `json:"NewFullName"`
	NewCourse     byte   `json:"NewCourse"`
	NewGraduates  byte   `json:"NewGroduates"`
	NewSpeciality string `json:"NewSpeciality"`
	NewGroupNum   int    `json:"NewNumber"`
	NewSemester   byte   `json:"NewSemester"`
	Error         string `json:"error,omitempty"`
}

type AppDelete_Student struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`
}

// -------------------- [ PDF / Документы ] --------------------

type AppPdfDock struct {
	Code  int    `json:"code"`
	File  []byte `json"File,omitempty"`
	Error string `json:"error,omitempty"`
}

// -------------------- [ Employers / Данные работодателя ] --------------------

type AppInf_EmpByStudentId struct {
	Code      int     `json:"code"`
	Employers EInfEmp `json:"Employers"`
	Error     string  `json:"error,omitempty"`
}

// type AppCreate_Employers struct {
// 	Code  int    `json:"code"`
// 	Error string `json:"error,omitempty"`
// }

type AppUpdate_EmpbyStudentId struct {
	Code             int    `json:"code"`
	StudId           int    `json:"StudId"`
	NewEnterprise    string `json:"NewEnterprise"`
	NewWorkStartDate string `json:"NewWorkStartDate"`
	NewJobTitle      string `json:"NewJobTitle"`
	Error            string `json:"error,omitempty"`
}

type AppDelete_EmpbyStudentId struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`
}
