package repository

import (
	"CRM_System/Backend/internal/db"
	"CRM_System/Backend/internal/modeles"
	"database/sql"
	"errors"
	"fmt"
)

func InfAllGrp() ([]modeles.EinfGroup, error) {
	const query = `
		SELECT Id, Well, Speciality, GClass, Semester, GroupNum
		FROM einf_groups
		ORDER BY GClass, Speciality, GroupNum`

	db.Init()

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []modeles.EinfGroup

	for rows.Next() {
		var group modeles.EinfGroup
		if err := rows.Scan(
			&group.Id,
			&group.Well,
			&group.Speciality,
			&group.GClass,
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

func CrtGrp(well byte, gClass byte, speciality string, groupNum int, semester byte) (int, error) {
	const query = `
		INSERT INTO einf_groups (Well, gClass, Speciality, GroupNum, Semester)
		VALUES (?, ?, ?, ?, ?)`

	db.Init()

	res, err := db.DB.Exec(query, well, gClass, speciality, groupNum, semester)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func UpdateGrp(groupId int, newWell byte, newGClass byte, newSpeciality string, newGroupNum int, newSemester byte) (bool, error) {
	const query = `
		UPDATE einf_groups
		SET Well = ?, GClass = ?, Speciality = ?, GroupNum = ?, Semester = ?
		WHERE Id = ?`

	db.Init()

	res, err := db.DB.Exec(query, newWell, newGClass, newSpeciality, newGroupNum, newSemester, groupId)
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

func GetGroupIDByParams(well byte, gClass byte, speciality string, groupNum int, semester byte) (int, error) {
	var groupID int

	err := db.DB.QueryRow(`
		SELECT id FROM einf_groups
		WHERE Speciality = $1 AND GroupNum = $2 AND Semester = $3 AND Well = $4 AND GClass = $5
	`, speciality, groupNum, semester, well, gClass).Scan(&groupID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("группа с указанными параметрами не найдена")
		}
		return 0, fmt.Errorf("ошибка при получении ID группы: %w", err)
	}

	return groupID, nil
}
