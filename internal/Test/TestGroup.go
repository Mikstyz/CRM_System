package Test

import (
	"CRM_System/internal/routes"
	"errors"
	"log" // подставь путь к моделям
)

func Test_InfAllGroup() (int, int, error) {
	var Ok, Bad int

	log.Println("[INFO] Получаем все группы...")
	groups, err := routes.Inf_AllGroup()
	if err != nil {
		log.Printf("Ошибка при получении групп: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[INFO] Группы успешно получены, найдено %d записей.", len(groups))
	Ok++

	return Ok, Bad, nil
}

func Test_CreateGroup() (int, int, error) {
	var Ok, Bad int

	log.Println("[INFO] Создаём тестовую группу...")
	course, graduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupIds, err := routes.Create_Group(course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if groupIds[0] == 0 || groupIds[1] == 0 {
		log.Println("[ERROR] Ошибка: ID групп не должны быть 0")
		Bad++
		return Ok, Bad, errors.New("нулевые ID групп")
	}
	log.Printf("Группа создана с ID=[%d, %d]", groupIds[0], groupIds[1])
	Ok++

	// Удаляем созданную группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err := routes.Delete_GroupById(course, graduates, speciality, groupNum)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[INFO] Тестовая группа успешно удалена.")
	Ok++

	return Ok, Bad, nil
}

func Test_UpdateGroupById() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, graduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupIds, err := routes.Create_Group(course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("Группа создана с ID=[%d, %d]", groupIds[0], groupIds[1])
	Ok++

	// Обновляем группу
	log.Printf("[INFO] Обновляем группу с ID=%d...", groupIds[0])
	newCourse, newGraduates, newSpeciality, newGroupNum := byte(2), byte(5), "Кибербезопасность", 102
	ok, err := routes.Update_GroupById(groupIds[0], newCourse, newGraduates, newSpeciality, newGroupNum)
	if err != nil || !ok {
		log.Printf("Ошибка при обновлении группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка обновления группы")
	}
	log.Println("[INFO] Группа успешно обновлена.")
	Ok++

	// Проверяем обновление
	log.Println("[INFO] Проверяем обновление группы...")
	groupId, err := routes.GetGroupId_GroupIdByInfo(newCourse, newGraduates, newSpeciality, newGroupNum, 1)
	if err != nil || groupId != groupIds[0] {
		log.Printf("Ошибка: обновлённая группа не найдена или неверный ID: %v, ожидался ID=%d, получен=%d", err, groupIds[0], groupId)
		Bad++
		return Ok, Bad, errors.New("ошибка проверки обновления группы")
	}
	log.Println("[INFO] Обновление группы подтверждено.")
	Ok++

	// Удаляем тестовую группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err = routes.Delete_GroupById(newCourse, newGraduates, newSpeciality, newGroupNum)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[INFO] Тестовая группа успешно удалена.")
	Ok++

	return Ok, Bad, nil
}

func Test_DeleteGroupById() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, graduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupIds, err := routes.Create_Group(course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("Группа создана с ID=[%d, %d]", groupIds[0], groupIds[1])
	Ok++

	// Удаляем группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err := routes.Delete_GroupById(course, graduates, speciality, groupNum)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления группы")
	}
	log.Println("[INFO] Группа успешно удалена.")
	Ok++

	// Проверяем, что группа удалена
	log.Println("[INFO] Проверяем удаление группы...")
	_, err = routes.GetGroupId_GroupIdByInfo(course, graduates, speciality, groupNum, 1)
	if err == nil {
		log.Println("[ERROR] Ошибка: группа должна быть удалена, но она всё ещё существует.")
		Bad++
		return Ok, Bad, errors.New("группа не была удалена")
	}
	log.Println("[INFO] Группа успешно удалена, тест пройден.")
	Ok++

	return Ok, Bad, nil
}

func Test_GetGroupIdByInfo() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, graduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupIds, err := routes.Create_Group(course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("Группа создана с ID=[%d, %d]", groupIds[0], groupIds[1])
	Ok++

	// Проверяем получение ID для 1-го семестра
	log.Println("[INFO] Проверяем получение ID группы для 1-го семестра...")
	groupId, err := routes.GetGroupId_GroupIdByInfo(course, graduates, speciality, groupNum, 1)
	if err != nil || groupId != groupIds[0] {
		log.Printf("Ошибка при получении ID группы для 1-го семестра: %v, ожидался ID=%d, получен=%d", err, groupIds[0], groupId)
		Bad++
		return Ok, Bad, errors.New("ошибка получения ID группы для 1-го семестра")
	}
	log.Println("[INFO] ID группы для 1-го семестра успешно получен.")
	Ok++

	// Проверяем получение ID для 2-го семестра
	log.Println("[INFO] Проверяем получение ID группы для 2-го семестра...")
	groupId, err = routes.GetGroupId_GroupIdByInfo(course, graduates, speciality, groupNum, 2)
	if err != nil || groupId != groupIds[1] {
		log.Printf("Ошибка при получении ID группы для 2-го семестра: %v, ожидался ID=%d, получен=%d", err, groupIds[1], groupId)
		Bad++
		return Ok, Bad, errors.New("ошибка получения ID группы для 2-го семестра")
	}
	log.Println("[INFO] ID группы для 2-го семестра успешно получен.")
	Ok++

	// Удаляем тестовую группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err := routes.Delete_GroupById(course, graduates, speciality, groupNum)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[INFO] Тестовая группа успешно удалена.")
	Ok++

	return Ok, Bad, nil
}

func Test_DuplicateGroupAllData() (int, int, error) {
	var Ok, Bad int

	// Создаём тестовую группу
	log.Println("[INFO] Создаём тестовую группу...")
	course, graduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupIds, err := routes.Create_Group(course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("Группа создана с ID=[%d, %d]", groupIds[0], groupIds[1])
	Ok++

	// Добавляем предметы в оба семестра
	log.Println("[INFO] Добавляем предметы в 1-й семестр...")
	subjectsSem1 := []string{"Математика", "Программирование"}
	for _, subj := range subjectsSem1 {
		subjectId, err := routes.Add_SubjectByGroupId(groupIds[0], subj)
		if err != nil {
			log.Printf("Ошибка при добавлении предмета %s в 1-й семестр: %v", subj, err)
			Bad++
			return Ok, Bad, errors.New("ошибка добавления предмета в 1-й семестр")
		}
		log.Printf("Предмет %s добавлен в 1-й семестр с ID=%d", subj, subjectId)
		Ok++
	}

	log.Println("[INFO] Добавляем предметы во 2-й семестр...")
	subjectsSem2 := []string{"Физика", "Базы данных"}
	for _, subj := range subjectsSem2 {
		subjectId, err := routes.Add_SubjectByGroupId(groupIds[1], subj)
		if err != nil {
			log.Printf("Ошибка при добавлении предмета %s во 2-й семестр: %v", subj, err)
			Bad++
			return Ok, Bad, errors.New("ошибка добавления предмета во 2-й семестр")
		}
		log.Printf("Предмет %s добавлен во 2-й семестр с ID=%d", subj, subjectId)
		Ok++
	}

	// Дублируем группу
	log.Println("[INFO] Дублируем группу со всеми данными...")
	newGroupNum := groupNum + 1
	success, err := routes.DublicateGroupAllData(course, graduates, speciality, groupNum)
	if err != nil || !success {
		log.Printf("Ошибка при дублировании группы: %v, success=%v", err, success)
		Bad++
		return Ok, Bad, errors.New("ошибка дублирования группы")
	}
	log.Println("[INFO] Группа успешно продублирована.")
	Ok++

	// Проверяем создание новой группы
	log.Println("[INFO] Проверяем создание новой группы...")
	newGroupId, err := routes.GetGroupId_GroupIdByInfo(course, graduates, speciality, newGroupNum, 1)
	if err != nil || newGroupId == 0 {
		log.Printf("Ошибка: новая группа не найдена: %v", err)
		Bad++
		return Ok, Bad, errors.New("новая группа не найдена")
	}
	log.Printf("Новая группа найдена с ID=%d", newGroupId)
	Ok++

	// Проверяем копирование предметов
	log.Println("[INFO] Проверяем копирование предметов для 1-го семестра новой группы...")
	newSubjects, err := routes.Inf_SubjectByGroupId(newGroupId)
	if err != nil || len(newSubjects) != len(subjectsSem1) {
		log.Printf("Ошибка при получении предметов новой группы: %v, ожидалось %d предметов, получено %d", err, len(subjectsSem1), len(newSubjects))
		Bad++
		return Ok, Bad, errors.New("ошибка копирования предметов")
	}
	log.Println("[INFO] Предметы для 1-го семестра новой группы успешно скопированы.")
	Ok++

	// Удаляем тестовую и новую группы
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err := routes.Delete_GroupById(course, graduates, speciality, groupNum)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[INFO] Тестовая группа успешно удалена.")
	Ok++

	log.Println("[INFO] Удаляем новую группу...")
	ok, err = routes.Delete_GroupById(course, graduates, speciality, newGroupNum)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении новой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления новой группы")
	}
	log.Println("[INFO] Новая группа успешно удалена.")
	Ok++

	return Ok, Bad, nil
}

func TestGroupALL() {
	var Ok, Bad int

	// Тестирование Inf_AllGroup
	ok, bad, err := Test_InfAllGroup()
	if err != nil {
		log.Fatal("Ошибка в Test_InfAllGroup: ", err)
	}
	logResults("InfAllGroup", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование Create_Group
	ok, bad, err = Test_CreateGroup()
	if err != nil {
		log.Fatal("Ошибка в Test_CreateGroup: ", err)
	}
	logResults("CreateGroup", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование Update_GroupById
	ok, bad, err = Test_UpdateGroupById()
	if err != nil {
		log.Fatal("Ошибка в Test_UpdateGroupById: ", err)
	}
	logResults("UpdateGroupById", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование Delete_GroupById
	ok, bad, err = Test_DeleteGroupById()
	if err != nil {
		log.Fatal("Ошибка в Test_DeleteGroupById: ", err)
	}
	logResults("DeleteGroupById", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование GetGroupId_GroupIdByInfo
	ok, bad, err = Test_GetGroupIdByInfo()
	if err != nil {
		log.Fatal("Ошибка в Test_GetGroupIdByInfo: ", err)
	}
	logResults("GetGroupIdByInfo", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование DublicateGroupAllData
	ok, bad, err = Test_DuplicateGroupAllData()
	if err != nil {
		log.Fatal("Ошибка в Test_DuplicateGroupAllData: ", err)
	}
	logResults("DublicateGroupAllData", ok, bad, err)
	Ok += ok
	Bad += bad

	// Итоговый вывод
	log.Println("\n[SUMMARY] Итоговые результаты:")
	log.Printf("TestGroupALL [%d\\%d]\nok: %d\nbad: %d\n", Ok, Bad, Ok, Bad)
}
