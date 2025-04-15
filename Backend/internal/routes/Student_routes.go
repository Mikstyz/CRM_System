package routes

import (
	repo "CRM_System/Backend/internal/db/repository"
	"CRM_System/Backend/internal/modeles"
	"fmt"
)

func Inf_AllStudent() ([]modeles.Student, error) {
	result, err := repo.InfStdByGroup()

	if err != nil {
		fmt.Printf("Ошибка при получении информации о студентах %v", err)

		return nil, err
	}

	return result, nil
}

func Inf_StudentByID(studentID int) (*modeles.Student, error) {
	result, err := repo.GetStudentByID(studentID)

	if err != nil {
		fmt.Printf("Ошибка при получении информации о студенте %v", err)

		return nil, err
	}

	return result, nil
}

func Create_Student(fullName string, well byte, gClass byte, speciality string, groupNum int, semester byte) (int, error) {
	result, err := repo.CreateStudentWithEmptyEmployment(fullName, well, gClass, speciality, groupNum, semester)

	if err != nil {
		fmt.Printf("Ошибка при создании студента %v", err)

		return 0, err
	}

	return result, nil
}

func Update_StudentById(studId int, newFullName string, newWell byte, newClass byte, newSpeciality string, newGroupNum int, newSemester byte) (bool, error) {
	result, err := repo.UpdateStd(studId, newFullName, newWell, newClass, newSpeciality, newGroupNum, newSemester)

	if err != nil {
		fmt.Printf("Ошибка при обновлении данных студента %v", err)

		return false, err
	}

	return result, nil
}

func Delete_Student(studId int) (bool, error) {
	result, err := repo.DelStd(studId)

	if err != nil {
		fmt.Printf("Ошибка при удалении студента %v", err)

		return false, err
	}

	return result, nil
}
