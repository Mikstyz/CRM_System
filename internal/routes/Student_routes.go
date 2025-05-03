package routes

import (
	repo "CRM_System/internal/db/repository"
	"CRM_System/internal/models"
	"fmt"
)

func Inf_AllStudent() ([]models.Student, error) {
	result, err := repo.InfStdByGroup()

	if err != nil {
		fmt.Printf("Ошибка при получении информации о студентах: %v\n", err)
		return nil, err
	}

	fmt.Println("Получение информации о всех студентах: успешное")
	return result, nil
}

func Inf_StudentByID(studentID int) (models.Student, error) {
	result, err := repo.GetStudentByID(studentID)

	if err != nil {
		fmt.Printf("Ошибка при получении информации о студенте с ID %d: %v\n", studentID, err)
		return models.Student{}, err
	}

	fmt.Printf("Получение информации о студенте с ID %d: успешное\n", studentID)
	return result, nil
}

func Inf_StudentByGroup(groupId int) ([]models.Student, error) {
	result, err := repo.GetStudentByGroup(groupId)

	if err != nil {
		fmt.Println("Ошибка при получении студентов группы")
		return nil, err
	}

	fmt.Println("Успешное получение студентов группы")
	return result, nil
}

func Create_Student(fullName string, groupId int) (int, error) {
	result, err := repo.CreateStudentWithEmptyEmployment(fullName, groupId)

	if err != nil {
		fmt.Printf("Ошибка при создании студента %v: %v\n", fullName, err)
		return 0, err
	}

	fmt.Printf("Успешное создание студента %s\n", fullName)
	return result, nil
}

func Update_StudentById(studId int, newFullName string, newGroupId int) (bool, error) {
	result, err := repo.UpdateStd(studId, newFullName, newGroupId)

	if err != nil {
		fmt.Printf("Ошибка при обновлении данных студента с ID %d: %v\n", studId, err)
		return false, err
	}

	fmt.Printf("Успешное обновление данных студента с ID %d\n", studId)
	return result, nil
}

func Delete_Student(studId int) (bool, error) {
	result, err := repo.DelStd(studId)

	if err != nil {
		fmt.Printf("Ошибка при удалении студента с ID %d: %v\n", studId, err)
		return false, err
	}

	fmt.Printf("Успешное удаление студента с ID %d\n", studId)
	return result, nil
}
