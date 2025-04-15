package routes

import (
	repo "CRM_System/Backend/internal/db/repository"
	"CRM_System/Backend/internal/modeles"
	"fmt"
)

func Inf_AllGroup() ([]modeles.EinfGroup, error) {
	return repo.InfAllGrp()
}

func Create_Group(well byte, gClass byte, speciality string, groupNum int, semester byte) (int, error) {
	result, err := repo.CrtGrp(well, gClass, speciality, groupNum, semester)

	if err != nil {
		fmt.Printf("Ошибка при создании группы %v\n", err)
		return 0, err
	} else {
		fmt.Print("Успешное создание группы")
	}

	return result, nil
}

func Update_GroupById(groupId int, newWell byte, newGClass byte, newSpeciality string, newGroupNum int, newSemester byte) (bool, error) {
	result, err := repo.UpdateGrp(groupId, newWell, newGClass, newSpeciality, newGroupNum, newSemester)

	if err != nil {
		fmt.Printf("Ошибка при обновлении группы %v\n", err)
		return false, err
	} else {
		fmt.Print("Успешное обновление группы")
	}

	return result, nil
}

func Delete_GroupById(groupId int) (bool, error) {
	result, err := repo.DelGrp(groupId)

	if err != nil {
		fmt.Printf("Ошибка при удалении группы %v\n", err)
		return false, err
	} else {
		fmt.Print("Успешное удаление группы")
	}

	return result, nil
}

func GetGroupId_GroupIdByInfo(well byte, gClass byte, speciality string, groupNum int, semester byte) (int, error) {
	result, err := repo.GetGroupIDByParams(well, gClass, speciality, groupNum, semester)

	if err != nil {
		fmt.Printf("Ошибка при получении id группы %v\n", err)
		return 0, err
	} else {
		fmt.Print("Успешное получение id группы")
	}

	return result, nil
}
