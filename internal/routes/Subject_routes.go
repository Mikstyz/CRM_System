package routes

import (
	repo "CRM_System/internal/db/repository"
	"fmt"
)

func Inf_SubjectByGroupId(groupId int) ([]string, error) {
	result, err := repo.InfDisciplinesByGroup(groupId)

	if err != nil {
		fmt.Printf("Ошибка при получении предметов группы с ID %d: %v\n", groupId, err)
		return nil, err
	}

	fmt.Printf("Получение предметов для группы с ID %d: успешное\n", groupId)
	return result, nil
}

func Add_SubjectByGroupId(groupId int, newSubject string) (int, error) {
	result, err := repo.AddDisciplinesInGroup(groupId, newSubject)

	if err != nil {
		fmt.Printf("Ошибка при добавлении предмета %s в группу с ID %d: %v\n", newSubject, groupId, err)
		return 0, err
	}

	fmt.Printf("Успешное добавление предмета %s в группу с ID %d\n", newSubject, groupId)
	return result, nil
}

func Update_SubjectById(subjectId int, newSubject string) (bool, error) {
	result, err := repo.UpdateDisciplinesInGroup(subjectId, newSubject)

	if err != nil {
		fmt.Printf("Ошибка при обновлении предмета с ID %d: %v\n", subjectId, err)
		return false, err
	}

	fmt.Printf("Успешное обновление предмета с ID %d на %s\n", subjectId, newSubject)
	return result, nil
}

func Delete_SubjectById(subjectId int) (bool, error) {
	result, err := repo.DeleteDisciplinesInGroup(subjectId)

	if err != nil {
		fmt.Printf("Ошибка при удалении предмета с ID %d: %v\n", subjectId, err)
		return false, err
	}

	fmt.Printf("Успешное удаление предмета с ID %d\n", subjectId)
	return result, nil
}

func Delete_AllSubjectByGroupId(groupID int) (bool, error) {
	result, err := repo.DeleteDisciplinesByGroupId(groupID)

	if err != nil {
		fmt.Printf("Ошибка при удалении всех предметов для группы с ID %d: %v\n", groupID, err)
		return false, err
	}

	fmt.Printf("Успешное удаление всех предметов для группы с ID %d\n", groupID)
	return result, nil
}
