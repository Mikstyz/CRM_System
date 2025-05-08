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
	course, groduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupId, err := routes.Create_Group(course, groduates, speciality, groupNum)
	if err != nil {
		log.Printf("Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if groupId == 0 {
		log.Println("[ERROR] Ошибка: ID групп не должны быть 0")
		Bad++
		return Ok, Bad, errors.New("нулевые ID групп")
	}
	log.Printf("Группа создана с ID %d", groupId)
	Ok++

	// Удаляем созданную группу
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

func Test_UpdateGroupById() (int, int, error) {
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

	// Обновляем группу
	log.Printf("[INFO] Обновляем группу с ID=%d...", groupId)
	newCourse, newGroduates, newSpeciality, newGroupNum := byte(2), byte(5), "Кибербезопасность", 102
	ok, err := routes.Update_GroupById(groupId, newCourse, newGroduates, newSpeciality, newGroupNum)
	if err != nil || !ok {
		log.Printf("Ошибка при обновлении группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка обновления группы")
	}
	log.Println("[INFO] Группа успешно обновлена.")
	Ok++

	// Проверяем обновление
	log.Println("[INFO] Проверяем обновление группы...")
	newGroupId, err := routes.GetGroupId_GroupIdByInfo(newCourse, newGroduates, newSpeciality, newGroupNum)
	if err != nil || newGroupId != groupId {
		log.Printf("Ошибка: обновлённая группа не найдена или неверный ID: %v, ожидался ID=%d, получен=%d", err, groupId, newGroupId)
		Bad++
		return Ok, Bad, errors.New("ошибка проверки обновления группы")
	}
	log.Println("[INFO] Обновление группы подтверждено.")
	Ok++

	// Удаляем тестовую группу
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

func Test_DeleteGroupById() (int, int, error) {
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

	// Удаляем группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err := routes.Delete_GroupById(groupId)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления группы")
	}
	log.Println("[INFO] Группа успешно удалена.")
	Ok++

	// Проверяем, что группа удалена
	log.Println("[INFO] Проверяем удаление группы...")
	_, err = routes.GetGroupId_GroupIdByInfo(course, groduates, speciality, groupNum)
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
	course, groduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupId, err := routes.Create_Group(course, groduates, speciality, groupNum)
	if err != nil {
		log.Printf("Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("Группа создана с ID %d", groupId)
	Ok++

	// Проверяем получение ID группы
	log.Println("[INFO] Проверяем получение ID группы...")
	receivedGroupId, err := routes.GetGroupId_GroupIdByInfo(course, groduates, speciality, groupNum)
	if err != nil || receivedGroupId != groupId {
		log.Printf("Ошибка при получении ID группы: %v, ожидался ID=%d, получен=%d", err, groupId, receivedGroupId)
		Bad++
		return Ok, Bad, errors.New("ошибка получения ID группы")
	}
	log.Println("[INFO] ID группы успешно получен.")
	Ok++

	// Удаляем тестовую группу
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

func Test_DuplicateGroupAllData() (int, int, error) {
	var Ok, Bad int

	// Создаём тестовую группу
	log.Println("[INFO] Создаём тестовую группу...")
	course, groduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupId, err := routes.Create_Group(course, groduates, speciality, groupNum)
	if err != nil {
		log.Printf("Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if groupId == 0 {
		log.Println("[ERROR] Ошибка: ID группы не должен быть 0")
		Bad++
		return Ok, Bad, errors.New("нулевой ID группы")
	}
	log.Printf("Группа создана с ID %d", groupId)
	Ok++

	// Дублируем группу
	log.Println("[INFO] Дублируем группу...")
	duplicatedData, err := routes.DublicateGroupAllData(groupId)
	if err != nil {
		log.Printf("Ошибка при дублировании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if duplicatedData.Id == 0 {
		log.Println("[ERROR] Ошибка: ID новой группы не должен быть 0")
		Bad++
		return Ok, Bad, errors.New("нулевой ID новой группы")
	}
	if duplicatedData.Course != course || duplicatedData.Spesiality != speciality || duplicatedData.Groduates != groduates || duplicatedData.Number != groupNum+1 {
		log.Println("[ERROR] Ошибка: параметры новой группы не соответствуют ожидаемым")
		Bad++
		return Ok, Bad, errors.New("неверные параметры новой группы")
	}
	log.Printf("Группа успешно продублирована с ID %d", duplicatedData.Id)
	Ok++

	// Проверяем новую группу
	log.Println("[INFO] Проверяем новую группу...")
	newGroupId, err := routes.GetGroupId_GroupIdByInfo(course, groduates, speciality, groupNum+1)
	if err != nil || newGroupId == 0 {
		log.Printf("Ошибка: новая группа не найдена или неверный ID: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка проверки новой группы")
	}
	if newGroupId != duplicatedData.Id {
		log.Println("[ERROR] Ошибка: ID новой группы не совпадает с возвращённым")
		Bad++
		return Ok, Bad, errors.New("несоответствие ID новой группы")
	}
	log.Printf("Новая группа найдена с ID %d", newGroupId)
	Ok++

	// Удаляем тестовую группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err := routes.Delete_GroupById(groupId)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[INFO] Тестовая группа успешно удалена.")
	Ok++

	// Удаляем новую группу
	log.Println("[INFO] Удаляем продублированную группу...")
	ok, err = routes.Delete_GroupById(newGroupId)
	if err != nil || !ok {
		log.Printf("Ошибка при удалении новой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления новой группы")
	}
	log.Println("[INFO] Новая группа успешно удалена.")
	Ok++

	return Ok, Bad, nil
}

func TestGroupALL() (int, int) {
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

	return Ok, Bad
}
