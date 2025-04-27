package Test

import (
	"CRM_System/internal/routes"
	"errors"
	"fmt"
	"log"
	"reflect"
)

//Тест получения предметов сразу на 2 семестра

func Test_Inf_DisciplinesByGroupData() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, grads, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupIds, err := routes.Create_Group(course, grads, speciality, groupNum)
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Группа создана с ID=[%d, %d]", groupIds[0], groupIds[1])
	Ok++
	groupId1, groupId2 := groupIds[0], groupIds[1] // ID для семестров 1 и 2

	// Проверяем пустой список предметов
	log.Println("[INFO] Проверяем пустой список предметов...")
	disciplines, err := routes.Inf_DisciplinesByGroupData(speciality, groupNum, int(course), int(grads))
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	expectedEmpty := map[int][]string{1: {}, 2: {}}
	if !compareDisciplines(disciplines, expectedEmpty) {
		log.Printf("[ERROR] Список предметов должен быть пустым, получено=%v", disciplines)
		Bad++
		return Ok, Bad, errors.New("список предметов не пуст")
	}
	log.Println("[SUCCESS] Пустой список предметов подтверждён")
	Ok++

	// Добавляем предметы для семестра 1
	log.Println("[INFO] Добавляем предмет для семестра 1...")
	subject1 := "Математика"
	subjectId1, err := routes.Add_SubjectByGroupId(groupId1, subject1)
	if err != nil {
		log.Printf("[ERROR] Ошибка при добавлении предмета '%s': %v", subject1, err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Предмет '%s' добавлен с ID=%d", subject1, subjectId1)
	Ok++

	// Добавляем предметы для семестра 2
	log.Println("[INFO] Добавляем предмет для семестра 2...")
	subject2 := "Программирование"
	subjectId2, err := routes.Add_SubjectByGroupId(groupId2, subject2)
	if err != nil {
		log.Printf("[ERROR] Ошибка при добавлении предмета '%s': %v", subject2, err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Предмет '%s' добавлен с ID=%d", subject2, subjectId2)
	Ok++

	// Проверяем список предметов после добавления
	log.Println("[INFO] Проверяем список предметов после добавления...")
	disciplines, err = routes.Inf_DisciplinesByGroupData(speciality, groupNum, int(course), int(grads))
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[DEBUG] Список предметов: %v", disciplines)
	expected := map[int][]string{
		1: {subject1},
		2: {subject2},
	}
	if !compareDisciplines(disciplines, expected) {
		log.Printf("[ERROR] Данные не совпадают, получено=%v, ожидаемо=%v", disciplines, expected)
		Bad++
		return Ok, Bad, errors.New("список предметов не совпадает")
	}
	log.Println("[SUCCESS] Данные предметов корректны")
	Ok++

	// Удаляем предметы
	log.Printf("[INFO] Удаляем предмет с ID=%d...", subjectId1)
	ok, err := routes.Delete_SubjectById(subjectId1)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении предмета '%s': %v", subject1, err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления предмета")
	}
	log.Printf("[SUCCESS] Предмет '%s' успешно удалён", subject1)
	Ok++

	log.Printf("[INFO] Удаляем предмет с ID=%d...", subjectId2)
	ok, err = routes.Delete_SubjectById(subjectId2)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении предмета '%s': %v", subject2, err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления предмета")
	}
	log.Printf("[SUCCESS] Предмет '%s' успешно удалён", subject2)
	Ok++

	// Проверяем пустой список после удаления
	log.Println("[INFO] Проверяем пустой список предметов после удаления...")
	disciplines, err = routes.Inf_DisciplinesByGroupData(speciality, groupNum, int(course), int(grads))
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if !compareDisciplines(disciplines, expectedEmpty) {
		log.Printf("[ERROR] Список предметов должен быть пустым, получено=%v", disciplines)
		Bad++
		return Ok, Bad, errors.New("список предметов не пуст после удаления")
	}
	log.Println("[SUCCESS] Пустой список предметов подтверждён")
	Ok++

	// Удаляем группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err = routes.Delete_GroupById(course, grads, speciality, groupNum)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[SUCCESS] Тестовая группа успешно удалена")
	Ok++

	// Проверяем, что предметы удалены после удаления группы
	log.Println("[INFO] Проверяем отсутствие предметов после удаления группы...")
	disciplines, err = routes.Inf_DisciplinesByGroupData(speciality, groupNum, int(course), int(grads))
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if !compareDisciplines(disciplines, expectedEmpty) {
		log.Printf("[ERROR] Предметы не должны существовать после удаления группы, получено=%v", disciplines)
		Bad++
		return Ok, Bad, errors.New("предметы существуют после удаления группы")
	}
	log.Println("[SUCCESS] Предметы отсутствуют после удаления группы")
	Ok++

	// Проверяем ошибку при неверных данных
	log.Println("[INFO] Проверяем ошибку при неверных данных...")
	disciplines, err = routes.Inf_DisciplinesByGroupData("НеИнформатика", 999, 9, 99)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов для неверных данных: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if !compareDisciplines(disciplines, expectedEmpty) {
		log.Printf("[ERROR] Неверные данные должны вернуть пустой список, получено=%v", disciplines)
		Bad++
		return Ok, Bad, errors.New("неверные данные не вернули пустой список")
	}
	log.Println("[SUCCESS] Пустой список получен для неверных данных")
	Ok++

	log.Printf("[INFO] Результат теста: %d успешных, %d неудачных", Ok, Bad)
	return Ok, Bad, nil
}

