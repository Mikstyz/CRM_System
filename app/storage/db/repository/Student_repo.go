package repository

import (
	"database/sql"
	"fmt"
	"log"

	db "CRM_System/app/storage/db"
	"CRM_System/app/storage/models"

	_ "modernc.org/sqlite"
)

// ----------------------Информация о студенте----------------------
// InfStudentByID возвращает данные студента по его ID.
func InfStudentByID(studentID int) (models.Student, error) {
	log.Println("[db][Student] - Получение информации о студенте по ID")

	const query = `
        SELECT 
            s.id, s.FullName, s.GroupId,
            s.enterprise, s.workstartdate, s.jobtitle
        FROM students s
        JOIN einf_groups g ON s.GroupId = g.Id
        WHERE s.id = ?`

	db.Init()

	var (
		student       models.Student
		enterprise    sql.NullString
		workStartDate sql.NullString
		jobTitle      sql.NullString
	)

	err := db.DB.QueryRow(query, studentID).Scan(
		&student.ID,
		&student.FullName,
		&student.GroupId,
		&enterprise,
		&workStartDate,
		&jobTitle,
	)
	if err != nil {
		log.Printf("[db][Student] - Ошибка при получении студента: %v", err)
		return models.Student{}, err
	}

	student.Enterprise = nil
	if enterprise.Valid {
		student.Enterprise = &enterprise.String
	}
	student.WorkStartDate = nil
	if workStartDate.Valid {
		student.WorkStartDate = &workStartDate.String
	}
	student.JobTitle = nil
	if jobTitle.Valid {
		student.JobTitle = &jobTitle.String
	}

	return student, nil
}

// InfIdByNameByGroup возвращает ID студента по имени и параметрам группы через GroupId.
func InfIdByNameByGroup(StudentName string, Course byte, Speciality string, Groduates byte, Number int) (int, error) {
	const query = `
	SELECT s.Id
	FROM students s
	JOIN einf_groups g ON s.GroupId = g.Id
	WHERE s.FullName = ? AND g.Course = ? AND g.Speciality = ? AND g.Groudates = ? AND g.GroupNum = ?
	`

	db.Init()

	var studentId int
	err := db.DB.QueryRow(query, StudentName, Course, Speciality, Groduates, Number).Scan(&studentId)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("[db][Student] - студент не найден по указанным данным")
		}
		log.Printf("[db][Student] - Ошибка при получении ID студента: %v", err)
		return 0, err
	}

	return studentId, nil
}

// ----------------------Информация о студентах----------------------
// InfStudentByGroup возвращает студентов по ID группы.
func InfStudentByGroup(GroupId int) ([]models.Student, error) {
	log.Printf("[db][Student] - Получение студентов по группе ID=%d", GroupId)

	const query = `
        SELECT 
            s.id, s.FullName, s.GroupId,
            s.enterprise, s.workstartdate, s.jobtitle
        FROM students s
        JOIN einf_groups g ON s.GroupId = g.Id
        WHERE g.Id = ?`

	db.Init()

	rows, err := db.DB.Query(query, GroupId)
	if err != nil {
		return []models.Student{}, fmt.Errorf("[db][Student] - не удалось получить студентов: %v", err)
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var (
			s             models.Student
			enterprise    sql.NullString
			workStartDate sql.NullString
			jobTitle      sql.NullString
		)

		err := rows.Scan(
			&s.ID,
			&s.FullName,
			&s.GroupId,
			&enterprise,
			&workStartDate,
			&jobTitle,
		)
		if err != nil {
			return []models.Student{}, fmt.Errorf("[db][Student] - ошибка при чтении строки: %v", err)
		}

		s.Enterprise = nil
		if enterprise.Valid {
			s.Enterprise = &enterprise.String
		}
		s.WorkStartDate = nil
		if workStartDate.Valid {
			s.WorkStartDate = &workStartDate.String
		}
		s.JobTitle = nil
		if jobTitle.Valid {
			s.JobTitle = &jobTitle.String
		}

		students = append(students, s)
	}

	if err := rows.Err(); err != nil {
		return []models.Student{}, fmt.Errorf("[db][Student] - ошибка после чтения строк: %v", err)
	}

	return students, nil
}

