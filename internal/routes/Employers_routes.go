package routes

import (
	repo "CRM_System/internal/db/repository"
	"CRM_System/internal/models"
	"CRM_System/internal/utils"
	"fmt"
)

// поменять входные данные на имя студента и Id группы
// поменять входные данные на все поля (группы) (инфы про место работы) (имя студента)
func GenerateFilledPDF(dataPdf models.GeneratePDF) ([]byte, error) {
	data := models.PdfDoc{
		Name:          dataPdf.StudentName,
		Enterprise:    dataPdf.Enterprise,
		WorkStartDate: dataPdf.WorkStartDate,
		JobTitle:      dataPdf.JobTitle,
	}

	//проверка на null

	PdfDock, err := utils.GenerateFilledPDF(data)
	if err != nil {
		fmt.Printf("Ошибка при генерации PDF для студента %s: %v\n", dataPdf.StudentName, err)
		return nil, fmt.Errorf("ошибка при генерации PDF: %w", err)
	}

	fmt.Printf("Успешное создание PDF для студента %s (ID: %d)\n", dataPdf.StudentName)
	return PdfDock, nil
}

func Inf_EmpByStudentId(StudentId int) (models.EInfEmp, error) {
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
