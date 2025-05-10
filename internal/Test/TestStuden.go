package Test

import (
	"CRM_System/internal/models"
	"CRM_System/internal/routes"
	"errors"
	"log"
)

// Хелпер для создания тестовой группы
func createTestGroup(course, graduates byte, speciality string, groupNum int) (int, error) {
	groupId, err := routes.Create_Group(course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("Ошибка при создании группы: %v", err)
		return 0, err
	}
	log.Printf("Группа создана с ID %d", groupId)
	return groupId, nil
}

// Хелпер для создания тестового студента
func createTestStudent(fullName string, groupId int, enterprise, workStartDate, jobTitle string) (int, error) {
	studentID, err := routes.Create_Student(fullName, groupId, enterprise, workStartDate, jobTitle)
	if err != nil {
		log.Printf("Ошибка при создании студента: %v", err)
		return 0, err
	}
	log.Printf("Студент создан с ID=%d", studentID)
	return studentID, nil
}

// Хелпер для проверки данных студента
func checkStudentData(student models.Student, expectedFullName string, expectedGroupId int, expectedEnterprise, expectedWorkStartDate, expectedJobTitle string) error {
	if student.FullName != expectedFullName ||
		student.GroupId != expectedGroupId ||
		student.Enterprise == nil || *student.Enterprise != expectedEnterprise ||
		student.WorkStartDate == nil || *student.WorkStartDate != expectedWorkStartDate ||
		student.JobTitle == nil || *student.JobTitle != expectedJobTitle {
		return errors.New("неверные данные студента")
	}
	return nil
}

// Хелпер для удаления студента
func deleteTestStudent(studentID int) error {
	ok, err := routes.Delete_Student(studentID)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении студента: %v", err)
		return errors.New("ошибка удаления студента")
	}
	log.Println("[INFO] Студент успешно удалён.")
	return nil
}

// Хелпер для удаления группы
func deleteTestGroup(groupId int) error {
	ok, err := routes.Delete_GroupById(groupId)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении группы: %v", err)
		return errors.New("ошибка удаления группы")
	}
	log.Println("[INFO] Группа успешно удалена.")
	return nil
}

// Хелпер для логирования результатов
func logTestResult(testName string, ok, bad int, err error) {
	log.Printf("[RESULT] %s: ok=%d, bad=%d", testName, ok, bad)
	if err != nil {
		log.Printf("[ERROR] %s: %v", testName, err)
	}
}

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

	// Создаём группу
	log.Println("[INFO] Создаём тестовую группу...")
	groupId, err := createTestGroup(1, 4, "Информатика", 101)
	if err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++

	// Создаём студента
	log.Println("[INFO] Создаём тестового студента...")
	fullName, enterprise, workStartDate, jobTitle := "Иван Иванов", "Tech Corp", "2023-06-01", "Developer"
	studentID, err := createTestStudent(fullName, groupId, enterprise, workStartDate, jobTitle)
	if err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++

	// Получаем студента по ID
	log.Println("[INFO] Получаем студента по ID...")
	student, err := routes.Inf_StudentByID(studentID)
	if err != nil {
		log.Printf("Ошибка при получении студента с ID=%d: %v", studentID, err)
		Bad++
		return Ok, Bad, err
	}
	if err := checkStudentData(student, fullName, groupId, enterprise, workStartDate, jobTitle); err != nil {
		log.Printf("Ошибка: %v, получено=%v", err, student)
		Bad++
		return Ok, Bad, err
	}
	log.Println("[INFO] Студент успешно получен по ID.")
	Ok++

	// Удаляем данные
	if err := deleteTestStudent(studentID); err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++
	if err := deleteTestGroup(groupId); err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++

	return Ok, Bad, nil
}

