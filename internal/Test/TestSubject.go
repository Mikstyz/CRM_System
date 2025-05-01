package Test

import (
	"CRM_System/internal/models"
	"CRM_System/internal/routes"
	"errors"
	"fmt"
	"log"
	"reflect"
	"sort"
)

// Test_Inf_DisciplinesByGroupData тестирует получение дисциплин по данным группы.
// Test_Inf_DisciplinesByGroupData тестирует получение дисциплин по данным группы.
// Test_Inf_DisciplinesByGroupData тестирует получение списка предметов по данным группы.
func Test_Inf_DisciplinesByGroupData() (int, int, error) {
	var Ok, Bad int

	// Создаём группу для теста
	log.Println("[INFO] Создаём тестовую группу...")
	course, grads, speciality, groupNum := byte(1), byte(4), "Информатика", 101
	groupId, err := routes.Create_Group(course, grads, speciality, groupNum)
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании группы: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Группа создана с ID=%d", groupId)
	Ok++

	// Проверяем пустой список предметов
	log.Println("[INFO] Проверяем пустой список предметов...")
	disciplines, err := routes.Inf_DisciplinesByGroupData(speciality, groupNum, int(course), int(grads))
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	expectedEmpty := models.DisciplinesBySemester{
		FirstSemester:  []models.SemesterDiscipline{},
		SecondSemester: []models.SemesterDiscipline{},
	}
	// Проверяем, что оба слайса либо nil, либо пустые
	if !isEmptyDisciplinesBySemester(disciplines, expectedEmpty) {
		log.Printf("[ERROR] Список предметов должен быть пустым, получено=%v", disciplines)
		Bad++
		return Ok, Bad, errors.New("список предметов не пуст")
	}
	log.Println("[SUCCESS] Пустой список предметов подтверждён")
	Ok++

	// Добавляем предметы для семестра 1
	log.Println("[INFO] Добавляем предметы для семестра 1...")
	subjectsSem1 := []string{"Математика", "Алгебра"}
	var subjectIdsSem1 []int
	for _, subject := range subjectsSem1 {
		subjectId, err := routes.Add_SubjectByGroupId(groupId, subject, 1)
		if err != nil {
			log.Printf("[ERROR] Ошибка при добавлении предмета '%s': %v", subject, err)
			Bad++
			return Ok, Bad, err
		}
		log.Printf("[SUCCESS] Предмет '%s' добавлен с ID=%d", subject, subjectId)
		subjectIdsSem1 = append(subjectIdsSem1, subjectId)
		Ok++
	}

	// Добавляем предметы для семестра 2
	log.Println("[INFO] Добавляем предметы для семестра 2...")
	subjectsSem2 := []string{"Программирование", "Базы данных"}
	var subjectIdsSem2 []int
	for _, subject := range subjectsSem2 {
		subjectId, err := routes.Add_SubjectByGroupId(groupId, subject, 2)
		if err != nil {
			log.Printf("[ERROR] Ошибка при добавлении предмета '%s': %v", subject, err)
			Bad++
			return Ok, Bad, err
		}
		log.Printf("[SUCCESS] Предмет '%s' добавлен с ID=%d", subject, subjectId)
		subjectIdsSem2 = append(subjectIdsSem2, subjectId)
		Ok++
	}

	// Проверяем список предметов после добавления
	log.Println("[INFO] Проверяем список предметов после добавления...")
	disciplines, err = routes.Inf_DisciplinesByGroupData(speciality, groupNum, int(course), int(grads))
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[DEBUG] Список предметов: %v", disciplines)

	// Извлекаем только Title из полученных данных
	var actualFirstSemTitles, actualSecondSemTitles []string
	for _, subj := range disciplines.FirstSemester {
		actualFirstSemTitles = append(actualFirstSemTitles, subj.Title)
	}
	for _, subj := range disciplines.SecondSemester {
		actualSecondSemTitles = append(actualSecondSemTitles, subj.Title)
	}

	// Сортируем ожидаемые Title в алфавитном порядке, как в базе
	sortedSubjectsSem1 := make([]string, len(subjectsSem1))
	copy(sortedSubjectsSem1, subjectsSem1)
	sort.Strings(sortedSubjectsSem1)
	sortedSubjectsSem2 := make([]string, len(subjectsSem2))
	copy(sortedSubjectsSem2, subjectsSem2)
	sort.Strings(sortedSubjectsSem2)

	// Сравниваем только Title
	if !reflect.DeepEqual(actualFirstSemTitles, sortedSubjectsSem1) {
		log.Printf("[ERROR] Первый семестр: ожидалось %v, получено %v", sortedSubjectsSem1, actualFirstSemTitles)
		Bad++
		return Ok, Bad, errors.New("список предметов первого семестра не совпадает")
	}
	if !reflect.DeepEqual(actualSecondSemTitles, sortedSubjectsSem2) {
		log.Printf("[ERROR] Второй семестр: ожидалось %v, получено %v", sortedSubjectsSem2, actualSecondSemTitles)
		Bad++
		return Ok, Bad, errors.New("список предметов второго семестра не совпадает")
	}
	log.Println("[SUCCESS] Данные предметов корректны")
	Ok++

	// Удаляем предметы
	for i, subjectId := range subjectIdsSem1 {
		log.Printf("[INFO] Удаляем предмет с ID=%d...", subjectId)
		ok, err := routes.Delete_SubjectById(subjectId)
		if err != nil || !ok {
			log.Printf("[ERROR] Ошибка при удалении предмета '%s': %v", subjectsSem1[i], err)
			Bad++
			return Ok, Bad, errors.New("ошибка удаления предмета")
		}
		log.Printf("[SUCCESS] Предмет '%s' успешно удалён", subjectsSem1[i])
		Ok++
	}
	for i, subjectId := range subjectIdsSem2 {
		log.Printf("[INFO] Удаляем предмет с ID=%d...", subjectId)
		ok, err := routes.Delete_SubjectById(subjectId)
		if err != nil || !ok {
			log.Printf("[ERROR] Ошибка при удалении предмета '%s': %v", subjectsSem2[i], err)
			Bad++
			return Ok, Bad, errors.New("ошибка удаления предмета")
		}
		log.Printf("[SUCCESS] Предмет '%s' успешно удалён", subjectsSem2[i])
		Ok++
	}

	// Проверяем пустой список после удаления
	log.Println("[INFO] Проверяем пустой список предметов после удаления...")
	disciplines, err = routes.Inf_DisciplinesByGroupData(speciality, groupNum, int(course), int(grads))
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if !isEmptyDisciplinesBySemester(disciplines, expectedEmpty) {
		log.Printf("[ERROR] Список предметов должен быть пустым, получено=%v", disciplines)
		Bad++
		return Ok, Bad, errors.New("список предметов не пуст после удаления")
	}
	log.Println("[SUCCESS] Пустой список предметов подтверждён")
	Ok++

	// Удаляем группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err := routes.Delete_GroupById(groupId)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[SUCCESS] Тестовая группа успешно удалена")
	Ok++

	// Проверяем отсутствие предметов после удаления группы
	log.Println("[INFO] Проверяем отсутствие предметов после удаления группы...")
	disciplines, err = routes.Inf_DisciplinesByGroupData(speciality, groupNum, int(course), int(grads))
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if !isEmptyDisciplinesBySemester(disciplines, expectedEmpty) {
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
	if !isEmptyDisciplinesBySemester(disciplines, expectedEmpty) {
		log.Printf("[ERROR] Неверные данные должны вернуть пустой список, получено=%v", disciplines)
		Bad++
		return Ok, Bad, errors.New("неверные данные не вернули пустой список")
	}
	log.Println("[SUCCESS] Пустой список получен для неверных данных")
	Ok++

	log.Printf("[INFO] Результат теста: %d успешных, %d неудачных", Ok, Bad)
	return Ok, Bad, nil
}

// isEmptyDisciplinesBySemester проверяет, что DisciplinesBySemester пуст (оба слайса nil или пустые).
func isEmptyDisciplinesBySemester(actual, expected models.DisciplinesBySemester) bool {
	isFirstEmpty := (actual.FirstSemester == nil || len(actual.FirstSemester) == 0) &&
		(expected.FirstSemester == nil || len(expected.FirstSemester) == 0)
	isSecondEmpty := (actual.SecondSemester == nil || len(actual.SecondSemester) == 0) &&
		(expected.SecondSemester == nil || len(expected.SecondSemester) == 0)
	return isFirstEmpty && isSecondEmpty
}

// Test_InfSubjectByGroupId тестирует получение списка предметов по ID группы и семестру.
func Test_InfSubjectByGroupId() (int, int, error) {
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

	// Проверяем пустой список предметов для 1-го семестра
	log.Println("[INFO] Проверяем пустой список предметов для 1-го семестра...")
	subjects, err := routes.Inf_SubjectByGroupId(groupId, 1)
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
	log.Println("[SUCCESS] Пустой список предметов для 1-го семестра подтверждён")
	Ok++

	// Добавляем предметы для 1-го семестра
	log.Println("[INFO] Добавляем предметы в 1-й семестр...")
	subjectsSem1 := []string{"Математика", "Алгебра"}
	var subjectIdsSem1 []int
	for _, subject := range subjectsSem1 {
		subjectId, err := routes.Add_SubjectByGroupId(groupId, subject, 1)
		if err != nil {
			log.Printf("[ERROR] Ошибка при добавлении предмета '%s': %v", subject, err)
			Bad++
			return Ok, Bad, err
		}
		log.Printf("[SUCCESS] Предмет '%s' добавлен с ID=%d", subject, subjectId)
		subjectIdsSem1 = append(subjectIdsSem1, subjectId)
		Ok++
	}

	// Проверяем список предметов для 1-го семестра
	log.Println("[INFO] Проверяем список предметов для 1-го семестра после добавления...")
	subjects, err = routes.Inf_SubjectByGroupId(groupId, 1)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[DEBUG] Список предметов: %v", subjects)
	if err := checkSubjects(subjects, subjectsSem1, "проверка списка предметов для 1-го семестра"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Println("[SUCCESS] Список предметов для 1-го семестра успешно получен")
	Ok++

	// Удаляем предметы
	for i, subjectId := range subjectIdsSem1 {
		log.Printf("[INFO] Удаляем предмет с ID=%d...", subjectId)
		ok, err := routes.Delete_SubjectById(subjectId)
		if err != nil || !ok {
			log.Printf("[ERROR] Ошибка при удалении предмета '%s': %v", subjectsSem1[i], err)
			Bad++
			return Ok, Bad, errors.New("ошибка удаления предмета")
		}
		log.Printf("[SUCCESS] Предмет '%s' успешно удалён", subjectsSem1[i])
		Ok++
	}

	// Проверяем пустой список после удаления
	log.Println("[INFO] Проверяем пустой список предметов для 1-го семестра после удаления...")
	subjects, err = routes.Inf_SubjectByGroupId(groupId, 1)
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
	log.Println("[SUCCESS] Пустой список предметов для 1-го семестра подтверждён")
	Ok++

	// Удаляем группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err := routes.Delete_GroupById(groupId)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[SUCCESS] Тестовая группа успешно удалена")
	Ok++

	// Проверяем отсутствие предметов после удаления группы
	log.Println("[INFO] Проверяем отсутствие предметов после удаления группы...")
	subjects, err = routes.Inf_SubjectByGroupId(groupId, 1)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if len(subjects) != 0 {
		log.Printf("[ERROR] Предметы не должны существовать после удаления группы, получено=%v", subjects)
		Bad++
		return Ok, Bad, errors.New("предметы существуют после удаления группы")
	}
	log.Println("[SUCCESS] Предметы отсутствуют после удаления группы")
	Ok++

	return Ok, Bad, nil
}

// Test_AddSubjectByGroupId тестирует добавление предметов в группу.
func Test_AddSubjectByGroupId() (int, int, error) {
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

	// Проверяем пустой список предметов
	log.Println("[INFO] Проверяем пустой список предметов для 1-го семестра...")
	subjects, err := routes.Inf_SubjectByGroupId(groupId, 1)
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
	log.Println("[INFO] Добавляем первый предмет для 1-го семестра...")
	subject1 := "Математика"
	subjectId1, err := routes.Add_SubjectByGroupId(groupId, subject1, 1)
	if err != nil {
		log.Printf("[ERROR] Ошибка при добавлении первого предмета: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Первый предмет добавлен с ID=%d", subjectId1)
	Ok++

	// Проверяем список после добавления первого предмета
	log.Println("[INFO] Проверяем список после добавления первого предмета...")
	subjects, err = routes.Inf_SubjectByGroupId(groupId, 1)
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
	log.Println("[INFO] Добавляем второй предмет для 1-го семестра...")
	subject2 := "Физика"
	subjectId2, err := routes.Add_SubjectByGroupId(groupId, subject2, 1)
	if err != nil {
		log.Printf("[ERROR] Ошибка при добавлении второго предмета: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Второй предмет добавлен с ID=%d", subjectId2)
	Ok++

	// Проверяем список после добавления второго предмета
	log.Println("[INFO] Проверяем список после добавления второго предмета...")
	subjects, err = routes.Inf_SubjectByGroupId(groupId, 1)
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
	ok, err := routes.Delete_GroupById(groupId)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[SUCCESS] Тестовая группа успешно удалена")
	Ok++

	// Проверяем отсутствие предметов после удаления группы
	log.Println("[INFO] Проверяем отсутствие предметов после удаления группы...")
	subjects, err = routes.Inf_SubjectByGroupId(groupId, 1)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if len(subjects) != 0 {
		log.Printf("[ERROR] Предметы не должны существовать после удаления группы, получено=%v", subjects)
		Bad++
		return Ok, Bad, errors.New("предметы существуют после удаления группы")
	}
	log.Println("[SUCCESS] Предметы отсутствуют после удаления группы")
	Ok++

	return Ok, Bad, nil
}

// Test_UpdateSubjectById тестирует обновление предмета.
func Test_UpdateSubjectById() (int, int, error) {
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

	// Добавляем предмет
	log.Println("[INFO] Добавляем предмет для обновления...")
	subject := "Математика"
	subjectId, err := routes.Add_SubjectByGroupId(groupId, subject, 1)
	if err != nil {
		log.Printf("[ERROR] Ошибка при добавлении предмета: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Предмет добавлен с ID=%d", subjectId)
	Ok++

	// Проверяем список предметов
	log.Println("[INFO] Проверяем список предметов перед обновлением...")
	subjects, err := routes.Inf_SubjectByGroupId(groupId, 1)
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
	subjects, err = routes.Inf_SubjectByGroupId(groupId, 1)
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
	ok, err = routes.Delete_GroupById(groupId)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[SUCCESS] Тестовая группа успешно удалена")
	Ok++

	return Ok, Bad, nil
}

// Test_DeleteSubjectById тестирует удаление одного предмета.
func Test_DeleteSubjectById() (int, int, error) {
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

	// Добавляем предмет
	log.Println("[INFO] Добавляем предмет для удаления...")
	subject := "Математика"
	subjectId, err := routes.Add_SubjectByGroupId(groupId, subject, 1)
	if err != nil {
		log.Printf("[ERROR] Ошибка при добавлении предмета: %v", err)
		Bad++
		return Ok, Bad, err
	}
	log.Printf("[SUCCESS] Предмет добавлен с ID=%d", subjectId)
	Ok++

	// Проверяем список предметов
	log.Println("[INFO] Проверяем список предметов перед удалением...")
	subjects, err := routes.Inf_SubjectByGroupId(groupId, 1)
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
	subjects, err = routes.Inf_SubjectByGroupId(groupId, 1)
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
	ok, err = routes.Delete_GroupById(groupId)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[SUCCESS] Тестовая группа успешно удалена")
	Ok++

	return Ok, Bad, nil
}

// Test_DeleteAllSubjectByGroupId тестирует удаление всех предметов группы.
func Test_DeleteAllSubjectByGroupId() (int, int, error) {
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

	// Добавляем предметы
	log.Println("[INFO] Добавляем предметы для 1-го семестра...")
	subjectsSem1 := []string{"Математика", "Физика"}
	var subjectIdsSem1 []int
	for _, subject := range subjectsSem1 {
		subjectId, err := routes.Add_SubjectByGroupId(groupId, subject, 1)
		if err != nil {
			log.Printf("[ERROR] Ошибка при добавлении предмета '%s': %v", subject, err)
			Bad++
			return Ok, Bad, err
		}
		log.Printf("[SUCCESS] Предмет '%s' добавлен с ID=%d", subject, subjectId)
		subjectIdsSem1 = append(subjectIdsSem1, subjectId)
		Ok++
	}

	log.Println("[INFO] Добавляем предметы для 2-го семестра...")
	subjectsSem2 := []string{"Программирование", "Базы данных"}
	var subjectIdsSem2 []int
	for _, subject := range subjectsSem2 {
		subjectId, err := routes.Add_SubjectByGroupId(groupId, subject, 2)
		if err != nil {
			log.Printf("[ERROR] Ошибка при добавлении предмета '%s': %v", subject, err)
			Bad++
			return Ok, Bad, err
		}
		log.Printf("[SUCCESS] Предмет '%s' добавлен с ID=%d", subject, subjectId)
		subjectIdsSem2 = append(subjectIdsSem2, subjectId)
		Ok++
	}

	// Проверяем список предметов перед удалением
	log.Println("[INFO] Проверяем список предметов перед удалением...")
	subjectsSem1Check, err := routes.Inf_SubjectByGroupId(groupId, 1)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов для 1-го семестра: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if err := checkSubjects(subjectsSem1Check, subjectsSem1, "проверка списка предметов для 1-го семестра"); err != nil {
		log.Printf("[ERROR] Ошибка: %v", err)
		Bad++
		return Ok, Bad, err
	}
	subjectsSem2Check, err := routes.Inf_SubjectByGroupId(groupId, 2)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов для 2-го семестра: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if err := checkSubjects(subjectsSem2Check, subjectsSem2, "проверка списка предметов для 2-го семестра"); err != nil {
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
	subjects, err := routes.Inf_SubjectByGroupId(groupId, 1)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов для 1-го семестра: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if len(subjects) != 0 {
		log.Printf("[ERROR] Список предметов для 1-го семестра должен быть пустым, получено=%v", subjects)
		Bad++
		return Ok, Bad, errors.New("список предметов не пуст после удаления всех предметов")
	}
	subjects, err = routes.Inf_SubjectByGroupId(groupId, 2)
	if err != nil {
		log.Printf("[ERROR] Ошибка при получении предметов для 2-го семестра: %v", err)
		Bad++
		return Ok, Bad, err
	}
	if len(subjects) != 0 {
		log.Printf("[ERROR] Список предметов для 2-го семестра должен быть пустым, получено=%v", subjects)
		Bad++
		return Ok, Bad, errors.New("список предметов не пуст после удаления всех предметов")
	}
	log.Println("[SUCCESS] Пустой список после удаления всех предметов подтверждён")
	Ok++

	// Удаляем группу
	log.Println("[INFO] Удаляем тестовую группу...")
	ok, err = routes.Delete_GroupById(groupId)
	if err != nil || !ok {
		log.Printf("[ERROR] Ошибка при удалении тестовой группы: %v", err)
		Bad++
		return Ok, Bad, errors.New("ошибка удаления тестовой группы")
	}
	log.Println("[SUCCESS] Тестовая группа успешно удалена")
	Ok++

	return Ok, Bad, nil
}

// TestSubjectsALL тестирует все функции Subjects.
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

	// Тестирование Inf_DisciplinesByGroupData
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

// checkSubjects проверяет, что два слайса одинаковы.
func checkSubjects(actual, expected []string, action string) error {
	if !reflect.DeepEqual(actual, expected) {
		return fmt.Errorf("%s: ожидался список предметов %v, получено %v", action, expected, actual)
	}
	return nil
}

// logResults логирует результаты теста.
