package repository

import (
	"CRM_System/internal/db"
	"CRM_System/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func InfAllGrp() ([]models.EinfGroup, error) {
	const query = `
		SELECT Id, Course, Speciality, Groudates, Semester, GroupNum
		FROM einf_groups
		ORDER BY Course, Speciality, GroupNum`

	db.Init()

	rows, err := db.DB.Query(query)
	if err != nil {
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
			&group.Semester,
			&group.Number,
		); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func InfGroupById(GroupId int) (*models.EinfGroup, error) {
	const query = `SELECT Id, Course, Speciality, Groudates, Semester, GroupNum FROM einf_groups WHERE id = ?`

	db.Init()

	var Group models.EinfGroup
	err := db.DB.QueryRow(query, GroupId).Scan(
		&Group.Course,
		&Group.Speciality,
		&Group.Groudates,
		&Group.Semester,
		&Group.Number,
	)
	if err != nil {
		log.Printf("Ошибка при получении данных группы: %v", err)
		return nil, err
	}

	return &Group, nil
}

func CrtGrp(course byte, groudates byte, speciality string, groupNum int) ([2]int, error) {
	const query = `
		INSERT INTO einf_groups (Course, Groudates, Speciality, GroupNum, Semester)
		VALUES (?, ?, ?, ?, ?)`

	db.Init()

	tx, err := db.DB.Begin()
	if err != nil {
		return [2]int{}, err
	}

	var ids [2]int

	// Первая запись с semester = 1
	res1, err := tx.Exec(query, course, groudates, speciality, groupNum, 1)
	if err != nil {
		tx.Rollback()
		return [2]int{}, err
	}

	id1, err := res1.LastInsertId()
	if err != nil {
		tx.Rollback()
		return [2]int{}, err
	}
	ids[0] = int(id1)

	// Вторая запись с semester = 2
	res2, err := tx.Exec(query, course, groudates, speciality, groupNum, 2)
	if err != nil {
		tx.Rollback()
		return [2]int{}, err
	}

	id2, err := res2.LastInsertId()
	if err != nil {
		tx.Rollback()
		return [2]int{}, err
	}
	ids[1] = int(id2)

	err = tx.Commit()
	if err != nil {
		return [2]int{}, err
	}

	return ids, nil
}

func UpdateGrp(groupId int, newCourse byte, newGroudates byte, newSpeciality string, newGroupNum int) (bool, error) {
	const query = `
		UPDATE einf_groups
		SET Course = ?, Groudates = ?, Speciality = ?, GroupNum = ?
		WHERE Id = ?`

	db.Init()

	res, err := db.DB.Exec(query, newCourse, newGroudates, newSpeciality, newGroupNum, groupId)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func DelGrp(course byte, graduates byte, speciality string, groupNum int) (bool, error) {
	const query = `DELETE FROM einf_groups WHERE Course = ? AND Groudates = ? AND Speciality = ? AND GroupNum = ?`

	log.Printf("Удаляем группу: Course=%v, Groudates=%v, Speciality=%s, GroupNum=%v", course, graduates, speciality, groupNum)
	res, err := db.DB.Exec(query, course, graduates, speciality, groupNum)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("RowsAffected error: %v", err)
		return false, err
	}

	return rowsAffected > 0, nil
}

func GetGroupIDByParams(course byte, groudates byte, speciality string, groupNum int, semester byte) (int, error) {
	var groupID int

	err := db.DB.QueryRow(`
		SELECT id FROM einf_groups
		WHERE Speciality = $1 AND GroupNum = $2 AND Semester = $3 AND Course = $4 AND Groudates = $5
	`, speciality, groupNum, semester, course, groudates).Scan(&groupID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("группа с указанными параметрами не найдена")
		}
		return 0, fmt.Errorf("ошибка при получении ID группы: %w", err)
	}

	return groupID, nil
}