// Функция для сравнения данных
func compareDisciplines(actual, expected map[int][]string) bool {
	if len(actual) != len(expected) {
		return false
	}
	for key, value := range expected {
		if !equal(value, actual[key]) {
			return false
		}
	}
	return true
}

// Функция для проверки, что два слайса одинаковы
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Вспомогательная функция для проверки списка предметов
func checkSubjects(actual, expected []string, action string) error {
	if !reflect.DeepEqual(actual, expected) {
		return fmt.Errorf("%s: ожидался список предметов %v, получено %v", action, expected, actual)
	}
	return nil
}

// Тест получения списка предметов
func Test_InfSubjectByGroupId() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, graduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupIds, err := routes.Create_Group(course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Группа создана с ID=[%d, %d]", groupIds[0], groupIds[1])
	Ok++
	groupId := groupIds[0] // Используем первый ID группы

	// Проверяем пустой список предметов
	log.Println("[INFO] Проверяем пустой список предметов...")
	subjects, err := routes.Inf_SubjectByGroupId(groupId)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if len(subjects) != 0 {
		log.Printf("[ERROR] Список предметов должен быть пустым, получено=%v", subjects)
		Bad++
		return Ok, Bad, errors.New("список предметов не пуст")
	}
	log.Println("[SUCCESS] Пустой список предметов подтверждён")
	Ok++

	// Добавляем предмет
	log.Println("[INFO] Добавляем предмет в группу...")
	subject := "Математика"
	subjectId, err := routes.Add_SubjectByGroupId(groupId, subject)
	if err != nil {
		log.Printf("[ERROR] Ошибка при добавлении предмета: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Предмет добавлен с ID=%d", subjectId)
	Ok++

	// Проверяем список предметов
	log.Println("[INFO] Проверяем список предметов после добавления...")
	subjects, err = routes.Inf_SubjectByGroupId(groupId)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[DEBUG] Список предметов: %v", subjects)
	if err := checkSubjects(subjects, []string{subject}, "проверка списка предметов"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Println("[SUCCESS] Список предметов успешно получен")
	Ok++

	// Удаляем предмет
	log.Printf("[INFO] Удаляем предмет с ID=%d...", subjectId)
	ok, err := routes.Delete_SubjectById(subjectId)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении предмета: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления предмета")
	}
	log.Println("[SUCCESS] Предмет успешно удалён")
	Ok++

	// Проверяем пустой список после удаления
	log.Println("[INFO] Проверяем пустой список предметов после удаления...")
	subjects, err = routes.Inf_SubjectByGroupId(groupId)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if len(subjects) != 0 {
		log.Printf("[ERROR] Список предметов должен быть пустым, получено=%v", subjects)
		Bad++
		return Ok, Bad, errors.New("список предметов не пуст после удаления")
	}
	log.Println("[SUCCESS] Пустой список предметов подтверждён")
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

	// Проверяем, что предметы удалены после удаления группы
	log.Println("[INFO] Проверяем отсутствие предметов после удаления группы...")
	subjects, _ = routes.Inf_SubjectByGroupId(groupId)
	if len(subjects) != 0 {
		log.Printf("[ERROR - 1] Предметы не должны существовать после удаления группы, получено=%v", subjects)
		Bad++
		return Ok, Bad, errors.New("предметы существуют после удаления группы")
	}
	log.Println("[SUCCESS] Предметы отсутствуют после удаления группы")
	Ok++

	return Ok, Bad, nil
}

// Тест добавления предметов
func Test_AddSubjectByGroupId() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, graduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupIds, err := routes.Create_Group(course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Группа создана с ID=[%d, %d]", groupIds[0], groupIds[1])
	Ok++
	groupId := groupIds[0]

	// Проверяем пустой список предметов
	log.Println("[INFO] Проверяем пустой список предметов...")
	subjects, err := routes.Inf_SubjectByGroupId(groupId)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if len(subjects) != 0 {
		log.Printf("[ERROR] Список предметов должен быть пустым, получено=%v", subjects)
		Bad++
		return Ok, Bad, errors.New("список предметов не пуст")
	}
	log.Println("[SUCCESS] Пустой список предметов подтверждён")
	Ok++

	// Добавляем первый предмет
	log.Println("[INFO] Добавляем первый предмет...")
	subject1 := "Математика"
	subjectId1, err := routes.Add_SubjectByGroupId(groupId, subject1)
	if err != nil {
		log.Printf("[ERROR] Ошибка при добавлении первого предмета: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Первый предмет добавлен с ID=%d", subjectId1)
	Ok++

	// Проверяем список после добавления первого предмета
	log.Println("[INFO] Проверяем список после добавления первого предмета...")
	subjects, err = routes.Inf_SubjectByGroupId(groupId)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[DEBUG] Список предметов: %v", subjects)
	if err := checkSubjects(subjects, []string{subject1}, "проверка списка после добавления первого предмета"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Println("[SUCCESS] Список после добавления первого предмета подтверждён")
	Ok++

	// Добавляем второй предмет
	log.Println("[INFO] Добавляем второй предмет...")
	subject2 := "Физика"
	subjectId2, err := routes.Add_SubjectByGroupId(groupId, subject2)
	if err != nil {
		log.Printf("[ERROR] Ошибка при добавлении второго предмета: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Второй предмет добавлен с ID=%d", subjectId2)
	Ok++

	// Проверяем список после добавления второго предмета
	log.Println("[INFO] Проверяем список после добавления второго предмета...")
	subjects, err = routes.Inf_SubjectByGroupId(groupId)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[DEBUG] Список предметов: %v", subjects)
	if err := checkSubjects(subjects, []string{subject1, subject2}, "проверка списка после добавления второго предмета"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Println("[SUCCESS] Список после добавления второго предмета подтверждён")
	Ok++

	// Удаляем группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err := routes.Delete_GroupById(course, graduates, speciality, groupNum)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[SUCCESS] Тестовая группа успешно удалена")
	Ok++

	// Проверяем, что предметы удалены после удаления группы
	log.Println("[INFO] Проверяем отсутствие предметов после удаления группы...")
	subjects, _ = routes.Inf_SubjectByGroupId(groupId)
	if len(subjects) != 0 {
		log.Printf("[ERROR - 2] Предметы не должны существовать после удаления группы, получено=%v", subjects)
		Bad++
		return Ok, Bad, errors.New("предметы существуют после удаления группы")
	}
	log.Println("[SUCCESS] Предметы отсутствуют после удаления группы")
	Ok++

	return Ok, Bad, nil
}

// Тест обновления предмета
func Test_UpdateSubjectById() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, graduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupIds, err := routes.Create_Group(course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Группа создана с ID=[%d, %d]", groupIds[0], groupIds[1])
	Ok++
	groupId := groupIds[0]

	// Добавляем предмет
	log.Println("[INFO] Добавляем предмет для обновления...")
	subject := "Математика"
	subjectId, err := routes.Add_SubjectByGroupId(groupId, subject)
	if err != nil {
		log.Printf("[ERROR] Ошибка при добавлении предмета: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Предмет добавлен с ID=%d", subjectId)
	Ok++

	// Проверяем список предметов
	log.Println("[INFO] Проверяем список предметов перед обновлением...")
	subjects, err := routes.Inf_SubjectByGroupId(groupId)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[DEBUG] Список предметов: %v", subjects)
	if err := checkSubjects(subjects, []string{subject}, "проверка списка перед обновлением"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Println("[SUCCESS] Список предметов перед обновлением подтверждён")
	Ok++

	// Обновляем предмет
	log.Println("[INFO] Обновляем предмет...")
	newSubject := "Высшая математика"
	ok, err := routes.Update_SubjectById(subjectId, newSubject)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при обновлении предмета: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка обновления предмета")
	}
	log.Println("[SUCCESS] Предмет успешно обновлён")
	Ok++

	// Проверяем список после обновления
	log.Println("[INFO] Проверяем список после обновления предмета...")
	subjects, err = routes.Inf_SubjectByGroupId(groupId)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[DEBUG] Список предметов: %v", subjects)
	if err := checkSubjects(subjects, []string{newSubject}, "проверка списка после обновления"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Println("[SUCCESS] Список после обновления подтверждён")
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

// Тест удаления одного предмета
func Test_DeleteSubjectById() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, graduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupIds, err := routes.Create_Group(course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Группа создана с ID=[%d, %d]", groupIds[0], groupIds[1])
	Ok++
	groupId := groupIds[0]

	// Добавляем предмет
	log.Println("[INFO] Добавляем предмет для удаления...")
	subject := "Математика"
	subjectId, err := routes.Add_SubjectByGroupId(groupId, subject)
	if err != nil {
		log.Printf("[ERROR] Ошибка при добавлении предмета: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Предмет добавлен с ID=%d", subjectId)
	Ok++

	// Проверяем список предметов
	log.Println("[INFO] Проверяем список предметов перед удалением...")
	subjects, err := routes.Inf_SubjectByGroupId(groupId)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[DEBUG] Список предметов: %v", subjects)
	if err := checkSubjects(subjects, []string{subject}, "проверка списка перед удалением"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Println("[SUCCESS] Список предметов перед удалением подтверждён")
	Ok++

	// Удаляем предмет
	log.Printf("[INFO] Удаляем предмет с ID=%d...", subjectId)
	ok, err := routes.Delete_SubjectById(subjectId)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении предмета: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления предмета")
	}
	log.Println("[SUCCESS] Предмет успешно удалён")
	Ok++

	// Проверяем пустой список после удаления
	log.Println("[INFO] Проверяем пустой список после удаления предмета...")
	subjects, err = routes.Inf_SubjectByGroupId(groupId)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if len(subjects) != 0 {
		log.Printf("[ERROR] Список предметов должен быть пустым, получено=%v", subjects)
		Bad++
		return Ok, Bad, errors.New("список предметов не пуст после удаления")
	}
	log.Println("[SUCCESS] Пустой список после удаления подтверждён")
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

// Тест удаления всех предметов группы
func Test_DeleteAllSubjectByGroupId() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, graduates, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupIds, err := routes.Create_Group(course, graduates, speciality, groupNum)
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Группа создана с ID=[%d, %d]", groupIds[0], groupIds[1])
	Ok++
	groupId := groupIds[0]

	// Добавляем два предмета
	log.Println("[INFO] Добавляем первый предмет...")
	subject1 := "Математика"
	subjectId1, err := routes.Add_SubjectByGroupId(groupId, subject1)
	if err != nil {
		log.Printf("[ERROR] Ошибка при добавлении первого предмета: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Первый предмет добавлен с ID=%d", subjectId1)
	Ok++

	log.Println("[INFO] Добавляем второй предмет...")
	subject2 := "Физика"
	subjectId2, err := routes.Add_SubjectByGroupId(groupId, subject2)
	if err != nil {
		log.Printf("[ERROR] Ошибка при добавлении второго предмета: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Второй предмет добавлен с ID=%d", subjectId2)
	Ok++

	// Проверяем список предметов
	log.Println("[INFO] Проверяем список предметов перед удалением...")
	subjects, err := routes.Inf_SubjectByGroupId(groupId)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[DEBUG] Список предметов: %v", subjects)
	if err := checkSubjects(subjects, []string{subject1, subject2}, "проверка списка перед удалением всех предметов"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Println("[SUCCESS] Список предметов перед удалением подтверждён")
	Ok++

	// Удаляем все предметы
	log.Printf("[INFO] Удаляем все предметы для группы с ID=%d...", groupId)
	ok, err := routes.Delete_AllSubjectByGroupId(groupId)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении всех предметов: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления всех предметов")
	}
	log.Println("[SUCCESS] Все предметы успешно удалены")
	Ok++

	// Проверяем пустой список после удаления
	log.Println("[INFO] Проверяем пустой список после удаления всех предметов...")
	subjects, err = routes.Inf_SubjectByGroupId(groupId)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if len(subjects) != 0 {
		log.Printf("[ERROR] Список предметов должен быть пустым, получено=%v", subjects)
		Bad++
		return Ok, Bad, errors.New("список предметов не пуст после удаления всех предметов")
	}
	log.Println("[SUCCESS] Пустой список после удаления всех предметов подтверждён")
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

// Тест всех функций Subjects
func TestSubjectsALL() (int, int) {
	var Ok, Bad int

	// Тестирование Inf_SubjectByGroupId
	ok, bad, err := Test_InfSubjectByGroupId()
	if err != nil {
		log.Fatal("Ошибка в Test_InfSubjectByGroupId: ", err)
	}
	logResults("InfSubjectByGroupId", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование Add_SubjectByGroupId
	ok, bad, err = Test_AddSubjectByGroupId()
	if err != nil {
		log.Fatal("Ошибка в Test_AddSubjectByGroupId: ", err)
	}
	logResults("AddSubjectByGroupId", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование Update_SubjectById
	ok, bad, err = Test_UpdateSubjectById()
	if err != nil {
		log.Fatal("Ошибка в Test_UpdateSubjectById: ", err)
	}
	logResults("UpdateSubjectById", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование Delete_SubjectById
	ok, bad, err = Test_DeleteSubjectById()
	if err != nil {
		log.Fatal("Ошибка в Test_DeleteSubjectById: ", err)
	}
	logResults("DeleteSubjectById", ok, bad, err)
	Ok += ok
	Bad += bad

	// Тестирование Delete_AllSubjectByGroupId
	ok, bad, err = Test_DeleteAllSubjectByGroupId()
	if err != nil {
		log.Fatal("Ошибка в Test_DeleteAllSubjectByGroupId: ", err)
	}
	logResults("DeleteAllSubjectByGroupId", ok, bad, err)
	Ok += ok
	Bad += bad

	ok, bad, err = Test_Inf_DisciplinesByGroupData()
	if err != nil {
		log.Fatal("Ошибка в Test_Inf_DisciplinesByGroupData: ", err)
	}
	logResults("Test_Inf_DisciplinesByGroupData", ok, bad, err)
	Ok += ok
	Bad += bad

	// Итоговый вывод
	log.Println("\n[SUMMARY] Итоговые результаты:")
	log.Printf("TestSubjectsALL [%d\\%d]\nok: %d\nbad: %d\n", Ok, Bad, Ok, Bad)

	return Ok, Bad
}
