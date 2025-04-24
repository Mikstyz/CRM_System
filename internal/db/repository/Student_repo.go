package repository

import (
	db "CRM_System/internal/db"
	"CRM_System/internal/models"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func InfStdByGroup() ([]models.Student, error) {
	const query = `SELECT id, FullName, Well, Course, Speciality, GroupNum, Semester FROM students`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить студентов: %v", err)
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		err := rows.Scan(&s.ID, &s.FullName, &s.Groudates, &s.Course, &s.Speciality, &s.GroupNum, &s.Semester)
		if err != nil {
			return nil, fmt.Errorf("ошибка при чтении строки: %v", err)
		}
		students = append(students, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка после чтения строк: %v", err)
	}

	return students, nil
}

func GetStudentByID(studentID int) (*models.Student, error) {
	log.Println("Получение информации о студенте по ID")

	const query = `
		SELECT FullName, Speciality, GroupNum, Semester, Well, Course
		FROM students
		WHERE id = ?`

	db.Init()

	var student models.Student
	err := db.DB.QueryRow(query, studentID).Scan(
		&student.FullName,
		&student.Speciality,
		&student.GroupNum,
		&student.Semester,
		&student.Groudates,
		&student.Course,
	)
	if err != nil {
		log.Printf("Ошибка при получении студента: %v", err)
		return nil, err
	}

	return &student, nil
}

// Создание студента
func CrtStd(fullName string, Course byte, Groudates byte, speciality string, groupNum int, semester byte) (int, error) {
	log.Println("Создаю нового студента...")

	const query = `
		INSERT INTO students (FullName, Well, Course, Speciality, GroupNum, Semester)
		VALUES (?, ?, ?, ?, ?, ?)`

	res, err := db.DB.Exec(query, fullName, Groudates, Course, speciality, groupNum, semester)
	if err != nil {
		return 0, fmt.Errorf("не удалось вставить студента: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("не получил id нового студента: %v", err)
	}

	log.Printf("Студент с ID=%d успешно создан\n", id)
	return int(id), nil
}

func UpdateStd(studId int, newFullName string, newCourse byte, newGroudates byte, newSpeciality string, newGroupNum int, newSemester byte) (bool, error) {
	log.Printf("Обновляю данные студента с ID=%d...\n", studId)

	const query = `
		UPDATE students 
		SET FullName = ?, Well = ?, Course = ?, Speciality = ?, GroupNum = ?, Semester = ?
		WHERE id = ?`

	res, err := db.DB.Exec(query, newFullName, newGroudates, newCourse, newSpeciality, newGroupNum, newSemester, studId)
	if err != nil {
		return false, fmt.Errorf("не удалось обновить данные студента: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("не удалось получить количество обновленных строк: %v", err)
	}

	if rowsAffected > 0 {
		log.Printf("Данные студента с ID=%d успешно обновлены\n", studId)
		return true, nil
	}

	log.Printf("Студент с ID=%d не был обновлён (строки не были изменены)\n", studId)
	return false, nil
}

func DelStd(studId int) (bool, error) {
	log.Printf("Удаляю студента с ID=%d...\n", studId)

	const query = `DELETE FROM students WHERE id = ?`

	res, err := db.DB.Exec(query, studId)
	if err != nil {
		return false, fmt.Errorf("не удалось удалить студента: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("не удалось получить количество удаленных строк: %v", err)
	}

	if rowsAffected > 0 {
		log.Printf("Студент с ID=%d успешно удалён\n", studId)
		return true, nil
	}

	log.Printf("Студент с ID=%d не был удалён (строки не были изменены)\n", studId)
	return false, nil
}

// Продвинутое создание студента
func CreateStudentWithEmptyEmployment(fullName string, Course byte, Groudates byte, speciality string, groupNum int, semester byte) (int, error) {
	log.Println("Создаю нового студента и пустую запись о трудоустройстве (в транзакции)...")

	tx, err := db.DB.Begin()
	if err != nil {
		return 0, fmt.Errorf("не удалось начать транзакцию: %v", err)
	}

	// Шаг 1: вставляем студента
	const insertStudent = `
		INSERT INTO students (FullName, Well, Course, Speciality, GroupNum, Semester)
		VALUES (?, ?, ?, ?, ?, ?)`

	res, err := tx.Exec(insertStudent, fullName, Groudates, Course, speciality, groupNum, semester)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("не удалось вставить студента: %v", err)
	}

	studentID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("не получил id нового студента: %v", err)
	}

	// Шаг 2: вставляем пустую запись работодателя
	const insertEmployer = `
		INSERT INTO employers (studid, enterprise, workstartdate, jobtitle)
		VALUES (?, '', '', '')`

	_, err = tx.Exec(insertEmployer, studentID)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("не удалось вставить запись о трудоустройстве: %v", err)
	}

	// Всё ок — коммитим
	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("не удалось завершить транзакцию: %v", err)
	}

	log.Printf("✅ Студент с ID=%d и пустое трудоустройство успешно созданы\n", studentID)
	return int(studentID), nil
}
