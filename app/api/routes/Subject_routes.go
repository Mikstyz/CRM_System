package routes

import (
	repo "CRM_System/app/storage/db/repository"
	"CRM_System/app/storage/models"
	"fmt"
)

// ----------------------information--------------------------
func Inf_DisciplinesByGroupId(groupId int, Semester byte) ([]string, error) {
	result, err := repo.FetchDisciplinesByGroup(groupId, Semester)

	if err != nil {
		fmt.Printf("[Routes][Disciplines] - Ошибка при получении предметов группы с ID %d: %v\n", groupId, err)
		return nil, err
	}

	fmt.Printf("[Routes][Disciplines] - Получение предметов для группы с ID %d: успешное\n", groupId)
	return result, nil
}

func Inf_DisciplinesByGroupData(Speciality string, GroupNum int, Course int, Groudates int) (models.DisciplinesBySemester, error) {
	result, err := repo.FetchDisciplinesByGroupData(Speciality, GroupNum, Course, Groudates)

	if err != nil {
		fmt.Printf("[Routes][Disciplines] - Ошибка при получении предметов %v", err)
		return result, err
	}

	fmt.Printf("[Routes][Disciplines] - Успешное получение предметов для группы")
	return result, nil
}

// ----------------------Manager--------------------------
func Add_DisciplinesByGroupId(groupId int, newSubject string, Semester byte) (int, error) {
	result, err := repo.AddDisciplinesInGroup(groupId, newSubject, Semester)

	if err != nil {
		fmt.Printf("[Routes][Disciplines] - Ошибка при добавлении предмета %s в группу с ID %d: %v\n", newSubject, groupId, err)
		return 0, err
	}

	fmt.Printf("[Routes][Disciplines] - Успешное добавление предмета %s в группу с ID %d\n", newSubject, groupId)
	return result, nil
}

func Update_DisciplinesById(subjectId int, newSubject string) (bool, error) {
	result, err := repo.UpdateDisciplinesInGroup(subjectId, newSubject)

	if err != nil {
		fmt.Printf("[Routes][Disciplines] - Ошибка при обновлении предмета с ID %d: %v\n", subjectId, err)
		return false, err
	}

	fmt.Printf("[Routes][Disciplines] - Успешное обновление предмета с ID %d на %s\n", subjectId, newSubject)
	return result, nil
}

func Delete_DisciplinesById(subjectId int) (bool, error) {
	result, err := repo.DeleteDisciplinesInGroup(subjectId)

	if err != nil {
		fmt.Printf("[Routes][Disciplines] - Ошибка при удалении предмета с ID %d: %v\n", subjectId, err)
		return false, err
	}

	fmt.Printf("[Routes][Disciplines] - Успешное удаление предмета с ID %d\n", subjectId)
	return result, nil
}

func Delete_AllDisciplinesByGroupId(groupID int) (bool, error) {
	result, err := repo.DeleteDisciplinesByGroupId(groupID)

	if err != nil {
		fmt.Printf("[Routes][Disciplines] - Ошибка при удалении всех предметов для группы с ID %d: %v\n", groupID, err)
		return false, err
	}

	fmt.Printf("[Routes][Disciplines] - Успешное удаление всех предметов для группы с ID %d\n", groupID)
	return result, nil
}
