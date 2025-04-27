package Test

import (
	"CRM_System/internal/routes"
	"errors"
	"log"
)

// Вспомогательная функция для проверки, что поле не nil и равно ожидаемому значению

func checkField(field *string, expected, fieldName string) error {
	if field == nil {
		return errors.New(fieldName + " не должно быть nil")
	}
	if *field != expected {
		return errors.New(fieldName + " имеет неверное значение, ожидалось: " + expected + ", получено: " + *field)
	}
	return nil
}

func logResults(testName string, ok, bad int, err error) {
	if err != nil {
		log.Printf("[ERROR] ❌ %s test failed: %v\n", testName, err)
	} else {
		log.Printf("[SUCCESS] ✅ %s test passed: %d, ❌ failed: %d\n", testName, ok, bad)
	}
}

// Тест получения данных о трудоустройстве
func Test_InfEmpByStudentId() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, graduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupId, err := routes.Create_Group(course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Группа создана с ID=%d", groupId)
	Ok++

	// Создаём студента
	log.Println("[INFO] Создаём тестового студента...")
	fullName := "Иван Иванов"
	studentID, err := routes.Create_Student(fullName, course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании студента: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Студент создан с ID=%d", studentID)
	Ok++

	// Проверяем пустые данные о трудоустройстве после создания студента
	log.Println("[INFO] Проверяем пустые данные о трудоустройстве...")
	emp, err := routes.Inf_EmpByStudentId(studentID)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if emp.Enterprise != nil || emp.WorkStartDate != nil || emp.JobTitle != nil {
		log.Printf("[ERROR] Данные о трудоустройстве должны быть nil, получено=%v", emp)
		Bad++
		return Ok, Bad, errors.New("данные о трудоустройстве не пусты")
	}
	log.Println("[SUCCESS] Пустые данные о трудоустройстве подтверждены")
	Ok++

	// Обновляем данные о трудоустройстве
	log.Println("[INFO] Обновляем данные о трудоустройстве...")
	enterprise, workStartDate, jobTitle := "ООО Ромашка", "2023-01-01", "Программист"
	ok, err := routes.Update_EmpbyStudentId(studentID, enterprise, workStartDate, jobTitle)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при обновлении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка обновления данных о трудоустройстве")
	}
	log.Println("[SUCCESS] Данные о трудоустройстве обновлены")
	Ok++

	// Получаем данные о трудоустройстве
	log.Println("[INFO] Получаем данные о трудоустройстве по ID студента...")
	emp, err = routes.Inf_EmpByStudentId(studentID)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[DEBUG] Данные о трудоустройстве: Enterprise=%v, WorkStartDate=%v, JobTitle=%v", emp.Enterprise, emp.WorkStartDate, emp.JobTitle)
	if err := checkField(emp.Enterprise, enterprise, "enterprise"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if err := checkField(emp.WorkStartDate, workStartDate, "work_start_date"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if err := checkField(emp.JobTitle, jobTitle, "job_title"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Println("[SUCCESS] Данные о трудоустройстве успешно получены")
	Ok++

	// Удаляем данные о трудоустройстве
	log.Printf("[INFO] Удаляем данные о трудоустройстве для студента с ID=%d...", studentID)
	ok, err = routes.Delete_EmpbyStudentId(studentID)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления данных о трудоустройстве")
	}
	log.Println("[SUCCESS] Данные о трудоустройстве успешно удалены")
	Ok++

	// Проверяем, что данные о трудоустройстве удалены
	log.Println("[INFO] Проверяем, что данные о трудоустройстве удалены...")
	_, err = routes.Inf_EmpByStudentId(studentID)
	if err == nil {
		log.Printf("[ERROR] Данные о трудоустройстве должны быть удалены, но они всё ещё существуют")
		Bad++
		return Ok, Bad, errors.New("данные о трудоустройстве не были удалены")
	}
	log.Println("[SUCCESS] Данные о трудоустройстве успешно удалены, тест пройден")
	Ok++

	// Удаляем студента
	log.Printf("[INFO] Удаляем студента с ID=%d...", studentID)
	ok, err = routes.Delete_Student(studentID)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении студента: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления студента")
	}
	log.Println("[SUCCESS] Студент успешно удалён")
	Ok++

	// Проверяем, что данные о трудоустройстве удалены после удаления студента
	log.Println("[INFO] Проверяем, что данные о трудоустройстве удалены после удаления студента...")
	_, err = routes.Inf_EmpByStudentId(studentID)
	if err == nil {
		log.Printf("[ERROR] Данные о трудоустройстве должны быть удалены после удаления студента")
		Bad++
		return Ok, Bad, errors.New("данные о трудоустройстве не были удалены после удаления студента")
	}
	log.Println("[SUCCESS] Данные о трудоустройстве отсутствуют после удаления студента")
	Ok++

	// Удаляем группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err = routes.Delete_GroupById(course, graduates, speciality, groupNum)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[SUCCESS] Тестовая группа успешно удалена")
	Ok++

	return Ok, Bad, nil
}

// Тест создания данных о трудоустройстве
func Test_CreateEmployers() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, graduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupId, err := routes.Create_Group(course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Группа создана с ID %d", groupId)
	Ok++

	// Создаём студента
	log.Println("[INFO] Создаём тестового студента...")
	fullName := "Иван Иванов"
	studentID, err := routes.Create_Student(fullName, course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании студента: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Студент создан с ID=%d", studentID)
	Ok++

	// Проверяем пустые данные о трудоустройстве после создания студента
	log.Println("[INFO] Проверяем пустые данные о трудоустройстве...")
	emp, err := routes.Inf_EmpByStudentId(studentID)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if emp.Enterprise != nil || emp.WorkStartDate != nil || emp.JobTitle != nil {
		log.Printf("[ERROR] Данные о трудоустройстве должны быть nil, получено=%v", emp)
		Bad++
		return Ok, Bad, errors.New("данные о трудоустройстве не пусты")
	}
	log.Println("[SUCCESS] Пустые данные о трудоустройстве подтверждены")
	Ok++

	// Удаляем запись о трудоустройстве, чтобы проверить Create_Employers
	log.Printf("[INFO] Удаляем данные о трудоустройстве для студента с ID=%d...", studentID)
	ok, err := routes.Delete_EmpbyStudentId(studentID)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления данных о трудоустройстве")
	}
	log.Println("[SUCCESS] Данные о трудоустройстве успешно удалены")
	Ok++

	// Проверяем, что запись удалена
	log.Println("[INFO] Проверяем, что данные о трудоустройстве удалены...")
	_, err = routes.Inf_EmpByStudentId(studentID)
	if err == nil {
		log.Printf("[ERROR] Данные о трудоустройстве должны быть удалены, но они всё ещё существуют")
		Bad++
		return Ok, Bad, errors.New("данные о трудоустройстве не были удалены")
	}
	log.Println("[SUCCESS] Данные о трудоустройстве успешно удалены")
	Ok++

	// Создаём данные о трудоустройстве через Create_Employers
	log.Println("[INFO] Создаём данные о трудоустройстве через Create_Employers...")
	enterprise, workStartDate, jobTitle := "ООО Ромашка", "2023-01-01", "Программист"
	routes.Create_Employers(studentID, enterprise, workStartDate, jobTitle)
	log.Println("[SUCCESS] Данные о трудоустройстве созданы")
	Ok++

	// Проверяем создание
	log.Println("[INFO] Проверяем создание данных о трудоустройстве...")
	emp, err = routes.Inf_EmpByStudentId(studentID)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[DEBUG] Данные о трудоустройстве: Enterprise=%v, WorkStartDate=%v, JobTitle=%v", emp.Enterprise, emp.WorkStartDate, emp.JobTitle)
	if err := checkField(emp.Enterprise, enterprise, "enterprise"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if err := checkField(emp.WorkStartDate, workStartDate, "work_start_date"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if err := checkField(emp.JobTitle, jobTitle, "job_title"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Println("[SUCCESS] Данные о трудоустройстве успешно созданы и проверены")
	Ok++

	// Удаляем данные о трудоустройстве
	log.Printf("[INFO] Удаляем данные о трудоустройстве для студента с ID=%d...", studentID)
	ok, err = routes.Delete_EmpbyStudentId(studentID)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления данных о трудоустройства")
	}
	log.Println("[SUCCESS] Данные о трудоустройстве успешно удалены")
	Ok++

	// Удаляем студента
	log.Printf("[INFO] Удаляем студента с ID=%d...", studentID)
	ok, err = routes.Delete_Student(studentID)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении студента: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления студента")
	}
	log.Println("[SUCCESS] Студент успешно удалён")
	Ok++

	// Проверяем, что данные о трудоустройстве удалены после удаления студента
	log.Println("[INFO] Проверяем, что данные о трудоустройстве удалены после удаления студента...")
	_, err = routes.Inf_EmpByStudentId(studentID)
	if err == nil {
		log.Printf("[ERROR] Данные о трудоустройстве должны быть удалены после удаления студента")
		Bad++
		return Ok, Bad, errors.New("данные о трудоустройстве не были удалены после удаления студента")
	}
	log.Println("[SUCCESS] Данные о трудоустройстве отсутствуют после удаления студента")
	Ok++

	// Удаляем группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err = routes.Delete_GroupById(course, graduates, speciality, groupNum)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[SUCCESS] Тестовая группа успешно удалена")
	Ok++

	return Ok, Bad, nil
}

// Тест обновления данных о трудоустройстве
func Test_UpdateEmpByStudentId() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, graduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupId, err := routes.Create_Group(course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Группа создана с ID %d", groupId)
	Ok++

	// Создаём студента
	log.Println("[INFO] Создаём тестового студента...")
	fullName := "Иван Иванов"
	studentID, err := routes.Create_Student(fullName, course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании студента: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Студент создан с ID=%d", studentID)
	Ok++

	// Проверяем пустые данные о трудоустройстве после создания студента
	log.Println("[INFO] Проверяем пустые данные о трудоустройстве...")
	emp, err := routes.Inf_EmpByStudentId(studentID)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if emp.Enterprise != nil || emp.WorkStartDate != nil || emp.JobTitle != nil {
		log.Printf("[ERROR] Данные о трудоустройстве должны быть nil, получено=%v", emp)
		Bad++
		return Ok, Bad, errors.New("данные о трудоустройстве не пусты")
	}
	log.Println("[SUCCESS] Пустые данные о трудоустройстве подтверждены")
	Ok++

	// Обновляем данные о трудоустройстве (первое обновление)
	log.Println("[INFO] Обновляем данные о трудоустройстве (первое обновление)...")
	enterprise, workStartDate, jobTitle := "ООО Ромашка", "2023-01-01", "Программист"
	ok, err := routes.Update_EmpbyStudentId(studentID, enterprise, workStartDate, jobTitle)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при первом обновлении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка первого обновления данных о трудоустройстве")
	}
	log.Println("[SUCCESS] Данные о трудоустройстве успешно обновлены (первое обновление)")
	Ok++

	// Проверяем первое обновление
	log.Println("[INFO] Проверяем первое обновление данных о трудоустройстве...")
	emp, err = routes.Inf_EmpByStudentId(studentID)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[DEBUG] Данные о трудоустройстве: Enterprise=%v, WorkStartDate=%v, JobTitle=%v", emp.Enterprise, emp.WorkStartDate, emp.JobTitle)
	if err := checkField(emp.Enterprise, enterprise, "enterprise"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if err := checkField(emp.WorkStartDate, workStartDate, "work_start_date"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if err := checkField(emp.JobTitle, jobTitle, "job_title"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Println("[SUCCESS] Первое обновление данных о трудоустройстве подтверждено")
	Ok++

	// Обновляем данные о трудоустройстве (второе обновление)
	log.Println("[INFO] Обновляем данные о трудоустройстве (второе обновление)...")
	newEnterprise, newWorkStartDate, newJobTitle := "ООО Лютик", "2024-01-01", "Старший программист"
	ok, err = routes.Update_EmpbyStudentId(studentID, newEnterprise, newWorkStartDate, newJobTitle)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при втором обновлении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка второго обновления данных о трудоустройстве")
	}
	log.Println("[SUCCESS] Данные о трудоустройстве успешно обновлены (второе обновление)")
	Ok++

	// Проверяем второе обновление
	log.Println("[INFO] Проверяем второе обновление данных о трудоустройстве...")
	emp, err = routes.Inf_EmpByStudentId(studentID)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[DEBUG] Данные о трудоустройстве: Enterprise=%v, WorkStartDate=%v, JobTitle=%v", emp.Enterprise, emp.WorkStartDate, emp.JobTitle)
	if err := checkField(emp.Enterprise, newEnterprise, "enterprise"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if err := checkField(emp.WorkStartDate, newWorkStartDate, "work_start_date"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if err := checkField(emp.JobTitle, newJobTitle, "job_title"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Println("[SUCCESS] Второе обновление данных о трудоустройстве подтверждено")
	Ok++

	// Удаляем данные о трудоустройстве
	log.Printf("[INFO] Удаляем данные о трудоустройстве для студента с ID=%d...", studentID)
	ok, err = routes.Delete_EmpbyStudentId(studentID)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления данных о трудоустройства")
	}
	log.Println("[SUCCESS] Данные о трудоустройстве успешно удалены")
	Ok++

	// Удаляем студента
	log.Printf("[INFO] Удаляем студента с ID=%d...", studentID)
	ok, err = routes.Delete_Student(studentID)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении студента: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления студента")
	}
	log.Println("[SUCCESS] Студент успешно удалён")
	Ok++

	// Проверяем, что данные о трудоустройстве удалены после удаления студента
	log.Println("[INFO] Проверяем, что данные о трудоустройстве удалены после удаления студента...")
	_, err = routes.Inf_EmpByStudentId(studentID)
	if err == nil {
		log.Printf("[ERROR] Данные о трудоустройстве должны быть удалены после удаления студента")
		Bad++
		return Ok, Bad, errors.New("данные о трудоустройстве не были удалены после удаления студента")
	}
	log.Println("[SUCCESS] Данные о трудоустройстве отсутствуют после удаления студента")
	Ok++

	// Удаляем группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err = routes.Delete_GroupById(course, graduates, speciality, groupNum)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[SUCCESS] Тестовая группа успешно удалена")
	Ok++

	return Ok, Bad, nil
}

// Тест удаления данных о трудоустройстве
func Test_DeleteEmpByStudentId() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, graduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupId, err := routes.Create_Group(course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Группа создана с ID %d", groupId)
	Ok++

	// Создаём студента
	log.Println("[INFO] Создаём тестового студента...")
	fullName := "Иван Иванов"
	studentID, err := routes.Create_Student(fullName, course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании студента: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Студент создан с ID=%d", studentID)
	Ok++

	// Проверяем пустые данные о трудоустройстве после создания студента
	log.Println("[INFO] Проверяем пустые данные о трудоустройстве...")
	emp, err := routes.Inf_EmpByStudentId(studentID)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if emp.Enterprise != nil || emp.WorkStartDate != nil || emp.JobTitle != nil {
		log.Printf("[ERROR] Данные о трудоустройстве должны быть nil, получено=%v", emp)
		Bad++
		return Ok, Bad, errors.New("данные о трудоустройстве не пусты")
	}
	log.Println("[SUCCESS] Пустые данные о трудоустройстве подтверждены")
	Ok++

	// Обновляем данные о трудоустройстве
	log.Println("[INFO] Обновляем данные о трудоустройстве...")
	enterprise, workStartDate, jobTitle := "ООО Ромашка", "2023-01-01", "Программист"
	ok, err := routes.Update_EmpbyStudentId(studentID, enterprise, workStartDate, jobTitle)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при обновлении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка обновления данных о трудоустройстве")
	}
	log.Println("[SUCCESS] Данные о трудоустройстве успешно обновлены")
	Ok++

	// Проверяем обновление
	log.Println("[INFO] Проверяем обновление данных о трудоустройстве...")
	emp, err = routes.Inf_EmpByStudentId(studentID)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[DEBUG] Данные о трудоустройстве: Enterprise=%v, WorkStartDate=%v, JobTitle=%v", emp.Enterprise, emp.WorkStartDate, emp.JobTitle)
	if err := checkField(emp.Enterprise, enterprise, "enterprise"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if err := checkField(emp.WorkStartDate, workStartDate, "work_start_date"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if err := checkField(emp.JobTitle, jobTitle, "job_title"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Println("[SUCCESS] Данные о трудоустройстве успешно обновлены и проверены")
	Ok++

	// Удаляем данные о трудоустройстве
	log.Printf("[INFO] Удаляем данные о трудоустройстве для студента с ID=%d...", studentID)
	ok, err = routes.Delete_EmpbyStudentId(studentID)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении данных о трудоустройстве: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления данных о трудоустройства")
	}
	log.Println("[SUCCESS] Данные о трудоустройстве успешно удалены")
	Ok++

	// Проверяем удаление
	log.Println("[INFO] Проверяем удаление данных о трудоустройстве...")
	_, err = routes.Inf_EmpByStudentId(studentID)
	if err == nil {
		log.Printf("[ERROR] Данные о трудоустройстве должны быть удалены, но они всё ещё существуют")
		Bad++
		return Ok, Bad, errors.New("данные о трудоустройстве не были удалены")
	}
	log.Println("[SUCCESS] Данные о трудоустройстве успешно удалены, тест пройден")
	Ok++

	// Удаляем студента
	log.Printf("[INFO] Удаляем студента с ID=%d...", studentID)
	ok, err = routes.Delete_Student(studentID)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении студента: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления студента")
	}
	log.Println("[SUCCESS] Студент успешно удалён")
	Ok++

	// Проверяем, что данные о трудоустройстве удалены после удаления студента
	log.Println("[INFO] Проверяем, что данные о трудоустройстве удалены после удаления студента...")
	_, err = routes.Inf_EmpByStudentId(studentID)
	if err == nil {
		log.Printf("[ERROR] Данные о трудоустройстве должны быть удалены после удаления студента")
		Bad++
		return Ok, Bad, errors.New("данные о трудоустройстве не были удалены после удаления студента")
	}
	log.Println("[SUCCESS] Данные о трудоустройстве отсутствуют после удаления студента")
	Ok++

	// Удаляем группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err = routes.Delete_GroupById(course, graduates, speciality, groupNum)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[SUCCESS] Тестовая группа успешно удалена")
	Ok++

	return Ok, Bad, nil
}

// Тест всех функций Employers
func TestEmployersALL() (int, int) {
	var Ok, Bad int

	// Тестирование Inf_EmpByStudentId
	ok, bad, err := Test_InfEmpByStudentId()
	if err != nil {
		log.Fatal("Ошибка в Test_InfEmpByStudentId: ", err)
	}
	logResults("InfEmpByStudentId", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование Create_Employers
	ok, bad, err = Test_CreateEmployers()
	if err != nil {
		log.Fatal("Ошибка в Test_CreateEmployers: ", err)
	}
	logResults("CreateEmployers", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование Update_EmpByStudentId DRAMAT
	ok, bad, err = Test_UpdateEmpByStudentId()
	if err != nil {
		log.Fatal("Ошибка в Test_UpdateEmpByStudentId: ", err)
	}
	logResults("UpdateEmpByStudentId", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование Delete_EmpByStudentId
	ok, bad, err = Test_DeleteEmpByStudentId()
	if err != nil {
		log.Fatal("Ошибка в Test_DeleteEmpByStudentId: ", err)
	}
	logResults("DeleteEmpByStudentId", ok, bad, err)
	Ok += ok
	Bad += bad

	// Итоговый вывод
	log.Println("\n[SUMMARY] Итоговые результаты:")
	log.Printf("TestEmployersALL [%d\\%d]\nok: %d\nbad: %d\n", Ok, Bad, Ok, Bad)

	return ok, bad
}
