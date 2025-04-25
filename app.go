package main

import (
	dtos "CRM_System/internal/DTOs"
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

func (a *App) Inf_StudentByID(dto dtos.InfStudByIdDTO) models.Inf_Student {
	var info models.Inf_Student
	stud, err := routes.Inf_StudentByID(dto.StudentID)
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

func (a *App) CreateStudent(dto dtos.CreateStudDTO) models.OutStudentAPI {
	var res models.OutStudentAPI

	id, err := routes.Create_Student(dto.FullName, dto.Course, dto.Groudates, dto.Speciality, dto.Number, dto.Semester)

	code := 200

	if err != nil {
		code = 500
		res.Error = err.Error()
		return res
	}

	res.Code = code
	res.Id = id
	res.FullName = dto.FullName
	res.Speciality = dto.Speciality
	res.Number = dto.Number
	res.Semester = dto.Semester
	res.Course = dto.Course
	res.Groudates = dto.Groudates

	return res
}

func (a *App) UpdateStudentById(dto dtos.UpdateStudDTO) models.OutStudentAPI {
	var res models.OutStudentAPI

	status, err := routes.Update_StudentById(dto.StudId, dto.NewFullName, dto.NewCourse, dto.NewGroudates, dto.NewSpeciality, dto.NewNumber, dto.NewSemester)

	code := 200

	if err != nil || status == false {
		code = 500
		res.Error = err.Error()
	}

	res.Code = code
	res.Id = dto.StudId
	res.FullName = dto.NewFullName
	res.Speciality = dto.NewSpeciality
	res.Number = dto.NewNumber
	res.Semester = dto.NewSemester
	res.Course = dto.NewCourse
	res.Groudates = dto.NewGroudates

	return res
}

func (a *App) DeleteStudents(dto dtos.DeleteStudDTO) models.Remove {
	var res models.Remove
	status, err := routes.Delete_Student(dto.StudId)
	code := 200
	if err != nil || status == false {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code

	return res
}

//----------------- Employers --------------

func (a *App) Inf_EmpByStudentId(dto dtos.InfEmployersDTO) models.Inf_Employer {
	var info models.Inf_Employer
	emp, err := routes.Inf_EmpByStudentId(dto.StudInt)
	code := 200
	if err != nil {
		code = 500
		info.Error = err.Error()
	}
	info.Code = code
	info.Employer = emp
	return info
}

func (a *App) Update_EmpbyStudentId(dto dtos.UpdateEmployersDTO) models.OutEmployer {
	var res models.OutEmployer
	_, err := routes.Update_EmpbyStudentId(dto.StudId, dto.NewEnterprise, dto.NewWorkStartDate, dto.NewJobTitle)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	} else {
		res.Enterprise = dto.NewEnterprise
		res.WorkStartDate = dto.NewWorkStartDate
		res.JobTitle = dto.NewJobTitle
	}
	res.Code = code
	return res
}

func (a *App) Delete_EmpbyStudentId(dto dtos.DeleteEmployersDTO) models.Remove {
	var res models.Remove
	_, err := routes.Delete_EmpbyStudentId(dto.StudId)
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

func (a *App) Create_Group(dto dtos.CreateGroupDTO) models.OutGroupApi {
	var res models.OutGroupApi

	id, err := routes.Create_Group(dto.Course, dto.Groudates, dto.Speciality, dto.Number, dto.Semester)

	code := 200

	if err != nil {
		code = 500
		res.Error = err.Error()
	}

	res.Code = code

	res.Id = id
	res.Course = dto.Course
	res.Groudates = dto.Groudates
	res.Speciality = dto.Speciality
	res.Number = dto.Number
	res.Semester = dto.Semester

	return res
}

func (a *App) Update_GroupById(dto dtos.UpdateGroupDTO) models.OutGroupApi {
	var res models.OutGroupApi

	status, err := routes.Update_GroupById(dto.GroupId, dto.NewCourse, dto.NewGroudates, dto.NewSpeciality, dto.NewNumber, dto.NewSemester)

	code := 200

	if err != nil || status == false {
		code = 500
		res.Error = err.Error()
	}

	res.Code = code

	res.Id = dto.GroupId
	res.Course = dto.NewCourse
	res.Groudates = dto.NewGroudates
	res.Speciality = dto.NewSpeciality
	res.Number = dto.NewNumber
	res.Semester = dto.NewSemester

	return res
}

func (a *App) Delete_GroupById(dto dtos.DeleteGroupDTO) models.Remove {
	var res models.Remove
	_, err := routes.Delete_GroupById(dto.GroupId)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	return res
}

func (a *App) GetGroupId_GroupIdByInfo(dto dtos.GetGroupIdByInfo) models.GetId {
	var res models.GetId
	id, err := routes.GetGroupId_GroupIdByInfo(dto.Course, dto.Groudates, dto.Speciality, dto.Number, dto.Semester)
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

func (a *App) Inf_SubjectByGroupId(dto dtos.InfSubjectDTO) models.Inf_Subject {
	var info models.Inf_Subject
	subjects, err := routes.Inf_SubjectByGroupId(dto.GroupId)
	code := 200
	if err != nil {
		code = 500
		info.Error = err.Error()
	}
	info.Code = code
	info.Subject = subjects
	return info
}

func (a *App) Add_SubjectByGroupId(dto dtos.AddSubjectDTO) models.OutSubject {
	var res models.OutSubject

	id, err := routes.Add_SubjectByGroupId(dto.GroupId, dto.NewSubject)
	code := 200

	if err != nil {
		code = 500
		res.Error = err.Error()
	} else {
		res.Subject = dto.NewSubject
		res.GroupId = dto.GroupId
	}

	res.Code = code
	res.GroupId = id

	return res
}

func (a *App) Update_SubjectById(dto dtos.UpdateSubjectDTO) models.OutSubject {
	var res models.OutSubject

	status, err := routes.Update_SubjectById(dto.SubjectId, dto.NewSubject)
	code := 200

	if err != nil || !status {
		code = 500
		if err != nil {
			res.Error = err.Error()
		} else {
			res.Error = "не удалось обновить предмет"
		}
	} else {
		res.Subject = dto.NewSubject
	}

	res.Code = code
	return res
}

func (a *App) Delete_SubjectById(dto dtos.DeleteSubjectById) models.Remove {
	var res models.Remove
	_, err := routes.Delete_SubjectById(dto.SubjectId)
	code := 200
	if err != nil {
		code = 500
		res.Error = err.Error()
	}
	res.Code = code
	return res
}

func (a *App) Delete_AllSubjectByGroupId(dto dtos.DeleteAddSubjectByGroup) models.Remove {
	var res models.Remove
	_, err := routes.Delete_AllSubjectByGroupId(dto.GroupId)
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
