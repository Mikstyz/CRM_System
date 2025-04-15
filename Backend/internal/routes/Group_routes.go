package routes

import (
	repo "CRM_System/Backend/internal/db/repository"
	"CRM_System/Backend/internal/modeles"
	"fmt"
)

func Inf_AllGroup() ([]modeles.EinfGroup, error) {
	result, err := repo.InfAllGrp()

	if err != nil {
		fmt.Printf("Ошибка при получении всех групп: %v\n", err)
		return nil, err
	}

	fmt.Println("Получение всех групп: успешное")
	return result, nil
}

func Create_Group(well byte, gClass byte, speciality string, groupNum int, semester byte) (int, error) {
	result, err := repo.CrtGrp(well, gClass, speciality, groupNum, semester)

	if err != nil {
		fmt.Printf("Ошибка при создании группы: %v\n", err)
		return 0, err
	} else {
		fmt.Println("Успешное создание группы")
	}

	return result, nil
}

func Update_GroupById(groupId int, newWell byte, newGClass byte, newSpeciality string, newGroupNum int, newSemester byte) (bool, error) {
	result, err := repo.UpdateGrp(groupId, newWell, newGClass, newSpeciality, newGroupNum, newSemester)

	if err != nil {
		fmt.Printf("Ошибка при обновлении группы с ID %d: %v\n", groupId, err)
		return false, err
	} else {
		fmt.Printf("Успешное обновление группы с ID %d\n", groupId)
	}

	return result, nil
}

func Delete_GroupById(groupId int) (bool, error) {
	result, err := repo.DelGrp(groupId)

	if err != nil {
		fmt.Printf("Ошибка при удалении группы с ID %d: %v\n", groupId, err)
		return false, err
	} else {
		fmt.Printf("Успешное удаление группы с ID %d\n", groupId)
	}

	return result, nil
}

func GetGroupId_GroupIdByInfo(well byte, gClass byte, speciality string, groupNum int, semester byte) (int, error) {
	result, err := repo.GetGroupIDByParams(well, gClass, speciality, groupNum, semester)

	if err != nil {
		fmt.Printf("Ошибка при получении ID группы с параметрами: well=%d, gClass=%d, speciality=%s, groupNum=%d, semester=%d: %v\n", well, gClass, speciality, groupNum, semester, err)
		return 0, err
	} else {
		fmt.Println("Успешное получение ID группы")
	}

	return result, nil
}
