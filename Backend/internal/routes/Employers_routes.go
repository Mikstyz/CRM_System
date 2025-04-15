package routes

import (
	repo "CRM_System/Backend/internal/db/repository"
	"CRM_System/Backend/internal/modeles"
	"CRM_System/Backend/internal/utils"
	"fmt"
)

func GenerateFilledPDF(StudentId int, StudentName string, GroupId int) ([]byte, error) {
	InfEmp, err := repo.InfEmpById(StudentId)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении данных о трудоустройстве: %w", err)
	}

	subj, err := repo.InfDiscipAnDock(GroupId)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении списка предметов: %w", err)
	}

	data := modeles.PdfDoc{
		Name:          StudentName,
		Enterprise:    InfEmp.Enterprise,
		WorkStartDate: InfEmp.WorkStartDate,
		JobTitle:      InfEmp.JobTitle,
		SubjectArray:  subj,
	}

	PdfDock, err := utils.GenerateFilledPDF(data)
	if err != nil {
		return nil, fmt.Errorf("ошибка при генерации PDF: %w", err)
	}

	return PdfDock, nil
}

func Inf_EmpByStudentId(StudentId int) (modeles.EInfEmp, error) {
	return repo.InfEmpById(StudentId)
}

func Update_EmpbyStudentId(studId int, newEnterprise string, newWorkStartDate string, newJobTitle string) (bool, error) {
	result, err := repo.UpdateEmp(studId, newEnterprise, newWorkStartDate, newJobTitle)

	if err != nil {
		fmt.Printf("{routes}Ошибка при обновлении данных работодателя у студента %v", err)
		return false, err
	}

	fmt.Printf("{routes}Успешное обновление данных работодателя%v\n[ERORR: %v]", result, err)
	return result, nil
}

func Delete_EmpbyStudentId(studId int) (bool, error) {
	return repo.DelEmp(studId)
}
