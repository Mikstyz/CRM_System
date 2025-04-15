package repository

import (
	db "CRM_System/Backend/internal/db"
	"CRM_System/Backend/internal/modeles"
	"fmt"

	_ "modernc.org/sqlite"
)

func InfEmpById(StudId int) (modeles.EInfEmp, error) {
	const query = "SELECT Enterprise, WorkStartDate, JobTitle FROM employers WHERE studId = ?"

	db.Init()

	// Выполняем запрос
	rows, err := db.DB.Query(query, StudId)
	if err != nil {
		return modeles.EInfEmp{}, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}
	defer rows.Close()

	var employer modeles.EInfEmp

	// Проверяем, есть ли данные
	if rows.Next() {
		// Если данные есть, сканируем их в структуру
		if err := rows.Scan(&employer.Enterprise, &employer.WorkStartDate, &employer.JobTitle); err != nil {
			return modeles.EInfEmp{}, fmt.Errorf("ошибка при сканировании данных: %w", err)
		}
	} else {
		// Если данных нет, возвращаем ошибку "нет данных"
		return modeles.EInfEmp{}, fmt.Errorf("работодатель с таким ID не найден")
	}

	// Проверка на ошибки после выполнения запроса
	if err := rows.Err(); err != nil {
		return modeles.EInfEmp{}, fmt.Errorf("ошибка при обработке результата запроса: %w", err)
	}

	// Возвращаем найденную информацию о работодателе
	return employer, nil
}

func CreateEmp(studId int, enterprise string, workStartDate string, jobTitle string) (bool, error) {
	const query = `INSERT INTO employers (studid, enterprise, workstartdate, jobtitle)
		VALUES (?, ?, ?, ?)`

	db.Init()

	res, err := db.DB.Exec(query, studId, enterprise, workStartDate, jobTitle)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func UpdateEmp(studId int, newEnterprise string, newWorkStartDate string, newJobTitle string) (bool, error) {
	const query = `UPDATE employers 
		SET enterprise = ?, workstartdate = ?, jobtitle = ? 
		WHERE studid = ?`

	db.Init()

	res, err := db.DB.Exec(query, newEnterprise, newWorkStartDate, newJobTitle, studId)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func DelEmp(studId int) (bool, error) {
	const query = `DELETE FROM employers 
		WHERE studid = ?`

	db.Init()

	res, err := db.DB.Exec(query, studId)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
