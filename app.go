package main

import (
	dtos "CRM_System/internal/DTOs"
	models "CRM_System/internal/models"
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

// ----------------- Student ----------------
// ----------------------------------- Inf_AllStudent -----------------------------------
func (a *App) InfAllStudent(dto dtos.Inf_AllStudent) models.AppInf_AllStudent {
	students, err := routes.Inf_AllStudent()
	if err != nil {
		return models.AppInf_AllStudent{Code: 500, Error: err.Error()}
	}
	return models.AppInf_AllStudent{Code: 200, Students: students}
}

// ----------------------------------- Inf_StudentByID -----------------------------------
func (a *App) InfStudentByID(dto dtos.Inf_StudentByID) models.AppInf_StudentByID {
	student, err := routes.Inf_StudentByID(dto.StudentID)
	if err != nil {
		return models.AppInf_StudentByID{Code: 500, Error: err.Error()}
	}
	return models.AppInf_StudentByID{Code: 200, Student: student}
}

// ----------------------------------- Inf_StudentByGroup -----------------------------------
func (a *App) InfStudentByGroup(dto dtos.Inf_StudentByGroup) models.AppInf_StudentByGroup {
	students, err := routes.Inf_StudentByGroup(dto.Course, dto.Speciality, dto.GroupNum, dto.Semester)
	if err != nil {
		return models.AppInf_StudentByGroup{Code: 500, Error: err.Error()}
	}
	return models.AppInf_StudentByGroup{Code: 200, Students: students}
}

// ----------------------------------- Create_Student -----------------------------------
func (a *App) CreateStudent(dto dtos.Create_Student) models.AppCreate_Student {
	id, err := routes.Create_Student(dto.FullName, dto.Course, dto.Graduates, dto.Speciality, dto.GroupNum, dto.Semester)
	if err != nil {
		return models.AppCreate_Student{Code: 500, Error: err.Error()}
	}
	return models.AppCreate_Student{
		Code:       200,
		Id:         id,
		FullName:   dto.FullName,
		Course:     dto.Course,
		Graduates:  dto.Graduates,
		Speciality: dto.Speciality,
		GroupNum:   dto.GroupNum,
		Semester:   dto.Semester,
	}
}

// ----------------------------------- Update_StudentById -----------------------------------
func (a *App) UpdateStudentByID(dto dtos.Update_StudentById) models.AppUpdate_StudentById {
	status, err := routes.Update_StudentById(dto.StudId, dto.NewFullName, dto.NewCourse, dto.NewGraduates, dto.NewSpeciality, dto.NewNumber, dto.NewSemester)
	if err != nil || !status {
		return models.AppUpdate_StudentById{Code: 500, Error: err.Error()}
	}
	return models.AppUpdate_StudentById{
		Code:          200,
		NewId:         dto.StudId,
		NewFullName:   dto.NewFullName,
		NewCourse:     dto.NewCourse,
		NewGraduates:  dto.NewGraduates,
		NewSpeciality: dto.NewSpeciality,
		NewGroupNum:   dto.NewNumber,
		NewSemester:   dto.NewSemester,
	}
}

// ----------------------------------- Delete_Student -----------------------------------
func (a *App) DeleteStudent(dto dtos.Delete_Student) models.AppDelete_Student {
	status, err := routes.Delete_Student(dto.StudId)
	if err != nil || !status {
		return models.AppDelete_Student{Code: 500, Error: err.Error()}
	}
	return models.AppDelete_Student{Code: 200}
}

// ----------------- Employers --------------
// ----------------------------------- Inf_EmpByStudentId -----------------------------------
func (a *App) InfEmpByStudentID(dto dtos.Inf_EmpByStudentId) models.AppInf_EmpByStudentId {
	employers, err := routes.Inf_EmpByStudentId(dto.StudentId)
	if err != nil {
		return models.AppInf_EmpByStudentId{Code: 500, Error: err.Error()}
	}
	return models.AppInf_EmpByStudentId{Code: 200, Employers: employers}
}

// ----------------------------------- Create_Employers -----------------------------------
// func (a *App) Create_Employers(dto dtos.Create_Employers) models.AppCreate_Employers {
// 	var res models.AppCreate_Employers
// 	_, err := routes.Create_Employers()

// 	code := 200
// 	if err != nil {
// 		code = 500
// 		res.Error = err.Error()
// 	}
// 	res.Code = code
// 	return res
// }

// ----------------------------------- Update_EmpbyStudentId -----------------------------------
func (a *App) UpdateEmpByStudentID(dto dtos.Update_EmpbyStudentId) models.AppUpdate_EmpbyStudentId {
	status, err := routes.Update_EmpbyStudentId(dto.StudId, dto.NewEnterprise, dto.NewWorkStartDate, dto.NewJobTitle)
	if err != nil || !status {
		return models.AppUpdate_EmpbyStudentId{Code: 500, Error: err.Error()}
	}
	return models.AppUpdate_EmpbyStudentId{
		Code:             200,
		StudId:           dto.StudId,
		NewEnterprise:    dto.NewEnterprise,
		NewWorkStartDate: dto.NewWorkStartDate,
		NewJobTitle:      dto.NewJobTitle,
	}
}

// ----------------------------------- Delete_EmpbyStudentId -----------------------------------
func (a *App) DeleteEmpByStudentID(dto dtos.Delete_EmpbyStudentId) models.AppDelete_EmpbyStudentId {
	status, err := routes.Delete_EmpbyStudentId(dto.StudId)
	if err != nil || !status {
		return models.AppDelete_EmpbyStudentId{Code: 500, Error: err.Error()}
	}
	return models.AppDelete_EmpbyStudentId{Code: 200}
}

// ----------------- Group -------------------
// ----------------------------------- Inf_AllGroup -----------------------------------
// Inf_AllGroup получает список всех групп
func (a *App) InfAllGroup(dto dtos.InfAllGroupDTO) models.AppInf_AllGroup {
	groups, err := routes.Inf_AllGroup()
	if err != nil {
		return models.AppInf_AllGroup{Code: 500, Error: err.Error()}
	}
	return models.AppInf_AllGroup{Code: 200, Groups: groups}
}

// ----------------------------------- Create_Group -----------------------------------
// Create_Group создаёт новую группу
func (a *App) CreateGroup(dto dtos.CreateGroupDTO) models.AppCrtOutGroupApi {
	id, err := routes.Create_Group(dto.Course, dto.Graduates, dto.Speciality, dto.Number)
	if err != nil {
		return models.AppCrtOutGroupApi{Code: 500, Error: err.Error()}
	}

	groups := []models.EinfGroup{
		{
			Id:         id[0],
			Course:     dto.Course,
			Groudates:  dto.Graduates,
			Speciality: dto.Speciality,
			Number:     dto.Number,
			Semester:   1,
		},
		{
			Id:         id[1],
			Course:     dto.Course,
			Groudates:  dto.Graduates,
			Speciality: dto.Speciality,
			Number:     dto.Number,
			Semester:   2,
		},
	}

	return models.AppCrtOutGroupApi{Code: 200, Groups: groups}
}

// ----------------------------------- Update_GroupById -----------------------------------
// Update_GroupById обновляет группу по ID
func (a *App) UpdateGroupByID(dto dtos.UpdateGroupDTO) models.AppUpdateOutGroupApi {
	status, err := routes.Update_GroupById(dto.GroupId, dto.NewCourse, dto.NewGraduates, dto.NewSpeciality, dto.NewNumber)
	if err != nil || !status {
		errorMsg := "не удалось обновить группу"
		if err != nil {
			errorMsg = err.Error()
		}
		return models.AppUpdateOutGroupApi{Code: 500, Error: errorMsg}
	}
	return models.AppUpdateOutGroupApi{
		Code:       200,
		Id:         dto.GroupId,
		Course:     dto.NewCourse,
		Graduates:  dto.NewGraduates,
		Speciality: dto.NewSpeciality,
		Number:     dto.NewNumber,
	}
}

// ----------------------------------- Delete_GroupById -----------------------------------
// Delete_GroupById удаляет группу по данным
func (a *App) DeleteGroupByID(dto dtos.DeleteGroupDTO) models.AppRemove {
	_, err := routes.Delete_GroupById(dto.Course, dto.Graduates, dto.Speciality, dto.Number)
	if err != nil {
		return models.AppRemove{Code: 500, Error: err.Error()}
	}
	return models.AppRemove{Code: 200}
}

// ----------------------------------- GetGroupId_GroupIdByInfo -----------------------------------
// GetGroupId_GroupIdByInfo получает ID группы по данным
func (a *App) GetGroupIDByInfo(dto dtos.GetGroupIdByInfoDTO) models.AppGetId {
	id, err := routes.GetGroupId_GroupIdByInfo(dto.Course, dto.Graduates, dto.Speciality, dto.Number, dto.Semester)
	if err != nil {
		return models.AppGetId{Code: 500, Error: err.Error()}
	}
	return models.AppGetId{Code: 200, Id: id}
}

// ----------------------------------- DublicateGroupAllData -----------------------------------
// DublicateGroupAllData дублирует данные группы
func (a *App) DuplicateGroupAllData(dto dtos.DublicateGroupDTO) models.AppCrtOutGroupApi {
	id, err := routes.DublicateGroupAllData(dto.Course, dto.Graduates, dto.Speciality, dto.Number)
	if err != nil {
		return models.AppCrtOutGroupApi{Code: 500, Error: err.Error()}
	}

	groups := []models.EinfGroup{
		{
			Id:         id[0],
			Course:     dto.Course,
			Groudates:  dto.Graduates,
			Speciality: dto.Speciality,
			Number:     dto.Number,
			Semester:   1,
		},
		{
			Id:         id[1],
			Course:     dto.Course,
			Groudates:  dto.Graduates,
			Speciality: dto.Speciality,
			Number:     dto.Number,
			Semester:   2,
		},
	}

	return models.AppCrtOutGroupApi{Code: 200, Groups: groups}
}

// ----------------- Subject -----------------
// ----------------------------------- Inf_SubjectByGroupId -----------------------------------
// Inf_SubjectByGroupId получает список предметов для группы по ID
func (a *App) InfSubjectByGroupID(dto dtos.InfSubjectDTO) models.AppInf_Subject {
	subjects, err := routes.Inf_SubjectByGroupId(dto.GroupId)
	if err != nil {
		return models.AppInf_Subject{Code: 500, Error: err.Error()}
	}
	return models.AppInf_Subject{Code: 200, Subject: subjects}
}

// ----------------------------------- Inf_DisciplinesByGroupData -----------------------------------
// Inf_DisciplinesByGroupData получает предметы по данным группы
func (a *App) InfDisciplinesByGroupData(dto dtos.InfDisciplinesByGroupDataDTO) models.AppInf_Disciplines {
	disciplines, err := routes.Inf_DisciplinesByGroupData(dto.Speciality, dto.GroupNum, dto.Course, dto.Groudates)
	if err != nil {
		return models.AppInf_Disciplines{Code: 500, Error: err.Error()}
	}
	return models.AppInf_Disciplines{Code: 200, Disciplines: disciplines}
}

// ----------------------------------- Add_SubjectByGroupId -----------------------------------
// Add_SubjectByGroupId добавляет предмет для группы
func (a *App) AddSubjectByGroupID(dto dtos.AddSubjectDTO) models.AppOutSubject {
	id, err := routes.Add_SubjectByGroupId(dto.GroupId, dto.NewSubject)
	if err != nil {
		return models.AppOutSubject{Code: 500, Error: err.Error()}
	}
	return models.AppOutSubject{
		Code:    200,
		Id:      id,
		Subject: dto.NewSubject,
		GroupId: dto.GroupId,
	}
}

// ----------------------------------- Update_SubjectById -----------------------------------
// Update_SubjectById обновляет предмет по ID
func (a *App) UpdateSubjectByID(dto dtos.UpdateSubjectDTO) models.AppOutSubject {
	status, err := routes.Update_SubjectById(dto.SubjectId, dto.NewSubject)
	if err != nil || !status {
		errorMsg := "не удалось обновить предмет"
		if err != nil {
			errorMsg = err.Error()
		}
		return models.AppOutSubject{Code: 500, Error: errorMsg}
	}
	return models.AppOutSubject{
		Code:    200,
		Id:      dto.SubjectId,
		Subject: dto.NewSubject,
	}
}

// ----------------------------------- Delete_SubjectById -----------------------------------
// Delete_SubjectById удаляет предмет по ID
func (a *App) DeleteSubjectByID(dto dtos.DeleteSubjectDTO) models.AppRemove {
	_, err := routes.Delete_SubjectById(dto.SubjectId)
	if err != nil {
		return models.AppRemove{Code: 500, Error: err.Error()}
	}
	return models.AppRemove{Code: 200}
}

// ----------------------------------- Delete_AllSubjectByGroupId -----------------------------------
// Delete_AllSubjectByGroupId удаляет все предметы для группы
func (a *App) DeleteAllSubjectsByGroupID(dto dtos.DeleteAllSubjectByGroupDTO) models.AppRemove {
	_, err := routes.Delete_AllSubjectByGroupId(dto.GroupId)
	if err != nil {
		return models.AppRemove{Code: 500, Error: err.Error()}
	}
	return models.AppRemove{Code: 200}
}

// ----------------- PDF ----------------------
// ----------------------------------- GenerateFilledPDF -----------------------------------
func (a *App) GenerateFilledPDF(dockDATA models.GeneratePDF) models.AppPdfDock {
	file, err := routes.GenerateFilledPDF(dockDATA)
	if err != nil {
		return models.AppPdfDock{Code: 500, Error: err.Error()}
	}
	return models.AppPdfDock{Code: 200, File: file}
}
