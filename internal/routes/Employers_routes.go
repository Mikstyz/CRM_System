package routes

import (
	repo "CRM_System/internal/db/repository"
	"CRM_System/internal/modeles"
	"CRM_System/internal/utils"
	"fmt"
)

func GenerateFilledPDF(StudentId int, StudentName string, GroupId int) ([]byte, error) {
	InfEmp, err := repo.InfEmpById(StudentId)
	if err != nil {
		fmt.Printf("Ошибка при получении данных о трудоустройстве для студента с ID %d: %v\n", StudentId, err)
		return nil, fmt.Errorf("ошибка при получении данных о трудоустройстве: %w", err)
	}

	subj, err := repo.InfDiscipAnDock(GroupId)
	if err != nil {
		fmt.Printf("Ошибка при получении списка предметов для группы с ID %d: %v\n", GroupId, err)
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
		fmt.Printf("Ошибка при генерации PDF для студента с ID %d: %v\n", StudentId, err)
		return nil, fmt.Errorf("ошибка при генерации PDF: %w", err)
	}

	fmt.Printf("Успешное создание PDF для студента %s (ID: %d)\n", StudentName, StudentId)
	return PdfDock, nil
}

func Inf_EmpByStudentId(StudentId int) (modeles.EInfEmp, error) {
	result, err := repo.InfEmpById(StudentId)
	if err != nil {
		fmt.Printf("Ошибка при получении данных о трудоустройстве для студента с ID %d: %v\n", StudentId, err)
		return result, err
	}
	fmt.Printf("Получены данные о трудоустройстве для студента с ID %d\n", StudentId)
	return result, nil
}

func Update_EmpbyStudentId(studId int, newEnterprise string, newWorkStartDate string, newJobTitle string) (bool, error) {
	result, err := repo.UpdateEmp(studId, newEnterprise, newWorkStartDate, newJobTitle)

	if err != nil {
		fmt.Printf("{routes}Ошибка при обновлении данных работодателя у студента с ID %d: %v\n", studId, err)
		return false, err
	}

	fmt.Printf("{routes}Успешное обновление данных работодателя для студента с ID %d: %v\n", studId, result)
	return result, nil
}

func Delete_EmpbyStudentId(studId int) (bool, error) {
	result, err := repo.DelEmp(studId)

	if err != nil {
		fmt.Printf("Ошибка при удалении данных работодателя для студента с ID %d: %v\n", studId, err)
		return false, err
	}

	fmt.Printf("Успешное удаление данных работодателя для студента с ID %d\n", studId)
	return result, nil
}
