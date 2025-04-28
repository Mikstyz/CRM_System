package repository

import (
	"CRM_System/internal/db"
	"CRM_System/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

// InfAllGrp возвращает список всех групп.
func InfAllGrp() ([]models.EinfGroup, error) {
	const query = `
		SELECT Id, Course, Speciality, Groudates, GroupNum
		FROM einf_groups
		ORDER BY Course, Speciality, GroupNum`

	db.Init()

	rows, err := db.DB.Query(query)
	if err != nil {
		log.Printf("Ошибка при получении групп: %v", err)
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
			log.Printf("Ошибка при сканировании группы: %v", err)
			continue
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Ошибка при чтении групп: %v", err)
		return nil, err
	}

	return groups, nil
}

// InfGroupById возвращает данные группы по её ID.
func InfGroupById(GroupId int) (*models.EinfGroup, error) {
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
		log.Printf("Ошибка при получении данных группы: %v", err)
		return nil, err
	}

	return &Group, nil
}

func MaxNumberByParams(course byte, groudates byte, speciality string) (int, error) {
	const query = `
		SELECT COALESCE(MAX(GroupNum), 0)
		FROM groups
		WHERE Course = ? AND GroupDates = ? AND Speciality = ?
	`

	db.Init()

	var maxGroupNum int
	err := db.DB.QueryRow(query, course, groudates, speciality).Scan(&maxGroupNum)
	if err != nil {
		return 0, fmt.Errorf("не смогла получить максимальный GroupNum: %w", err)
	}

	return maxGroupNum, nil
}

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
		log.Printf("Ошибка при создании группы: %v", err)
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
		log.Printf("Ошибка при обновлении группы: %v", err)
		return false, err
	}

	return true, nil
}

// DelGrp удаляет группу по параметрам.
func DelGrp(course byte, graduates byte, speciality string, groupNum int) (bool, error) {
	const query = `DELETE FROM einf_groups WHERE Course = ? AND Groudates = ? AND Speciality = ? AND GroupNum = ? RETURNING Id`

	log.Printf("Удаляем группу: Course=%v, Groudates=%v, Speciality=%s, GroupNum=%v", course, graduates, speciality, groupNum)

	db.Init()

	var deletedId int
	err := db.DB.QueryRow(query, course, graduates, speciality, groupNum).Scan(&deletedId)
	if err == sql.ErrNoRows {
		log.Println("Группа не найдена")
		return false, nil
	}
	if err != nil {
		log.Printf("Ошибка при удалении группы: %v", err)
		return false, err
	}

	return true, nil
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
			return 0, fmt.Errorf("группа с указанными параметрами не найдена")
		}
		log.Printf("Ошибка при получении ID группы: %v", err)
		return 0, fmt.Errorf("ошибка при получении ID группы: %w", err)
	}

	return groupID, nil
}
