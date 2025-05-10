package repository

import (
	db "CRM_System/internal/db"
	"CRM_System/internal/models"
	"context"
	"database/sql"
	"log"
	"time"
)

// ----------------------information--------------------------
// InfDisciplinesByGroup возвращает список всех предметов для группы по её ID.
func InfDisciplinesByGroup(groupId int, semester byte) ([]string, error) {
	log.Println("Получение предметов группы")
	const query = `SELECT subject_name FROM group_subject WHERE group_id = ? AND semester = ?`

	db.Init()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.DB.QueryContext(ctx, query, groupId, semester)
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
			return nil, err
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

// InfDisciplinesByGroupData возвращает дисциплины для двух семестров по данным группы.
func InfDisciplinesByGroupData(Speciality string, GroupNum int, Course int, Groudates int) (models.DisciplinesBySemester, error) {
	log.Printf("Получение предметов для группы: Speciality=%s, GroupNum=%d, Course=%d, Groudates=%d", Speciality, GroupNum, Course, Groudates)

	const query = `
		SELECT gs.subject_name, gs.semester 
		FROM group_subject gs 
		JOIN einf_groups eg ON gs.group_id = eg.Id 
		WHERE eg.Speciality = ? AND eg.GroupNum = ? AND eg.Course = ? AND eg.Groudates = ? 
		ORDER BY gs.semester, gs.subject_name
	`

	db.Init()

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Ошибка при подготовке запроса: %v", err)
		return models.DisciplinesBySemester{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(Speciality, GroupNum, Course, Groudates)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return models.DisciplinesBySemester{}, err
	}
	defer rows.Close()

	var result models.DisciplinesBySemester

	for rows.Next() {
		var subjectName string
		var semester int
		if err := rows.Scan(&subjectName, &semester); err != nil {
			log.Printf("Ошибка при сканировании предмета: %v", err)
			return models.DisciplinesBySemester{}, err
		}

		discipline := models.SemesterDiscipline{
			Id:    semester,
			Title: subjectName,
		}

		switch semester {
		case 1:
			result.FirstSemester = append(result.FirstSemester, discipline)
		case 2:
			result.SecondSemester = append(result.SecondSemester, discipline)
		default:
			log.Printf("Неизвестный семестр: %d для предмета %s", semester, subjectName)
		}
	}

	if err := rows.Err(); err != nil {
		log.Printf("Ошибка при чтении данных: %v", err)
		return models.DisciplinesBySemester{}, err
	}

	log.Printf("Получено предметов: семестр 1 — %d шт., семестр 2 — %d шт.", len(result.FirstSemester), len(result.SecondSemester))
	return result, nil
}

// ----------------------Manager--------------------------
// AddDisciplinesInGroup добавляет новый предмет в группу.
func AddDisciplinesInGroup(groupId int, newSubject string, semester byte) (int, error) {
	log.Println("Добавление предмета в группу")
	const query = `INSERT INTO group_subject (group_id, subject_name, semester) VALUES (?, ?, ?) RETURNING id`

	db.Init()

	tx, err := db.DB.Begin()
	if err != nil {
		log.Printf("Ошибка при старте транзакции: %v", err)
		return 0, err
	}

	var newId int
	err = tx.QueryRow(query, groupId, newSubject, semester).Scan(&newId)
	if err != nil {
		tx.Rollback()
		log.Printf("Ошибка добавления предмета в группу: %v", err)
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Ошибка при коммите транзакции: %v", err)
		return 0, err
	}

	log.Printf("Успешное создание предмета с ID: %d", newId)
	return newId, nil
}

// UpdateDisciplinesInGroup обновляет предмет в группе.
func UpdateDisciplinesInGroup(subjectId int, newSubject string) (bool, error) {
	log.Println("Редактирование предмета у группы")
	const query = `UPDATE group_subject SET subject_name = ? WHERE id = ? RETURNING id`

	db.Init()

	var updatedId int
	err := db.DB.QueryRow(query, newSubject, subjectId).Scan(&updatedId)
	if err == sql.ErrNoRows {
		log.Println("Предмет не найден")
		return false, nil
	}
	if err != nil {
		log.Printf("Ошибка редактирования предмета у группы: %v", err)
		return false, err
	}

	log.Println("Успешное обновление данных предмета")
	return true, nil
}

// DeleteDisciplinesInGroup удаляет предмет из группы.
func DeleteDisciplinesInGroup(subjectId int) (bool, error) {
	log.Println("Удаление предмета у группы")
	const query = `DELETE FROM group_subject WHERE id = ? RETURNING id`

	db.Init()

	var deletedId int
	err := db.DB.QueryRow(query, subjectId).Scan(&deletedId)
	if err == sql.ErrNoRows {
		log.Println("Предмет не найден или уже удален")
		return false, nil
	}
	if err != nil {
		log.Printf("Ошибка при удалении предмета у группы: %v", err)
		return false, err
	}

	log.Println("Успешное удаление предмета")
	return true, nil
}

// DeleteDisciplinesByGroupId удаляет все предметы для группы.
func DeleteDisciplinesByGroupId(groupID int) (bool, error) {
	log.Println("Удаление всех предметов группы")
	const query = `DELETE FROM group_subject WHERE group_id = ?`

	db.Init()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := db.DB.ExecContext(ctx, query, groupID)
	if err != nil {
		log.Printf("Ошибка при удалении предметов группы: %v", err)
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при проверке затронутых строк: %v", err)
		return false, err
	}

	log.Printf("Удалено предметов: %d", rowsAffected)
	return rowsAffected > 0, nil
}

// CopyDisciplinesBetweenGroups копирует предметы из одной группы в другую.
func CopyDisciplinesBetweenGroups(oldGroupId int, newGroupId int) error {
	log.Println("Копирование предметов между группами")
	const query = `
		WITH copied AS (
			SELECT ? AS new_group_id, subject_name, semester
			FROM group_subject
			WHERE group_id = ?
		)
		INSERT INTO group_subject (group_id, subject_name, semester)
		SELECT new_group_id, subject_name, semester
		FROM copied`

	db.Init()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.DB.ExecContext(ctx, query, newGroupId, oldGroupId)
	if err != nil {
		log.Printf("Ошибка при копировании предметов: %v", err)
		return err
	}

	log.Println("Успешное копирование предметов")
	return nil
}