func Test_CreateStudent() (int, int, error) {
	var Ok, Bad int

	// Создаём группу
	log.Println("[INFO] Создаём тестовую группу...")
	groupId, err := createTestGroup(1, 4, "Информатика", 101)
	if err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++

	// Создаём студента
	log.Println("[INFO] Создаём тестового студента...")
	fullName, enterprise, workStartDate, jobTitle := "Иван Иванов", "Tech Corp", "2023-06-01", "Developer"
	studentID, err := createTestStudent(fullName, groupId, enterprise, workStartDate, jobTitle)
	if err != nil {
		Bad++
		return Ok, Bad, err
	}
	if studentID == 0 {
		log.Println("[ERROR] Ошибка: ID студента не должен быть 0")
		Bad++
		return Ok, Bad, errors.New("нулевой ID студента")
	}
	Ok++

	// Проверяем создание
	log.Println("[INFO] Проверяем создание студента...")
	student, err := routes.Inf_StudentByID(studentID)
	if err != nil {
		log.Printf("Ошибка при получении студента с ID=%d: %v", studentID, err)
		Bad++
		return Ok, Bad, err
	}
	if err := checkStudentData(student, fullName, groupId, enterprise, workStartDate, jobTitle); err != nil {
		log.Printf("Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Println("[INFO] Студент успешно создан и проверен.")
	Ok++

	// Удаляем данные
	if err := deleteTestStudent(studentID); err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++
	if err := deleteTestGroup(groupId); err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++

	return Ok, Bad, nil
}

func Test_UpdateStudentById() (int, int, error) {
	var Ok, Bad int

	// Создаём первую группу
	log.Println("[INFO] Создаём тестовую группу...")
	groupId, err := createTestGroup(1, 4, "Информатика", 101)
	if err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++

	// Создаём вторую группу
	log.Println("[INFO] Создаём вторую тестовую группу...")
	newGroupId, err := createTestGroup(2, 5, "Кибербезопасность", 102)
	if err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++

	// Создаём студента
	log.Println("[INFO] Создаём тестового студента...")
	fullName, enterprise, workStartDate, jobTitle := "Иван Иванов", "Tech Corp", "2023-06-01", "Developer"
	studentID, err := createTestStudent(fullName, groupId, enterprise, workStartDate, jobTitle)
	if err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++

	// Обновляем студента
	log.Printf("[INFO] Обновляем студента с ID=%d...", studentID)
	newFullName, newEnterprise, newWorkStartDate, newJobTitle := "Пётр Петров", "New Corp", "2024-01-01", "Senior Developer"
	ok, err := routes.Update_StudentById(studentID, newFullName, newGroupId, newEnterprise, newWorkStartDate, newJobTitle)
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
	if err := checkStudentData(student, newFullName, newGroupId, newEnterprise, newWorkStartDate, newJobTitle); err != nil {
		log.Printf("Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Println("[INFO] Обновление студента подтверждено.")
	Ok++

	// Удаляем данные
	if err := deleteTestStudent(studentID); err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++
	if err := deleteTestGroup(groupId); err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++
	if err := deleteTestGroup(newGroupId); err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++

	return Ok, Bad, nil
}

func Test_DeleteStudent() (int, int, error) {
	var Ok, Bad int

	// Создаём группу
	log.Println("[INFO] Создаём тестовую группу...")
	groupId, err := createTestGroup(1, 4, "Информатика", 101)
	if err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++

	// Создаём студента
	log.Println("[INFO] Создаём тестового студента...")
	fullName, enterprise, workStartDate, jobTitle := "Иван Иванов", "Tech Corp", "2023-06-01", "Developer"
	studentID, err := createTestStudent(fullName, groupId, enterprise, workStartDate, jobTitle)
	if err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++

	// Удаляем студента
	log.Printf("[INFO] Удаляем студента с ID=%d...", studentID)
	if err := deleteTestStudent(studentID); err != nil {
		Bad++
		return Ok, Bad, err
	}
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
	if err := deleteTestGroup(groupId); err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++

	return Ok, Bad, nil
}

func Test_InfStudentByGroup() (int, int, error) {
	var Ok, Bad int

	// Создаём группу
	log.Println("[INFO] Создаём тестовую группу...")
	groupId, err := createTestGroup(1, 4, "Информатика", 101)
	if err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++

	// Создаём студентов
	log.Println("[INFO] Создаём тестовых студентов...")
	studentData := []struct {
		fullName, enterprise, workStartDate, jobTitle string
	}{
		{"Иван Иванов", "Tech Corp", "2023-06-01", "Developer"},
		{"Пётр Петров", "New Corp", "2024-01-01", "Analyst"},
	}
	studentIDs := make([]int, 0, len(studentData))
	for _, data := range studentData {
		studentID, err := createTestStudent(data.fullName, groupId, data.enterprise, data.workStartDate, data.jobTitle)
		if err != nil {
			Bad++
			return Ok, Bad, err
		}
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
	if len(students) != len(studentData) {
		log.Printf("Ошибка: ожидалось %d студентов, получено %d", len(studentData), len(students))
		Bad++
		return Ok, Bad, errors.New("неверное количество студентов")
	}
	for i, student := range students {
		if err := checkStudentData(student, studentData[i].fullName, groupId, studentData[i].enterprise, studentData[i].workStartDate, studentData[i].jobTitle); err != nil {
			log.Printf("Ошибка: %v", err)
			Bad++
			return Ok, Bad, err
		}
	}
	log.Printf("[INFO] Студенты группы успешно получены, найдено %d записей.", len(students))
	Ok++

	// Удаляем студентов
	for _, studentID := range studentIDs {
		if err := deleteTestStudent(studentID); err != nil {
			Bad++
			return Ok, Bad, err
		}
		Ok++
	}

	// Удаляем группу
	if err := deleteTestGroup(groupId); err != nil {
		Bad++
		return Ok, Bad, err
	}
	Ok++

	return Ok, Bad, nil
}

func TestStudentALL() (int, int) {
	var Ok, Bad int

	tests := []struct {
		name string
		fn   func() (int, int, error)
	}{
		{"InfAllStudent", Test_InfAllStudent},
		{"InfStudentByID", Test_InfStudentByID},
		{"CreateStudent", Test_CreateStudent},
		{"UpdateStudentById", Test_UpdateStudentById},
		{"DeleteStudent", Test_DeleteStudent},
		{"InfStudentByGroup", Test_InfStudentByGroup},
	}

	for _, test := range tests {
		ok, bad, err := test.fn()
		logTestResult(test.name, ok, bad, err)
		Ok += ok
		Bad += bad
		if err != nil {
			log.Fatal("Ошибка в ", test.name, ": ", err)
		}
	}

	log.Println("\n[SUMMARY] Итоговые результаты:")
	log.Printf("TestStudentALL [%d\\%d]\nok: %d\nbad: %d\n", Ok, Bad, Ok, Bad)

	return Ok, Bad
}
