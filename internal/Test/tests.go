package Test

import (
	"CRM_System/Backend/internal/db/repository"
	"CRM_System/Backend/internal/routes"
	"fmt"
	"log"
)

func TestGenerateFilledPDF() {
	studentID, _ := routes.Create_Student("Иван Иванов", 85, 3, "Программная инженерия", 101, 2)
	groupID, _ := routes.Create_Group(85, 3, "Программная инженерия", 101, 2)

	fmt.Printf("[INFO] Создан студент с ID = %d\n", studentID)
	fmt.Printf("[INFO] Создана группа с ID = %d\n", groupID)

	// Обновляем работодателя
	if _, err := routes.Update_EmpbyStudentId(studentID, "ООО Новая Компания", "2023-01-01", "Старший программист"); err != nil {
		fmt.Printf("[ERROR] Ошибка при обновлении работодателя: %v\n", err)
	} else {
		fmt.Println("[INFO] Работодатель успешно обновлён.")
	}

	// Добавляем предметы
	for i := 0; i < 10; i++ {
		subj := fmt.Sprintf("Subject %d", i)
		if _, err := routes.Add_SubjectByGroupId(groupID, subj); err != nil {
			fmt.Printf("[ERROR] Не удалось добавить предмет %s: %v\n", subj, err)
		}
	}

	// Генерация PDF
	Dock, err := routes.GenerateFilledPDF(studentID, "Иван Иванов", groupID)
	if err != nil || Dock == nil {
		fmt.Printf("[ERROR] Ошибка при генерации документа: %v\n", err)
	} else {
		fmt.Println("[INFO] Успешная генерация документа.")
	}

	// Удаление студента
	if ok, err := routes.Delete_Student(studentID); !ok || err != nil {
		fmt.Printf("[ERROR] Не удалось удалить студента: %v\n", err)
	} else {
		fmt.Println("[INFO] Студент успешно удалён.")
	}

	if ok, err := routes.Delete_EmpbyStudentId(studentID); !ok || err != nil {
		fmt.Printf("[ERROR] Не удалось удалить данные о работодателе: %v\n", err)
	} else {
		fmt.Println("[INFO] Успешное удаление данных о работодателе.")
	}

	// Удаление предметов
	if ok, err := routes.Delete_AllSubjectByGroupId(groupID); !ok || err != nil {
		fmt.Printf("[ERROR] Не удалось удалить предметы группы: %v\n", err)
	} else {
		fmt.Println("[INFO] Предметы успешно удалены.")
	}

	// Удаление группы
	if ok, err := routes.Delete_GroupById(groupID); !ok || err != nil {
		fmt.Printf("[ERROR] Не удалось удалить группу: %v\n", err)
	} else {
		fmt.Println("[INFO] Группа успешно удалена.")
	}
}

func TestAll() {
	var ok, bad int
	var err error

	// Тестирование студентов
	ok, bad, err = Test_StudentCRUD()
	logResults("Student", ok, bad, err)
	studentOk, studentBad := ok, bad

	// Тестирование групп
	ok, bad, err = Test_GroupCRUD()
	logResults("Group", ok, bad, err)
	groupOk, groupBad := ok, bad

	// Тестирование предметов
	ok, bad, err = Test_Subject()
	logResults("Subject", ok, bad, err)
	subjectOk, subjectBad := ok, bad

	// Тестирование работодателей
	ok, bad, err = Test_Employers()
	logResults("Employers", ok, bad, err)
	employersOk, employersBad := ok, bad

	// Итоговые результаты
	printSummary(studentOk, studentBad, groupOk, groupBad, subjectOk, subjectBad, employersOk, employersBad)
}

func logResults(testName string, ok, bad int, err error) {
	if err != nil {
		log.Printf("[ERROR] ❌ %s test failed: %v\n", testName, err)
	} else {
		log.Printf("[SUCCESS] ✅ %s test passed: %d, ❌ failed: %d\n", testName, ok, bad)
	}
}

func printSummary(studentOk, studentBad, groupOk, groupBad, subjectOk, subjectBad, employersOk, employersBad int) {
	log.Println("\n[SUMMARY] Итоговые результаты:")
	log.Printf("Student: ✅ %d | ❌ %d\n", studentOk, studentBad)
	log.Printf("Group: ✅ %d | ❌ %d\n", groupOk, groupBad)
	log.Printf("Subject: ✅ %d | ❌ %d\n", subjectOk, subjectBad)
	log.Printf("Employers: ✅ %d | ❌ %d\n", employersOk, employersBad)
}

