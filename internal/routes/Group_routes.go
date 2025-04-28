package routes

import (
	repo "CRM_System/internal/db/repository"
	"CRM_System/internal/models"
	"fmt"
)

func Inf_AllGroup() ([]models.EinfGroup, error) {
	result, err := repo.InfAllGrp()

	if err != nil {
		fmt.Printf("Ошибка при получении всех групп: %v\n", err)
		return nil, err
	}

	fmt.Println("Получение всех групп: успешное")
	return result, nil
}

func Create_Group(course byte, groduates byte, speciality string, groupNum int) (int, error) {
	Id, err := repo.CrtGrp(course, groduates, speciality, groupNum)

	if err != nil {
		fmt.Printf("Ошибка при создании группы: %v\n", err)
		return 0, err
	} else {
		fmt.Println("Успешное создание группы")
	}

	return Id, nil
}

func Update_GroupById(groupId int, newCourse byte, newGroduates byte, newSpeciality string, newGroupNum int) (bool, error) {
	result, err := repo.UpdateGrp(groupId, newCourse, newGroduates, newSpeciality, newGroupNum)

	if err != nil {
		fmt.Printf("Ошибка при обновлении группы с ID %d: %v\n", groupId, err)
		return false, err
	} else {
		fmt.Printf("Успешное обновление группы с ID %d\n", groupId)
	}

	return result, nil
}

func Delete_GroupById(course byte, graduates byte, speciality string, groupNum int) (bool, error) {
	result, err := repo.DelGrp(course, graduates, speciality, groupNum)

	if err != nil {
		fmt.Printf("Ошибка при удалении группы %v\n", err)
		return false, err
	} else {
		fmt.Printf("Успешное удаление группы")
	}

	return result, nil
}

func GetGroupId_GroupIdByInfo(course byte, groduates byte, speciality string, groupNum int) (int, error) {
	result, err := repo.GetGroupIDByParams(course, groduates, speciality, groupNum)

	if err != nil {
		fmt.Printf("Ошибка при получении ID группы с параметрами: course=%d, groduates=%d, speciality=%s, groupNum=%d: %v\n", course, groduates, speciality, groupNum, err)
		return 0, err
	} else {
		fmt.Println("Успешное получение ID группы")
	}

	return result, nil
}

func DublicateGroupAllData(Course byte, Groduates byte, Speciality string, GroupNum int) (int, error) { 
	GroupIdOne, err := repo.GetGroupIDByParams(Course, Groduates, Speciality, GroupNum)
	if err != nil {
		return 0, err
	}

	NewNumber, err := repo.MaxNumberByParams(Course, Groduates, Speciality)
	if err != nil {
		return 0, err
	}

	newGroupIds, err := repo.CrtGrp(Course, Groduates, Speciality, NewNumber+1)
	if err != nil {
		return 0, err
	}

	err = repo.CopyDisciplinesBetweenGroups(GroupIdOne, newGroupIds)
	if err != nil {
		return 0, err
	}

	return newGroupIds, nil
}
