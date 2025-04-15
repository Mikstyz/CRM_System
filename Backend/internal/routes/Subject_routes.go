package routes

import (
	repo "CRM_System/Backend/internal/db/repository"
	"fmt"
)

func Inf_SubjectByGroupId(groupId int) ([]string, error) {
	result, err := repo.InfDisciplinesInGroup(groupId)

	if err != nil {
		fmt.Printf("Ошибка при получении преметодов группы %v", err)
		return nil, err
	}

	return result, nil
}

func Add_SubjectByGroupId(groupId int, newSubject string) (int, error) {
	result, err := repo.AddDisciplinesInGroup(groupId, newSubject)

	if err != nil {
		fmt.Printf("Ошибка при получении преметодов группы %v", err)
		return 0, err
	}

	return result, nil
}

func Update_SubjectById(subjectId int, newSubject string) (bool, error) {
	result, err := repo.UpdateDisciplinesInGroup(subjectId, newSubject)

	if err != nil {
		fmt.Printf("Ошибка при получении преметодов группы %v", err)
		return false, err
	}

	return result, nil
}

func Delete_SubjectById(subjectId int) (bool, error) {
	result, err := repo.DeleteDisciplinesInGroup(subjectId)

	if err != nil {
		fmt.Printf("Ошибка при получении преметодов группы %v", err)
		return false, err
	}

	return result, nil
}

func Delete_AllSubjectByGroupId(groupID int) (bool, error) {
	return repo.DeleteDisciplinesByGroupId(groupID)
}
