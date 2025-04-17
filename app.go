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

func Inf_AllStudent() ([]modeles.Student, error) {
	return routes.Inf_AllStudent()
}

func Inf_StudentByID(studentID int) (*modeles.Student, error) {
	return routes.Inf_StudentByID(studentID)
}

func Create_Student(fullName string, well byte, gClass byte, speciality string, groupNum int, semester byte) (int, error) {
	return routes.Create_Student(fullName, well, gClass, speciality, groupNum, semester)
}

func Update_StudentById(studId int, newFullName string, newWell byte, newClass byte, newSpeciality string, newGroupNum int, newSemester byte) (bool, error) {
	return routes.Update_StudentById(studId, newFullName, newWell, newClass, newSpeciality, newGroupNum, newSemester)
}

func Delete_Student(studId int) (bool, error) {
	return routes.Delete_Student(studId)
}

//----------------- Employers --------------

func Inf_EmpByStudentId(StudentId int) (modeles.EInfEmp, error) {
	return routes.Inf_EmpByStudentId(StudentId)
}

func Update_EmpbyStudentId(studId int, newEnterprise string, newWorkStartDate string, newJobTitle string) (bool, error) {
	return routes.Update_EmpbyStudentId(studId, newEnterprise, newWorkStartDate, newJobTitle)
}

func Delete_EmpbyStudentId(studId int) (bool, error) {
	return routes.Delete_EmpbyStudentId(studId)
}

//----------------- Group -------------------

func Inf_AllGroup() ([]modeles.EinfGroup, error) {
	return routes.Inf_AllGroup()
}

func Create_Group(well byte, gClass byte, speciality string, groupNum int, semester byte) (int, error) {
	return routes.Create_Group(well, gClass, speciality, groupNum, semester)
}

func Update_GroupById(groupId int, newWell byte, newGClass byte, newSpeciality string, newGroupNum int, newSemester byte) (bool, error) {
	return routes.Update_GroupById(groupId, newWell, newGClass, newSpeciality, newGroupNum, newSemester)
}

func Delete_GroupById(groupId int) (bool, error) {
	return routes.Delete_GroupById(groupId)
}

func GetGroupId_GroupIdByInfo(well byte, gClass byte, speciality string, groupNum int, semester byte) (int, error) {
	return routes.GetGroupId_GroupIdByInfo(well, gClass, speciality, groupNum, semester)
}

//----------------- Subject -----------------

func Inf_SubjectByGroupId(groupId int) ([]string, error) {
	return routes.Inf_SubjectByGroupId(groupId)
}

func Add_SubjectByGroupId(groupId int, newSubject string) (int, error) {
	return routes.Add_SubjectByGroupId(groupId, newSubject)
}

func Update_SubjectById(subjectId int, newSubject string) (bool, error) {
	return routes.Update_SubjectById(subjectId, newSubject)
}

func Delete_SubjectById(subjectId int) (bool, error) {
	return routes.Delete_SubjectById(subjectId)
}

func Delete_AllSubjectByGroupId(groupID int) (bool, error) {
	return routes.Delete_AllSubjectByGroupId(groupID)
}

//----------------- PDF ----------------------

func GenerateFilledPDF(StudentId int, StudentName string, GroupId int) ([]byte, error) {
	return routes.GenerateFilledPDF(StudentId, StudentName, GroupId)
}
