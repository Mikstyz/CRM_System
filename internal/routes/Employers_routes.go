package routes

import (
	repo "CRM_System/internal/db/repository"
	"CRM_System/internal/models"
	"CRM_System/internal/utils"
	"errors"
	"fmt"
	"strings"
)

// поменять входные данные на имя студента и Id группы
// поменять входные данные на все поля (группы) (инфы про место работы) (имя студента)
func GenerateFilledPDF(dataPdf models.GeneratePDF) ([]byte, error) {

	// Проверка на пустые обязательные поля
	if strings.TrimSpace(dataPdf.StudentName) == "" ||
		strings.TrimSpace(dataPdf.Enterprise) == "" ||
		strings.TrimSpace(dataPdf.WorkStartDate) == "" ||
		strings.TrimSpace(dataPdf.JobTitle) == "" {

		return nil, errors.New("обязательные поля не должны быть пустыми")
	}

	// Получаем ID студента по имени и данным группы
	StudId, err := repo.GetIdByNameByGroup(dataPdf.StudentName, dataPdf.Course, dataPdf.Speciality, dataPdf.Groduates, dataPdf.Number)
	if err != nil {
		// Логирование ошибки
		return nil, fmt.Errorf("не удалось получить ID студента: %w", err)
	}

	// Обновляем информацию о предприятии по ID студента
	status, err := Update_EmpbyStudentId(StudId, dataPdf.Enterprise, dataPdf.WorkStartDate, dataPdf.JobTitle)
	if err != nil {
		// Логирование ошибки обновления
		return nil, fmt.Errorf("ошибка при обновлении информации о студенте: %w", err)
	}
	if !status {
		// Когда статус обновления ложный, мы возвращаем ошибку с контекстом
		return nil, fmt.Errorf("не удалось обновить информацию о студенте (статус обновления: %v)", status)
	}

	// Формируем данные для заполнения PDF
	data := models.PdfDoc{
		Name:          dataPdf.StudentName,
		Enterprise:    dataPdf.Enterprise,
		WorkStartDate: dataPdf.WorkStartDate,
		JobTitle:      dataPdf.JobTitle,
	}

	// Генерация заполненного PDF
	PdfDock, err := utils.GenerateFilledPDF(data)
	if err != nil {
		// Логирование ошибки генерации PDF
		return nil, fmt.Errorf("ошибка при генерации PDF для студента %s: %w", dataPdf.StudentName, err)
	}

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

func Create_Employers(studId int, enterprise string, workStartDate string, jobTitle string) {
	repo.CrteEmp(studId, enterprise, workStartDate, jobTitle)
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
