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

func Create_Group(course byte, groduates byte, speciality string, groupNum int) ([2]int, error) {
	result, err := repo.CrtGrp(course, groduates, speciality, groupNum)

	if err != nil {
		fmt.Printf("Ошибка при создании группы: %v\n", err)
		return [2]int{0, 0}, err
	} else {
		fmt.Println("Успешное создание группы")
	}

	return result, nil
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
		fmt.Printf("Ошибка при удалении группы с ID %d: %v\n", err)
		return false, err
	} else {
		fmt.Printf("Успешное удаление группы")
	}

	return result, nil
}

func GetGroupId_GroupIdByInfo(course byte, groduates byte, speciality string, groupNum int, semester byte) (int, error) {
	result, err := repo.GetGroupIDByParams(course, groduates, speciality, groupNum, semester)

	if err != nil {
		fmt.Printf("Ошибка при получении ID группы с параметрами: course=%d, groduates=%d, speciality=%s, groupNum=%d, semester=%d: %v\n", course, groduates, speciality, groupNum, semester, err)
		return 0, err
	} else {
		fmt.Println("Успешное получение ID группы")
	}

	return result, nil
}

func DublicateGroupAllData(Course byte, Groduates byte, Speciality string, GroupNum int) ([2]int, error) {
	GroupIdOne, err := repo.GetGroupIDByParams(Course, Groduates, Speciality, GroupNum, 1)
	if err != nil {
		return [2]int{0, 0}, err
	}

	GroupIdTwo, err := repo.GetGroupIDByParams(Course, Groduates, Speciality, GroupNum, 2)
	if err != nil {
		return [2]int{0, 0}, err
	}

	newGroupIds, err := repo.CrtGrp(Course, Groduates, Speciality, GroupNum+1)
	if err != nil {
		return [2]int{0, 0}, err
	}

	err = repo.CopyDisciplinesBetweenGroups(GroupIdOne, newGroupIds[0])
	if err != nil {
		return [2]int{0, 0}, err
	}

	err = repo.CopyDisciplinesBetweenGroups(GroupIdTwo, newGroupIds[1])
	if err != nil {
		return [2]int{0, 0}, err
	}

	return newGroupIds, nil
}
