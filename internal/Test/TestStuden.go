package Test

import (
	"CRM_System/internal/routes"
	"errors"
	"log"
)

func Test_InfAllStudent() (int, int, error) {
	var Ok, Bad int

	log.Println("[INFO] Получаем всех студентов...")
	students, err := routes.Inf_AllStudent()
	if err != nil {
		log.Printf("Ошибка при получении студентов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[INFO] Студенты успешно получены, найдено %d записей.", len(students))
	Ok++

	return Ok, Bad, nil
}

func Test_InfStudentByID() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, groduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupId, err := routes.Create_Group(course, groduates, speciality, groupNum)
	if err != nil {
		log.Printf("Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("Группа создана с ID %d", groupId)
	Ok++

	// Создаём студента
	log.Println("[INFO] Создаём тестового студента...")
	fullName := "Иван Иванов"
	studentID, err := routes.Create_Student(fullName, groupId)
	if err != nil {
		log.Printf("Ошибка при создании студента: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("Студент создан с ID=%d", studentID)
	Ok++

	// Получаем студента по ID
	log.Println("[INFO] Получаем студента по ID...")
	student, err := routes.Inf_StudentByID(studentID)
	if err != nil {
		log.Printf("Ошибка при получении студента с ID=%d: %v", studentID, err)
		Bad++
		return Ok, Bad, err
	}
	if student.FullName != fullName || student.GroupId != groupId {
		log.Printf("Ошибка: неверные данные студента с ID=%d, ожидалось имя=%s, GroupId=%d, получено=%v", studentID, fullName, groupId, student)
		Bad++
		return Ok, Bad, errors.New("неверные данные студента")
	}
	log.Println("[INFO] Студент успешно получен по ID.")
	Ok++

	// Удаляем студента
	log.Printf("[INFO] Удаляем студента с ID=%d...", studentID)
	ok, err := routes.Delete_Student(studentID)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении студента: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления студента")
	}
	log.Println("[INFO] Студент успешно удалён.")
	Ok++

	// Удаляем группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err = routes.Delete_GroupById(groupId)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[INFO] Тестовая группа успешно удалена.")
	Ok++

	return Ok, Bad, nil
}

func Test_CreateStudent() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, groduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupId, err := routes.Create_Group(course, groduates, speciality, groupNum)
	if err != nil {
		log.Printf("Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("Группа создана с ID %d", groupId)
	Ok++

	// Создаём студента
	log.Println("[INFO] Создаём тестового студента...")
	fullName := "Иван Иванов"
	studentID, err := routes.Create_Student(fullName, groupId)
	if err != nil {
		log.Printf("Ошибка при создании студента: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if studentID == 0 {
		log.Println("[ERROR] Ошибка: ID студента не должен быть 0")
		Bad++
		return Ok, Bad, errors.New("нулевой ID студента")
	}
	log.Printf("Студент создан с ID=%d", studentID)
	Ok++

	// Проверяем создание
	log.Println("[INFO] Проверяем создание студента...")
	student, err := routes.Inf_StudentByID(studentID)
	if err != nil {
		log.Printf("Ошибка при получении студента с ID=%d: %v", studentID, err)
		Bad++
		return Ok, Bad, err
	}
	if student.FullName != fullName || student.GroupId != groupId {
		log.Printf("Ошибка: неверные данные студента с ID=%d, ожидалось имя=%s, GroupId=%d, получено=%v", studentID, fullName, groupId, student)
		Bad++
		return Ok, Bad, errors.New("неверные данные студента")
	}
	log.Println("[INFO] Студент успешно создан и проверен.")
	Ok++

	// Удаляем студента
	log.Printf("[INFO] Удаляем студента с ID=%d...", studentID)
	ok, err := routes.Delete_Student(studentID)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении студента: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления студента")
	}
	log.Println("[INFO] Студент успешно удалён.")
	Ok++

	// Удаляем группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err = routes.Delete_GroupById(groupId)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[INFO] Тестовая группа успешно удалена.")
	Ok++

	return Ok, Bad, nil
}

func Test_UpdateStudentById() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, groduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupId, err := routes.Create_Group(course, groduates, speciality, groupNum)
	if err != nil {
		log.Printf("Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("Группа создана с ID %d", groupId)
	Ok++

	// Создаём вторую группу для обновления
	log.Println("[INFO] Создаём вторую тестовую группу...")
	newCourse, newGroduates, newSpeciality, newGroupNum := byte(2), byte(5), "Кибербезопасность", 102
	newGroupId, err := routes.Create_Group(newCourse, newGroduates, newSpeciality, newGroupNum)
	if err != nil {
		log.Printf("Ошибка при создании второй группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("Вторая группа создана с ID %d", newGroupId)
	Ok++

	// Создаём студента
	log.Println("[INFO] Создаём тестового студента...")
	fullName := "Иван Иванов"
	studentID, err := routes.Create_Student(fullName, groupId)
	if err != nil {
		log.Printf("Ошибка при создании студента: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("Студент создан с ID=%d", studentID)
	Ok++

	// Обновляем студента
	log.Printf("[INFO] Обновляем студента с ID=%d...", studentID)
	newFullName := "Пётр Петров"
	ok, err := routes.Update_StudentById(studentID, newFullName, newGroupId)
	if err != nil || !ok {
		log.Printf("Ошибка при обновлении студента: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка обновления студента")
	}
	log.Println("[INFO] Студент успешно обновлён.")
	Ok++

	// Проверяем обновление
	log.Println("[INFO] Проверяем обновление студента...")
	student, err := routes.Inf_StudentByID(studentID)
	if err != nil {
		log.Printf("Ошибка при получении студента с ID=%d: %v", studentID, err)
		Bad++
		return Ok, Bad, err
	}
	if student.FullName != newFullName || student.GroupId != newGroupId {
		log.Printf("Ошибка: неверные данные студента с ID=%d, ожидалось имя=%s, GroupId=%d, получено=%v", studentID, newFullName, newGroupId, student)
		Bad++
		return Ok, Bad, errors.New("неверные данные обновлённого студента")
	}
	log.Println("[INFO] Обновление студента подтверждено.")
	Ok++

	// Удаляем студента
	log.Printf("[INFO] Удаляем студента с ID=%d...", studentID)
	ok, err = routes.Delete_Student(studentID)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении студента: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления студента")
	}
	log.Println("[INFO] Студент успешно удалён.")
	Ok++

	// Удаляем группы
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err = routes.Delete_GroupById(groupId)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[INFO] Тестовая группа успешно удалена.")
	Ok++

	log.Println("[INFO] Удаляем вторую тестовую группу...")
	ok, err = routes.Delete_GroupById(newGroupId)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении второй тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления второй тестовой группы")
	}
	log.Println("[INFO] Вторая тестовая группа успешно удалена.")
	Ok++

	return Ok, Bad, nil
}

func Test_DeleteStudent() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, groduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupId, err := routes.Create_Group(course, groduates, speciality, groupNum)
	if err != nil {
		log.Printf("Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("Группа создана с ID %d", groupId)
	Ok++

	// Создаём студента
	log.Println("[INFO] Создаём тестового студента...")
	fullName := "Иван Иванов"
	studentID, err := routes.Create_Student(fullName, groupId)
	if err != nil {
		log.Printf("Ошибка при создании студента: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("Студент создан с ID=%d", studentID)
	Ok++

	// Удаляем студента
	log.Printf("[INFO] Удаляем студента с ID=%d...", studentID)
	ok, err := routes.Delete_Student(studentID)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении студента: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления студента")
	}
	log.Println("[INFO] Студент успешно удалён.")
	Ok++

	// Проверяем удаление
	log.Println("[INFO] Проверяем удаление студента...")
	_, err = routes.Inf_StudentByID(studentID)
	if err == nil {
		log.Println("[ERROR] Ошибка: студент должен быть удалён, но он всё ещё существует.")
		Bad++
		return Ok, Bad, errors.New("студент не был удалён")
	}
	log.Println("[INFO] Студент успешно удалён, тест пройден.")
	Ok++

	// Удаляем группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err = routes.Delete_GroupById(groupId)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[INFO] Тестовая группа успешно удалена.")
	Ok++

	return Ok, Bad, nil
}

func Test_InfStudentByGroup() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, groduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupId, err := routes.Create_Group(course, groduates, speciality, groupNum)
	if err != nil {
		log.Printf("Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("Группа создана с ID %d", groupId)
	Ok++

	// Создаём двух студентов
	log.Println("[INFO] Создаём тестовых студентов...")
	studentNames := []string{"Иван Иванов", "Пётр Петров"}
	studentIDs := make([]int, 0, len(studentNames))
	for _, fullName := range studentNames {
		studentID, err := routes.Create_Student(fullName, groupId)
		if err != nil {
			log.Printf("Ошибка при создании студента %s: %v", fullName, err)
			Bad++
			return Ok, Bad, err
		}
		log.Printf("Студент %s создан с ID=%d", fullName, studentID)
		Ok++
		studentIDs = append(studentIDs, studentID)
	}

	// Получаем студентов по группе
	log.Println("[INFO] Получаем студентов по группе...")
	students, err := routes.Inf_StudentByGroup(groupId)
	if err != nil {
		log.Printf("Ошибка при получении студентов группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if len(students) != len(studentNames) {
		log.Printf("Ошибка: ожидалось %d студентов, получено %d", len(studentNames), len(students))
		Bad++
		return Ok, Bad, errors.New("неверное количество студентов")
	}
	for i, student := range students {
		if student.FullName != studentNames[i] || student.GroupId != groupId {
			log.Printf("Ошибка: неверные данные студента, ожидалось имя=%s, GroupId=%d, получено=%v", studentNames[i], groupId, student)
			Bad++
			return Ok, Bad, errors.New("неверные данные студента")
		}
	}
	log.Printf("[INFO] Студенты группы успешно получены, найдено %d записей.", len(students))
	Ok++

	// Удаляем студентов
	for _, studentID := range studentIDs {
		log.Printf("[INFO] Удаляем студента с ID=%d...", studentID)
		ok, err := routes.Delete_Student(studentID)
		if err != nil || !ok {
			log.Printf("Ошибка при удалении студента: %v", err)
			Bad++
			return Ok, Bad, errors.New("ошибка удаления студента")
		}
		log.Println("[INFO] Студент успешно удалён.")
		Ok++
	}

	// Удаляем группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err := routes.Delete_GroupById(groupId)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[INFO] Тестовая группа успешно удалена.")
	Ok++

	return Ok, Bad, nil
}

func TestStudentALL() (int, int) {
	var Ok, Bad int

	// Тестирование Inf_AllStudent
	ok, bad, err := Test_InfAllStudent()
	if err != nil {
		log.Fatal("Ошибка в Test_InfAllStudent: ", err)
	}
	logResults("InfAllStudent", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование Inf_StudentByID
	ok, bad, err = Test_InfStudentByID()
	if err != nil {
		log.Fatal("Ошибка в Test_InfStudentByID: ", err)
	}
	logResults("InfStudentByID", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование Create_Student
	ok, bad, err = Test_CreateStudent()
	if err != nil {
		log.Fatal("Ошибка в Test_CreateStudent: ", err)
	}
	logResults("CreateStudent", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование Update_StudentById
	ok, bad, err = Test_UpdateStudentById()
	if err != nil {
		log.Fatal("Ошибка в Test_UpdateStudentById: ", err)
	}
	logResults("UpdateStudentById", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование Delete_Student
	ok, bad, err = Test_DeleteStudent()
	if err != nil {
		log.Fatal("Ошибка в Test_DeleteStudent: ", err)
	}
	logResults("DeleteStudent", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование Inf_StudentByGroup
	ok, bad, err = Test_InfStudentByGroup()
	if err != nil {
		log.Fatal("Ошибка в Test_InfStudentByGroup: ", err)
	}
	logResults("InfStudentByGroup", ok, bad, err)
	Ok += ok
	Bad += bad

	// Итоговый вывод
	log.Println("\n[SUMMARY] Итоговые результаты:")
	log.Printf("TestStudentALL [%d\\%d]\nok: %d\nbad: %d\n", Ok, Bad, Ok, Bad)

	return Ok, Bad
}
