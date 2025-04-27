package repository

import (
	db "CRM_System/internal/db"
	"log"
)

func InfDisciplinesByGroup(groupId int) ([]string, error) {
	log.Println("Получение предметов группы")
	const query = `SELECT subject_name FROM group_subject WHERE group_id = ?`

	db.Init()

	rows, err := db.DB.Query(query, groupId)
	if err != nil {
		log.Printf("Ошибка при получении предметов группы: %v", err)
		return nil, err
	}
	defer rows.Close()

	var disciplines []string
	for rows.Next() {
		var subjectName string
		if err := rows.Scan(&subjectName); err != nil {
			log.Printf("Ошибка при сканировании предметов группы: %v", err)
			continue
		}
		disciplines = append(disciplines, subjectName)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Ошибка при чтении предметов группы: %v", err)
		return nil, err
	}

	log.Printf("Получено предметов: %d", len(disciplines))

	return disciplines, nil
}

func InfDisciplinesByGroupData(Speciality string, GroupNum int, Course int, Groudates int) (map[int][]string, error) {
	log.Printf("Получение предметов для группы: Speciality=%s, GroupNum=%d, Course=%d, Groudates=%d", Speciality, GroupNum, Course, Groudates)

	// Запрос для получения предметов и семестров
	const query = `
		SELECT gs.subject_name, eg.Semester 
		FROM group_subject gs
		JOIN einf_groups eg ON gs.group_id = eg.Id
		WHERE eg.Speciality = ? AND eg.GroupNum = ? AND eg.Course = ? AND eg.Groudates = ?`

	// Инициализация соединения с базой данных
	db.Init()

	rows, err := db.DB.Query(query, Speciality, GroupNum, Course, Groudates)
	if err != nil {
		log.Printf("Ошибка при получении предметов: %v", err)
		return nil, err
	}
	defer rows.Close()

	// Инициализация мапы для семестров
	disciplines := make(map[int][]string)

	// Обрабатываем данные
	for rows.Next() {
		var subjectName string
		var semester int
		if err := rows.Scan(&subjectName, &semester); err != nil {
			log.Printf("Ошибка при сканировании предмета: %v", err)
			continue
		}
		disciplines[semester] = append(disciplines[semester], subjectName)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Ошибка при чтении данных: %v", err)
		return nil, err
	}

	// Если мапа пустая, инициализируем пустые слайсы для семестров 1 и 2
	if len(disciplines) == 0 {
		disciplines[1] = []string{}
		disciplines[2] = []string{}
	}

	log.Printf("Получено предметов: семестр 1 — %d, семестр 2 — %d", len(disciplines[1]), len(disciplines[2]))
	return disciplines, nil
}

func AddDisciplinesInGroup(groupId int, newSubject string) (int, error) {
	log.Println("Добавление предмета в группу")
	const query = `INSERT INTO group_subject (group_id, subject_name) VALUES (?, ?) RETURNING id`

	// Инициализация подключения к базе данных, если нужно
	db.Init()

	var newId int
	err := db.DB.QueryRow(query, groupId, newSubject).Scan(&newId)
	if err != nil {
		log.Printf("Ошибка добавления предмета в группу: %v", err)
		return 0, err
	}

	log.Printf("Успешное создание предмета с ID: %d", newId)
	return newId, nil
}

func UpdateDisciplinesInGroup(subjectId int, newSubject string) (bool, error) {
	log.Println("Редактирование предмета у группы")
	const query = `UPDATE group_subject SET subject_name = ? WHERE id = ?`

	// Инициализация подключения к базе данных, если нужно
	db.Init()

	res, err := db.DB.Exec(query, newSubject, subjectId)
	if err != nil {
		log.Printf("Ошибка редактирования предмета у группы: %v", err)
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при проверке затронутых строк: %v", err)
		return false, err
	}

	if rowsAffected == 0 {
		log.Println("Предмет не найден")
		return false, nil
	}

	log.Printf("Затронуто записей: %d", rowsAffected)
	log.Println("Успешное обновление данных предмета")
	return true, nil
}

func DeleteDisciplinesInGroup(subjectId int) (bool, error) {
	log.Println("Удаление предмета у группы")
	const query = `DELETE FROM group_subject WHERE id = ?`

	// Инициализация подключения к базе данных, если нужно
	db.Init()

	res, err := db.DB.Exec(query, subjectId)
	if err != nil {
		log.Printf("Ошибка при удалении предмета у группы: %v", err)
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при проверке затронутых строк: %v", err)
		return false, err
	}

	if rowsAffected == 0 {
		log.Println("Предмет не найден или уже удален")
		return false, nil
	}

	log.Printf("Затронуто записей: %d", rowsAffected)
	log.Println("Успешное удаление предмета")
	return true, nil
}

func DeleteDisciplinesByGroupId(groupID int) (bool, error) {
	const query = `DELETE FROM group_subject WHERE group_id = ?`

	db.Init()

	result, err := db.DB.Exec(query, groupID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func CopyDisciplinesBetweenGroups(oldGroupId int, newGroupId int) error {
	const query = `
		INSERT INTO group_subject (group_id, subject_name)
		SELECT ?, subject_name
		FROM group_subject
		WHERE group_id = ?;
	`

	db.Init()

	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, newGroupId, oldGroupId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