func Test_StudentCRUD() (int, int, error) {
	var Ok int = 0
	var Bad int = 0

	// Получение всех записей
	log.Println("[INFO] Получаем всех студентов...")
	_, err := routes.Inf_AllStudent()
	if err != nil {
		log.Printf("Ошибка при получении студентов: %v", err)
		Bad++
	} else {
		log.Println("[INFO] Студенты успешно получены.")
		Ok++
	}

	// Добавление записи
	log.Println("[INFO] Добавляем нового студента...")
	studentID, err := routes.Create_Student("Иван Иванов", 85, 3, "Программная инженерия", 101, 2)
	if err != nil {
		log.Printf("Ошибка при добавлении студента: %v", err)
		Bad++
	} else {
		log.Printf("Студент с ID=%d успешно создан", studentID)
		Ok++
	}

	// Обновление записи
	log.Printf("[INFO] Обновляем студента с ID=%d...", studentID)
	ok, err := routes.Update_StudentById(studentID, "Пётр Петров", 90, 4, "Кибербезопасность", 102, 3)
	if err != nil || !ok {
		log.Printf("Ошибка при обновлении студента: %v", err)
		Bad++
	} else {
		log.Println("[INFO] Данные студента успешно обновлены.")
		Ok++
	}

	// Удаление записи
	log.Printf("[INFO] Удаляем студента с ID=%d...", studentID)
	ok, err = routes.Delete_Student(studentID)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении студента: %v", err)
		Bad++
	} else {
		log.Println("[INFO] Студент успешно удалён.")
		Ok++
	}

	// Проверка получения студента по ID
	log.Printf("[INFO] Получаем студента с ID=%d...", studentID)
	_, err = routes.Inf_StudentByID(studentID)
	if err == nil {
		log.Println("[ERROR] Ошибка: Студент должен быть удалён, но он всё ещё существует.")
		Bad++
	} else {
		log.Println("[INFO] Студент удалён, тест пройден.")
		Ok++
	}

	return Ok, Bad, nil
}

func Test_GroupCRUD() (int, int, error) {
	var Ok int = 0
	var Bad int = 0

	// Получение всех записей
	log.Println("[INFO] Получаем все группы...")
	_, err := routes.Inf_AllGroup()
	if err != nil {
		log.Printf("Ошибка при получении групп: %v", err)
		Bad++
	} else {
		log.Println("[INFO] Группы успешно получены.")
		Ok++
	}

	// Добавление записи
	log.Println("[INFO] Добавляем группу...")
	groupID, err := routes.Create_Group(1, 2, "Программная инженерия", 101, 1)
	if err != nil {
		log.Printf("Ошибка при добавлении группы: %v", err)
		Bad++
	} else {
		log.Printf("Группа с ID=%d успешно создана", groupID)
		Ok++
	}

	// Обновление записи
	log.Printf("[INFO] Обновляем группу с ID=%d...", groupID)
	ok, err := routes.Update_GroupById(groupID, 2, 3, "Кибербезопасность", 102, 2)
	if err != nil || !ok {
		log.Printf("Ошибка при обновлении группы: %v", err)
		Bad++
	} else {
		log.Println("[INFO] Данные группы успешно обновлены.")
		Ok++
	}

	// Удаление записи
	log.Printf("[INFO] Удаляем группу с ID=%d...", groupID)
	ok, err = routes.Delete_GroupById(groupID)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении группы: %v", err)
		Bad++
	} else {
		log.Println("[INFO] Группа успешно удалена.")
		Ok++
	}

	// Проверка получения ID группы по параметрам
	log.Println("[INFO] Проверяем удаление группы по параметрам...")
	_, err = routes.GetGroupId_GroupIdByInfo(1, 2, "Программная инженерия", 101, 1)
	if err == nil {
		log.Println("[ERROR] Ошибка: Группа должна быть удалена, но она всё ещё существует.")
		Bad++
	} else {
		log.Println("[INFO] Группа успешно удалена, тест пройден.")
		Ok++
	}

	return Ok, Bad, nil
}

