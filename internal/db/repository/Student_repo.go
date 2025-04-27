package repository

import (
	"database/sql"
	"fmt"
	"log"

	db "CRM_System/internal/db"
	"CRM_System/internal/models"

	_ "modernc.org/sqlite"
)

// InfStdByGroup возвращает список всех студентов.
func InfStdByGroup() ([]models.Student, error) {
	const query = `SELECT id, FullName, Groudates, Course, Speciality, GroupNum FROM students`

	db.Init()

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить студентов: %v", err)
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		err := rows.Scan(&s.ID, &s.FullName, &s.Groudates, &s.Course, &s.Speciality, &s.GroupNum)
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

// InfStudentByGroup возвращает студентов по ID группы.
func InfStudentByGroup(GroupId int) ([]models.Student, error) {
	log.Printf("Получение студентов по группе ID=%d", GroupId)

	const query = `
		SELECT s.id, s.FullName, s.Groudates, s.Course, s.Speciality, s.GroupNum
		FROM students s 
		JOIN einf_groups eg ON s.Course = eg.Course AND s.Speciality = eg.Speciality AND s.GroupNum = eg.GroupNum
		WHERE eg.Id = ?`

	db.Init()

	rows, err := db.DB.Query(query, GroupId)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить студентов: %v", err)
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		err := rows.Scan(&s.ID, &s.FullName, &s.Groudates, &s.Course, &s.Speciality, &s.GroupNum)
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

// GetIdByNameByGroup возвращает ID студента по имени и параметрам группы.
func GetIdByNameByGroup(StudentName string, Course byte, Speciality string, Groduates byte, Number int) (int, error) {
	const query = `SELECT Id FROM students WHERE FullName = ? AND Course = ? AND Speciality = ? AND Groudates = ? AND GroupNum = ?`

	db.Init()

	var studentId int
	err := db.DB.QueryRow(query, StudentName, Course, Speciality, Groduates, Number).Scan(&studentId)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("студент не найден по указанным данным")
		}
		log.Printf("Ошибка при получении ID студента: %v", err)
		return 0, err
	}

	return studentId, nil
}

// GetStudentByID возвращает данные студента по его ID.
func GetStudentByID(studentID int) (models.Student, error) {
	log.Println("Получение информации о студенте по ID")

	const query = `
		SELECT FullName, Speciality, GroupNum, Groudates, Course
		FROM students
		WHERE id = ?`

	db.Init()

	var student models.Student
	err := db.DB.QueryRow(query, studentID).Scan(
		&student.FullName,
		&student.Speciality,
		&student.GroupNum,
		&student.Groudates,
		&student.Course,
	)
	if err != nil {
		log.Printf("Ошибка при получении студента: %v", err)
		return models.Student{}, err
	}

	return student, nil
}

// GetStudentByGroup возвращает студентов по параметрам группы.
func GetStudentByGroup(course byte, speciality string, groupNum int) ([]models.Student, error) {
	const query = `
		SELECT id, FullName, Groudates, Course, Speciality, GroupNum
		FROM students 
		WHERE Course = ? AND Speciality = ? AND GroupNum = ?`

	db.Init()

	rows, err := db.DB.Query(query, course, speciality, groupNum)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить студентов: %v", err)
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		err := rows.Scan(&s.ID, &s.FullName, &s.Groudates, &s.Course, &s.Speciality, &s.GroupNum)
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

// CrtStd создаёт нового студента.
func CrtStd(fullName string, Course byte, Groudates byte, speciality string, groupNum int) (int, error) {
	log.Println("Создаю нового студента...")

	const query = `
		INSERT INTO students (FullName, Groudates, Course, Speciality, GroupNum)
		VALUES (?, ?, ?, ?, ?)
		RETURNING id`

	db.Init()

	var studentId int
	err := db.DB.QueryRow(query, fullName, Groudates, Course, speciality, groupNum).Scan(&studentId)
	if err != nil {
		log.Printf("Ошибка при создании студента: %v", err)
		return 0, fmt.Errorf("не удалось вставить студента: %v", err)
	}

	log.Printf("Студент с ID=%d успешно создан", studentId)
	return studentId, nil
}

// UpdateStd обновляет данные студента.
func UpdateStd(studId int, newFullName string, newCourse byte, newGroudates byte, newSpeciality string, newGroupNum int) (bool, error) {
	log.Printf("Обновляю данные студента с ID=%d...", studId)

	const query = `
		UPDATE students 
		SET FullName = ?, Groudates = ?, Course = ?, Speciality = ?, GroupNum = ?
		WHERE id = ?
		RETURNING id`

	db.Init()

	var updatedId int
	err := db.DB.QueryRow(query, newFullName, newGroudates, newCourse, newSpeciality, newGroupNum, studId).Scan(&updatedId)
	if err == sql.ErrNoRows {
		log.Printf("Студент с ID=%d не найден", studId)
		return false, nil
	}
	if err != nil {
		log.Printf("Ошибка при обновлении студента: %v", err)
		return false, fmt.Errorf("не удалось обновить данные студента: %v", err)
	}

	log.Printf("Данные студента с ID=%d успешно обновлены", studId)
	return true, nil
}

// DelStd удаляет студента по ID.
func DelStd(studId int) (bool, error) {
	log.Printf("Удаляю студента с ID=%d...", studId)

	const query = `DELETE FROM students WHERE id = ? RETURNING id`

	db.Init()

	var deletedId int
	err := db.DB.QueryRow(query, studId).Scan(&deletedId)
	if err == sql.ErrNoRows {
		log.Printf("Студент с ID=%d не найден", studId)
		return false, nil
	}
	if err != nil {
		log.Printf("Ошибка при удалении студента: %v", err)
		return false, fmt.Errorf("не удалось удалить студента: %v", err)
	}

	log.Printf("Студент с ID=%d успешно удалён", studId)
	return true, nil
}

// CreateStudentWithEmptyEmployment создаёт студента и пустую запись о трудоустройстве.
func CreateStudentWithEmptyEmployment(fullName string, Course byte, Groudates byte, speciality string, groupNum int) (int, error) {
	log.Println("Создаю нового студента и пустую запись о трудоустройстве (в транзакции)...")

	db.Init()

	tx, err := db.DB.Begin()
	if err != nil {
		log.Printf("Ошибка при начале транзакции: %v", err)
		return 0, fmt.Errorf("не удалось начать транзакцию: %v", err)
	}

	const insertStudent = `
		INSERT INTO students (FullName, Groudates, Course, Speciality, GroupNum)
		VALUES (?, ?, ?, ?, ?)
		RETURNING id`

	var studentID int
	err = tx.QueryRow(insertStudent, fullName, Groudates, Course, speciality, groupNum).Scan(&studentID)
	if err != nil {
		tx.Rollback()
		log.Printf("Ошибка при создании студента: %v", err)
		return 0, fmt.Errorf("не удалось вставить студента: %v", err)
	}

	const insertEmployer = `
		INSERT INTO employers (studid, enterprise, workstartdate, jobtitle)
		VALUES (?, NULL, NULL, NULL)`

	_, err = tx.Exec(insertEmployer, studentID)
	if err != nil {
		tx.Rollback()
		log.Printf("Ошибка при создании записи о трудоустройстве: %v", err)
		return 0, fmt.Errorf("не удалось вставить запись о трудоустройстве: %v", err)
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Ошибка при коммите транзакции: %v", err)
		return 0, fmt.Errorf("не удалось завершить транзакцию: %v", err)
	}

	log.Printf("Студент с ID=%d и пустое трудоустройство успешно созданы", studentID)
	return studentID, nil
}
