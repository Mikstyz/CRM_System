package repository

import (
	"CRM_System/internal/db"
	"CRM_System/internal/models"
	"database/sql"
	"errors"
	"fmt"
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

func CrtGrp(course byte, groudates byte, speciality string, groupNum int, semester byte) (int, error) {
	const query = `
		INSERT INTO einf_groups (Course, Groudates, Speciality, GroupNum, Semester)
		VALUES (?, ?, ?, ?, ?)`

	db.Init()

	res, err := db.DB.Exec(query, course, groudates, speciality, groupNum, semester)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func UpdateGrp(groupId int, newCourse byte, newGroudates byte, newSpeciality string, newGroupNum int, newSemester byte) (bool, error) {
	const query = `
		UPDATE einf_groups
		SET Course = ?, Groudates = ?, Speciality = ?, GroupNum = ?, Semester = ?
		WHERE Id = ?`

	db.Init()

	res, err := db.DB.Exec(query, newCourse, newGroudates, newSpeciality, newGroupNum, newSemester, groupId)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func DelGrp(groupId int) (bool, error) {
	const query = `DELETE FROM einf_groups WHERE Id = ?`

	db.Init()

	res, err := db.DB.Exec(query, groupId)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
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