func Test_Employers() (int, int, error) {
	// Шаг 1: Создаём студента с пустым бланком работодателя
	fullName := "Иван Иванов"
	well := byte(85)
	gClass := byte(3)
	speciality := "Программная инженерия"
	groupNum := 101
	semester := byte(2)

	studentID, err := repository.CreateStudentWithEmptyEmployment(fullName, well, gClass, speciality, groupNum, semester)
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании студента и его бланка работодателя: %v", err)
		return 0, 1, err
	}
	log.Printf("[INFO] Студент с ID=%d и его бланк работодателя успешно созданы.", studentID)

	var Ok, Bad int

	// Шаг 2: Проверяем, что бланк работодателя для студента был создан
	log.Println("[INFO] Проверяем наличие бланка работодателя для студента с ID=", studentID)
	employer, err := routes.Inf_EmpByStudentId(studentID)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении бланка работодателя: %v", err)
		Bad++
	} else {
		log.Printf("[INFO] Бланк работодателя для студента найден: %v", employer)
		Ok++
	}

	// Шаг 3: Обновляем данные работодателя
	log.Println("[INFO] Обновляем информацию о работодателе для студента с ID=", studentID)
	ok, err := routes.Update_EmpbyStudentId(studentID, "ООО Новая Компания", "2023-01-01", "Старший программист")
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при обновлении работодателя: %v", err)
		log.Printf("[INFO] Запрос не затронул ни одной строки (ok: %v)", ok)
		Bad++
	} else {
		log.Println("[INFO] Работодатель успешно обновлён.")
		Ok++
	}

	// Шаг 4: Получаем обновлённую информацию о работодателе
	log.Println("[INFO] Получаем обновлённую информацию о работодателе для студента с ID=", studentID)
	employer, err = routes.Inf_EmpByStudentId(studentID)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении обновлённого работодателя: %v", err)
		Bad++
	} else {
		log.Printf("[INFO] Работодатель для студента успешно обновлён: %v", employer)
		Ok++
	}

	// Шаг 5: Удаляем студента (и бланк работодателя)
	log.Println("[INFO] Удаляем студента с ID=", studentID)
	ok, err = routes.Delete_EmpbyStudentId(studentID)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении студента: %v", err)
		Bad++
	} else {
		log.Println("[INFO] Студент с ID=", studentID, " успешно удалён.")
		Ok++
	}
	// Шаг 6: Проверяем, что данные работодателя были удалены
	log.Println("[INFO] Проверяем данные работодателя после удаления студента с ID=", studentID)
	_, err = routes.Inf_EmpByStudentId(studentID)
	if err == nil {
		log.Println("[ERROR] Данные работодателя не должны быть найдены, но они всё ещё существуют!")
		log.Println(err)
		Bad++
	} else {
		log.Println("[INFO] Данные работодателя для студента с ID=", studentID, " успешно удалены.")
		Ok++
	}

	return Ok, Bad, nil
}

func Test_Subject() (int, int, error) {
	var Ok int = 0
	var Bad int = 0

	// Получение всех записей
	log.Println("[INFO] Получаем все предметы для группы с ID=101...")
	_, err := routes.Inf_SubjectByGroupId(101)
	if err != nil {
		log.Printf("Ошибка при получении предметов: %v", err)
		Bad++
	} else {
		log.Println("[INFO] Предметы успешно получены.")
		Ok++
	}

	// Добавление записи
	log.Println("[INFO] Добавляем предмет в группу...")
	subjectID, err := routes.Add_SubjectByGroupId(101, "Математика")
	if err != nil {
		log.Printf("Ошибка при добавлении предмета: %v", err)
		Bad++
	} else {
		log.Printf("Предмет с ID=%d успешно добавлен", subjectID)
		Ok++
	}

	// Обновление записи
	log.Printf("[INFO] Обновляем предмет с ID=%d...", subjectID)
	ok, err := routes.Update_SubjectById(subjectID, "Физика")
	if err != nil || !ok {
		log.Printf("Ошибка при обновлении предмета: %v", err)
		Bad++
	} else {
		log.Println("[INFO] Предмет успешно обновлён.")
		Ok++
	}

	// Удаление записи
	log.Printf("[INFO] Удаляем предмет с ID=%d...", subjectID)
	ok, err = routes.Delete_SubjectById(subjectID)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении предмета: %v", err)
		Bad++
	} else {
		log.Println("[INFO] Предмет успешно удалён.")
		Ok++
	}

	return Ok, Bad, nil
}
