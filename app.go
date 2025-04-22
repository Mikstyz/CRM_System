package main

import (
	"CRM_System/internal/modeles"
	"CRM_System/internal/routes"
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

func (a *App) Inf_AllStudent() modeles.Inf_AllStudents {
	var info modeles.Inf_AllStudents
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

func (a *App) Inf_StudentByID(studentID int) modeles.Inf_Student {
	var info modeles.Inf_Student
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

func (a *App) Create_Student(fullName string, well byte, gClass byte, speciality string, groupNum int, semester byte) modeles.Create {
	var res modeles.Create
	id, err := routes.Create_Student(fullName, well, gClass, speciality, groupNum, semester)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	res.Id = id
	return res
}

func (a *App) Update_StudentById(studId int, newFullName string, newWell byte, newClass byte, newSpeciality string, newGroupNum int, newSemester byte) modeles.Update {
	var res modeles.Update
	_, err := routes.Update_StudentById(studId, newFullName, newWell, newClass, newSpeciality, newGroupNum, newSemester)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	return res
}

func (a *App) Delete_Student(studId int) modeles.Remove {
	var res modeles.Remove
	_, err := routes.Delete_Student(studId)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	return res
}

//----------------- Employers --------------

func (a *App) Inf_EmpByStudentId(StudentId int) modeles.Inf_Employer {
	var info modeles.Inf_Employer
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

func (a *App) Update_EmpbyStudentId(studId int, newEnterprise string, newWorkStartDate string, newJobTitle string) modeles.Update {
	var res modeles.Update
	_, err := routes.Update_EmpbyStudentId(studId, newEnterprise, newWorkStartDate, newJobTitle)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	return res
}

func (a *App) Delete_EmpbyStudentId(studId int) modeles.Remove {
	var res modeles.Remove
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

func (a *App) Inf_AllGroup() modeles.Inf_AllGroup {
	var info modeles.Inf_AllGroup
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

func (a *App) Create_Group(well byte, gClass byte, speciality string, groupNum int, semester byte) modeles.Create {
	var res modeles.Create
	id, err := routes.Create_Group(well, gClass, speciality, groupNum, semester)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	res.Id = id
	return res
}

func (a *App) Update_GroupById(groupId int, newWell byte, newGClass byte, newSpeciality string, newGroupNum int, newSemester byte) modeles.Update {
	var res modeles.Update
	_, err := routes.Update_GroupById(groupId, newWell, newGClass, newSpeciality, newGroupNum, newSemester)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	return res
}

func (a *App) Delete_GroupById(groupId int) modeles.Remove {
	var res modeles.Remove
	_, err := routes.Delete_GroupById(groupId)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	return res
}

func (a *App) GetGroupId_GroupIdByInfo(well byte, gClass byte, speciality string, groupNum int, semester byte) modeles.GetId {
	var res modeles.GetId
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

func (a *App) Inf_SubjectByGroupId(groupId int) modeles.Inf_Subject {
	var info modeles.Inf_Subject
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

func (a *App) Add_SubjectByGroupId(groupId int, newSubject string) modeles.Create {
	var res modeles.Create
	id, err := routes.Add_SubjectByGroupId(groupId, newSubject)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	res.Id = id
	return res
}

func (a *App) Update_SubjectById(subjectId int, newSubject string) modeles.Update {
	var res modeles.Update
	_, err := routes.Update_SubjectById(subjectId, newSubject)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	return res
}

func (a *App) Delete_SubjectById(subjectId int) modeles.Remove {
	var res modeles.Remove
	_, err := routes.Delete_SubjectById(subjectId)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	return res
}

func (a *App) Delete_AllSubjectByGroupId(groupID int) modeles.Remove {
	var res modeles.Remove
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

func (a *App) GenerateFilledPDF(StudentId int, StudentName string, GroupId int) modeles.PdfDock {
	var res modeles.PdfDock
	file, err := routes.GenerateFilledPDF(StudentId, StudentName, GroupId)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	res.File = file
	return res
}