// InfStdByGroup возвращает список всех студентов с их GroupId.
func InfStdByGroup() ([]models.Student, error) {
	const query = `
        SELECT 
            s.id, s.FullName, s.GroupId,
            s.enterprise, s.workstartdate, s.jobtitle
        FROM students s
        JOIN einf_groups g ON s.GroupId = g.Id`

	db.Init()

	rows, err := db.DB.Query(query)
	if err != nil {
		return []models.Student{}, fmt.Errorf("[db][Student] - не удалось получить студентов: %v", err)
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var (
			s             models.Student
			enterprise    sql.NullString
			workStartDate sql.NullString
			jobTitle      sql.NullString
		)

		err := rows.Scan(
			&s.ID,
			&s.FullName,
			&s.GroupId,
			&enterprise,
			&workStartDate,
			&jobTitle,
		)
		if err != nil {
			return []models.Student{}, fmt.Errorf("[db][Student] - ошибка при чтении строки: %v", err)
		}

		s.Enterprise = nil
		if enterprise.Valid {
			s.Enterprise = &enterprise.String
		}
		s.WorkStartDate = nil
		if workStartDate.Valid {
			s.WorkStartDate = &workStartDate.String
		}
		s.JobTitle = nil
		if jobTitle.Valid {
			s.JobTitle = &jobTitle.String
		}

		students = append(students, s)
	}

	if err := rows.Err(); err != nil {
		return []models.Student{}, fmt.Errorf("[db][Student] - ошибка после чтения строк: %v", err)
	}

	return students, nil
}

// ----------------------Manager--------------------------
// CrtStd создаёт нового студента с пустыми полями трудоустройства.
func CrtStd(fullName string, groupId int, enterprise string, workstartdate string, jobtitle string) (int, error) {
	log.Println("[db][Student] - Создаю нового студента...")

	const insertStudentQuery = `
		INSERT INTO students (FullName, GroupId, enterprise, workstartdate, jobtitle)
		VALUES (?, ?, ?, ?, ?)
		RETURNING id`

	db.Init()

	var studentId int
	err := db.DB.QueryRow(insertStudentQuery, fullName, groupId, enterprise, workstartdate, jobtitle).Scan(&studentId)
	if err != nil {
		log.Printf("[db][Student] - Ошибка при создании студента: %v", err)
		return 0, fmt.Errorf("[db][Student] - не удалось вставить студента: %v", err)
	}

	log.Printf("[db][Student] - Студент с ID=%d успешно создан", studentId)
	return studentId, nil
}

// UpdateStd обновляет данные студента, включая поля трудоустройства.
func UpdateStd(studId int, newFullName string, newGroupId int, newEnterprise string, newWorkStartDate string, newJobTitle string) (bool, error) {
	log.Printf("[db][Student] - Обновляю данные студента с ID=%d...", studId)

	const updateQuery = `
		UPDATE students 
		SET FullName = ?, GroupId = ?, enterprise = ?, workstartdate = ?, jobtitle = ?
		WHERE id = ?
		RETURNING id`

	db.Init()

	var updatedId int
	err := db.DB.QueryRow(updateQuery, newFullName, newGroupId, newEnterprise, newWorkStartDate, newJobTitle, studId).Scan(&updatedId)
	if err == sql.ErrNoRows {
		log.Printf("[db][Student] - Студент с ID=%d не найден", studId)
		return false, nil
	}
	if err != nil {
		log.Printf("[db][Student] - Ошибка при обновлении студента: %v", err)
		return false, fmt.Errorf("не удалось обновить данные студента: %v", err)
	}

	log.Printf("[db][Student] - Данные студента с ID=%d успешно обновлены", studId)
	return true, nil
}

// DelStd удаляет студента по ID.
func DelStd(studId int) (bool, error) {
	log.Printf("[db][Student] - Удаляю студента с ID=%d...", studId)

	const query = `DELETE FROM students WHERE id = ? RETURNING id`

	db.Init()

	var deletedId int
	err := db.DB.QueryRow(query, studId).Scan(&deletedId)
	if err == sql.ErrNoRows {
		log.Printf("[db][Student] - Студент с ID=%d не найден", studId)
		return false, nil
	}
	if err != nil {
		log.Printf("[db][Student] - Ошибка при удалении студента: %v", err)
		return false, fmt.Errorf("[db][Student] - не удалось удалить студента: %v", err)
	}

	log.Printf("[db][Student] - Студент с ID=%d успешно удалён", studId)
	return true, nil
}
