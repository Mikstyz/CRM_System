package repository

import (
	db "CRM_System/Backend/internal/db"
	"log"
)

func InfDiscipAnDock(groupId int) ([]string, error) {
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

func InfDisciplinesInGroup(groupId int) ([]string, error) {
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
