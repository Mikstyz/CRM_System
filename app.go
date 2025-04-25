package main

import (
	"CRM_System/internal/models"
	routes "CRM_System/internal/routes"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

//----------------- Student ----------------

func (a *App) Inf_AllStudent() models.Inf_AllStudents {
	var info models.Inf_AllStudents

	stud, err := routes.Inf_AllStudent()

	code := 200

	if err != nil {
		code = 500
		info.Error = err.Error()
	}

	info.Code = code
	info.Students = stud

	return info
}

func (a *App) Inf_StudentByID(studentID int) models.Inf_Student {
	var info models.Inf_Student
	stud, err := routes.Inf_StudentByID(studentID)
	code := 200
	if err != nil {
		code = 500
		info.Error = err.Error()
	}
	info.Code = code
	info.Student = append(info.Student, *stud)
	return info
}

func (a *App) Inf_StudentbyGroup() {

}

func (a *App) CreateStudent(fullName string, course byte, groudates byte, speciality string, number int, Semester byte) models.OutStudentAPI {
	var res models.OutStudentAPI

	id, err := routes.Create_Student(fullName, course, groudates, speciality, number, Semester)

	code := 200

	if err != nil {
		code = 500
		res.Error = err.Error()
		return res
	}

	res.Code = code
	res.Id = id
	res.FullName = fullName
	res.Speciality = speciality
	res.Number = number
	res.Semester = Semester
	res.Course = course
	res.Groudates = groudates

	return res
}

func (a *App) UpdateStudentById(studId int, newFullName string, newCourse byte, newGroudates byte, newSpeciality string, newNumber int, newSemester byte) models.OutStudentAPI {
	var res models.OutStudentAPI

	status, err := routes.Update_StudentById(studId, newFullName, newCourse, newGroudates, newSpeciality, newNumber, newSemester)

	code := 200

	if err != nil || status == false {
		code = 500
		res.Error = err.Error()
	}

	res.Code = code
	res.Id = studId
	res.FullName = newFullName
	res.Speciality = newSpeciality
	res.Number = newNumber
	res.Semester = newSemester
	res.Course = newCourse
	res.Groudates = newGroudates

	return res
}

func (a *App) DeleteStudents(studId int) models.Remove {
	var res models.Remove
	status, err := routes.Delete_Student(studId)
	code := 200
	if err != nil || status == false {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code

	return res
}

//----------------- Employers --------------

func (a *App) Inf_EmpByStudentId(StudentId int) models.Inf_Employer {
	var info models.Inf_Employer
	emp, err := routes.Inf_EmpByStudentId(StudentId)
	code := 200
	if err != nil {
		code = 500
		info.Error = err.Error()
	}
	info.Code = code
	info.Employer = emp
	return info
}

func (a *App) Update_EmpbyStudentId(studId int, newEnterprise string, newWorkStartDate string, newJobTitle string) models.OutEmployer {
	var res models.OutEmployer
	_, err := routes.Update_EmpbyStudentId(studId, newEnterprise, newWorkStartDate, newJobTitle)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	} else {
		res.Enterprise = newEnterprise
		res.WorkStartDate = newWorkStartDate
		res.JobTitle = newJobTitle
	}
	res.Code = code
	return res
}

func (a *App) Delete_EmpbyStudentId(studId int) models.Remove {
	var res models.Remove
	_, err := routes.Delete_EmpbyStudentId(studId)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	return res
}

//----------------- Group -------------------

func (a *App) Inf_AllGroup() models.Inf_AllGroup {
	var info models.Inf_AllGroup
	groups, err := routes.Inf_AllGroup()
	code := 200
	if err != nil {
		code = 500
		info.Error = err.Error()
	}
	info.Code = code
	info.Groups = groups
	return info
}

func (a *App) Create_Group(Course byte, Groudates byte, speciality string, Number int, Semester byte) models.OutGroupApi {
	var res models.OutGroupApi

	id, err := routes.Create_Group(Course, Groudates, speciality, Number, Semester)

	code := 200

	if err != nil {
		code = 500
		res.Error = err.Error()
	}

	res.Code = code

	res.Id = id
	res.Course = Course
	res.Groudates = Groudates
	res.Speciality = speciality
	res.Number = Number
	res.Semester = Semester

	return res
}

func (a *App) Update_GroupById(groupId int, newCourse byte, newGroudates byte, newSpeciality string, newNumber int, newSemester byte) models.OutGroupApi {
	var res models.OutGroupApi

	status, err := routes.Update_GroupById(groupId, newCourse, newGroudates, newSpeciality, newNumber, newSemester)

	code := 200

	if err != nil || status == false {
		code = 500
		res.Error = err.Error()
	}

	res.Code = code

	res.Id = groupId
	res.Course = newCourse
	res.Groudates = newGroudates
	res.Speciality = newSpeciality
	res.Number = newNumber
	res.Semester = newSemester

	return res
}

func (a *App) Delete_GroupById(groupId int) models.Remove {
	var res models.Remove
	_, err := routes.Delete_GroupById(groupId)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	return res
}

func (a *App) GetGroupId_GroupIdByInfo(well byte, gClass byte, speciality string, groupNum int, semester byte) models.GetId {
	var res models.GetId
	id, err := routes.GetGroupId_GroupIdByInfo(well, gClass, speciality, groupNum, semester)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	res.Id = id
	return res
}

//----------------- Subject -----------------

func (a *App) Inf_SubjectByGroupId(groupId int) models.Inf_Subject {
	var info models.Inf_Subject
	subjects, err := routes.Inf_SubjectByGroupId(groupId)
	code := 200
	if err != nil {
		code = 500
		info.Error = err.Error()
	}
	info.Code = code
	info.Subject = subjects
	return info
}

func (a *App) Add_SubjectByGroupId(groupId int, newSubject string) models.OutSubject {
	var res models.OutSubject

	id, err := routes.Add_SubjectByGroupId(groupId, newSubject)
	code := 200

	if err != nil {
		code = 500
		res.Error = err.Error()
	} else {
		res.Subject = newSubject
		res.GroupId = groupId
	}

	res.Code = code
	res.GroupId = id

	return res
}

func (a *App) Update_SubjectById(subjectId int, newSubject string) models.OutSubject {
	var res models.OutSubject

	status, err := routes.Update_SubjectById(subjectId, newSubject)
	code := 200

	if err != nil || !status {
		code = 500
		if err != nil {
			res.Error = err.Error()
		} else {
			res.Error = "не удалось обновить предмет"
		}
	} else {
		res.Subject = newSubject
	}

	res.Code = code
	return res
}

func (a *App) Delete_SubjectById(subjectId int) models.Remove {
	var res models.Remove
	_, err := routes.Delete_SubjectById(subjectId)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	return res
}

func (a *App) Delete_AllSubjectByGroupId(groupID int) models.Remove {
	var res models.Remove
	_, err := routes.Delete_AllSubjectByGroupId(groupID)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	return res
}

//----------------- PDF ----------------------

func (a *App) GenerateFilledPDF(dockDATA models.GeneratePDF) models.PdfDock {
	var res models.PdfDock
	file, err := routes.GenerateFilledPDF(dockDATA)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	res.File = file
	return res
}
