package repository

import (
	db "CRM_System/app/storage/db"
	"CRM_System/app/storage/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

// --------------------------Информаия о всех группах--------------------------
// FetchAllGrp возвращает список всех групп.
func FetchAllGrp() ([]models.EinfGroup, error) {
	const query = `
		SELECT Id, Course, Speciality, Groudates, GroupNum
		FROM einf_groups
		ORDER BY Course, Speciality, GroupNum`

	db.Init()

	rows, err := db.DB.Query(query)
	if err != nil {
		log.Printf("[db][Group] - Ошибка при получении групп: %v", err)
		return nil, err
	}
	defer rows.Close()

	var groups []models.EinfGroup
	for rows.Next() {
		var group models.EinfGroup
		if err := rows.Scan(
			&group.Id,
			&group.Course,
			&group.Speciality,
			&group.Groudates,
			&group.Number,
		); err != nil {
			log.Printf("[db][Group] - Ошибка при сканировании группы: %v", err)
			continue
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[db][Group] - Ошибка при чтении групп: %v", err)
		return nil, err
	}

	return groups, nil
}

// FetchAllGrpWithSubjects возвращает список всех групп и их предметов.
func InfAllGrpWithSubjects() ([]models.InFGroupAndSubject, error) {
	const query = `
		SELECT 
			g.Id, g.Course, g.Speciality, g.Groudates, g.GroupNum,
			gs.id, gs.subject_name, gs.semester
		FROM einf_groups g
		LEFT JOIN group_subject gs ON gs.group_id = g.Id
		ORDER BY g.Course, g.Speciality, g.GroupNum, gs.semester
	`

	db.Init()

	rows, err := db.DB.Query(query)
	if err != nil {
		log.Printf("[db][Group] - Ошибка при получении групп с предметами: %v", err)
		return nil, err
	}
	defer rows.Close()

	groupMap := make(map[int]*models.InFGroupAndSubject)

	for rows.Next() {
		var (
			groupID     int
			course      byte
			speciality  string
			groudates   byte
			groupNum    int
			subjectID   sql.NullInt64
			subjectName sql.NullString
			semester    sql.NullInt64
		)

		if err := rows.Scan(
			&groupID,
			&course,
			&speciality,
			&groudates,
			&groupNum,
			&subjectID,
			&subjectName,
			&semester,
		); err != nil {
			log.Printf("[db][Group] - Ошибка при сканировании строки: %v", err)
			continue
		}

		if _, ok := groupMap[groupID]; !ok {
			groupMap[groupID] = &models.InFGroupAndSubject{
				Id:         groupID,
				Course:     course,
				Spesiality: speciality,
				Groduates:  groudates,
				Number:     groupNum,
				Subject: models.DisciplinesBySemester{
					OneSemester: []models.SemesterDiscipline{},
					TwoSemester: []models.SemesterDiscipline{},
				},
			}
		}

		if subjectID.Valid && subjectName.Valid && semester.Valid {
			discipline := models.SemesterDiscipline{
				Id:    int(subjectID.Int64),
				Title: subjectName.String,
			}

			if semester.Int64 == 1 {
				groupMap[groupID].Subject.OneSemester = append(groupMap[groupID].Subject.OneSemester, discipline)
			} else if semester.Int64 == 2 {
				groupMap[groupID].Subject.TwoSemester = append(groupMap[groupID].Subject.TwoSemester, discipline)
			}
		}
	}

	if err := rows.Err(); err != nil {
		log.Printf("[db][Group] - Ошибка при чтении строк: %v", err)
		return nil, err
	}

	var result []models.InFGroupAndSubject
	for _, grp := range groupMap {
		result = append(result, *grp)
	}

	return result, nil
}

// --------------------------Информация про группу--------------------------
// FetchGroupById возвращает данные группы по её ID.
func FetchGroupById(GroupId int) (models.EinfGroup, error) {
	const query = `SELECT Id, Course, Speciality, Groudates, GroupNum FROM einf_groups WHERE Id = ?`

	db.Init()

	var Group models.EinfGroup
	err := db.DB.QueryRow(query, GroupId).Scan(
		&Group.Id,
		&Group.Course,
		&Group.Speciality,
		&Group.Groudates,
		&Group.Number,
	)
	if err != nil {
		log.Printf("[db][Group] - Ошибка при получении данных группы: %v", err)
		return models.EinfGroup{}, err
	}

	return Group, nil
}

// FetchGrpWithSubjectsById возвращает инфу о группе и ее предметы.
func FetchGrpWithSubjectsById(groupId int) (models.InFGroupAndSubject, error) {
	const query = `
        SELECT 
            g.Id, g.Course, g.Speciality, g.Groudates, g.GroupNum,
            gs.id, gs.subject_name, gs.semester
        FROM einf_groups g
        LEFT JOIN group_subject gs ON gs.group_id = g.Id
        WHERE g.Id = ?
        ORDER BY g.Course, g.Speciality, g.GroupNum, gs.semester
    `

	db.Init()

	rows, err := db.DB.Query(query, groupId)
	if err != nil {
		log.Printf("[db][Group] - Ошибка при получении группы с предметами: %v", err)
		return models.InFGroupAndSubject{}, err
	}
	defer rows.Close()

	var result models.InFGroupAndSubject
	initialized := false

	for rows.Next() {
		var (
			groupID     int
			course      byte
			speciality  string
			groudates   byte
			groupNum    int
			subjectID   sql.NullInt64
			subjectName sql.NullString
			semester    sql.NullInt64
		)

		if err := rows.Scan(
			&groupID,
			&course,
			&speciality,
			&groudates,
			&groupNum,
			&subjectID,
			&subjectName,
			&semester,
		); err != nil {
			log.Printf("[db][Group] - Ошибка при сканировании строки: %v", err)
			continue
		}

		if !initialized {
			result = models.InFGroupAndSubject{
				Id:         groupID,
				Course:     course,
				Spesiality: speciality,
				Groduates:  groudates,
				Number:     groupNum,
				Subject: models.DisciplinesBySemester{
					OneSemester: []models.SemesterDiscipline{},
					TwoSemester: []models.SemesterDiscipline{},
				},
			}
			initialized = true
		}

		if subjectID.Valid && subjectName.Valid && semester.Valid {
			discipline := models.SemesterDiscipline{
				Id:    int(subjectID.Int64),
				Title: subjectName.String,
			}

			if semester.Int64 == 1 {
				result.Subject.OneSemester = append(result.Subject.OneSemester, discipline)
			} else if semester.Int64 == 2 {
				result.Subject.TwoSemester = append(result.Subject.TwoSemester, discipline)
			}
		}
	}

	if err := rows.Err(); err != nil {
		log.Printf("[db][Group] - Ошибка при чтении строк: %v", err)
		return models.InFGroupAndSubject{}, err
	}

	if !initialized {
		return models.InFGroupAndSubject{}, fmt.Errorf("[db][Group] - группа с ID %d не найдена", groupId)
	}

	return result, nil
}

// --------------------------Инструменты--------------------------
// MaxNumberByParams возвращает максимальный Id группы.
func MaxNumberByParams(course byte, groudates byte, speciality string) (int, error) {
	const query = `
		SELECT COALESCE(MAX(GroupNum), 0)
		FROM einf_groups
		WHERE Course = ? AND Groudates = ? AND Speciality = ?
	`

	db.Init()

	var maxGroupNum int
	err := db.DB.QueryRow(query, course, groudates, speciality).Scan(&maxGroupNum)
	if err != nil {
		return 0, fmt.Errorf("[db][Group] - не получилось получить максимальный GroupNum: %w", err)
	}

	return maxGroupNum, nil
}

// GetGroupIDByParams возвращает ID группы по параметрам.
func GetGroupIDByParams(course byte, groudates byte, speciality string, groupNum int) (int, error) {
	const query = `
		SELECT Id FROM einf_groups
		WHERE Speciality = ? AND GroupNum = ? AND Course = ? AND Groudates = ?`

	db.Init()

	var groupID int
	err := db.DB.QueryRow(query, speciality, groupNum, course, groudates).Scan(&groupID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("[db][Group] - группа с указанными параметрами не найдена")
		}
		log.Printf("[db][Group] - Ошибка при получении ID группы: %v", err)
		return 0, fmt.Errorf("[db][Group] - ошибка при получении ID группы: %w", err)
	}

	return groupID, nil
}

// ----------------------Manager--------------------------
// CrtGrp создаёт новую группу.
func CrtGrp(course byte, groudates byte, speciality string, groupNum int) (int, error) {
	const query = `
		INSERT INTO einf_groups (Course, Groudates, Speciality, GroupNum)
		VALUES (?, ?, ?, ?)
		RETURNING Id`

	db.Init()

	var groupId int
	err := db.DB.QueryRow(query, course, groudates, speciality, groupNum).Scan(&groupId)
	if err != nil {
		log.Printf("[db][Group] - Ошибка при создании группы: %v", err)
		return 0, err
	}

	return groupId, nil
}

// UpdateGrp обновляет данные группы.
func UpdateGrp(groupId int, newCourse byte, newGroudates byte, newSpeciality string, newGroupNum int) (bool, error) {
	const query = `
		UPDATE einf_groups
		SET Course = ?, Groudates = ?, Speciality = ?, GroupNum = ?
		WHERE Id = ?
		RETURNING Id`

	db.Init()

	var updatedId int
	err := db.DB.QueryRow(query, newCourse, newGroudates, newSpeciality, newGroupNum, groupId).Scan(&updatedId)
	if err == sql.ErrNoRows {
		log.Printf("Группа с ID=%d не найдена", groupId)
		return false, nil
	}
	if err != nil {
		log.Printf("[db][Group] - Ошибка при обновлении группы: %v", err)
		return false, err
	}

	return true, nil
}

// DelGrp удаляет группу по параметрам.
func DelGrp(GroupId int) (bool, error) {
	const query = `DELETE FROM einf_groups WHERE Id = ?`

	log.Printf("[db][Group] - Удаляем группу: GroupId=%v", GroupId)

	db.Init()

	res, err := db.DB.Exec(query, GroupId)
	if err != nil {
		log.Printf("Ошибка при удалении группы: %v", err)
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("[db][Group] - Ошибка при получении количества удалённых строк: %v", err)
		return false, err
	}

	if rowsAffected == 0 {
		log.Println("[db][Group] - Группа не найдена")
		return false, nil
	}

	return true, nil
}
