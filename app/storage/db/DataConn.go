package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

// DB — глобальная переменная для работы с базой данных
var DB *sql.DB

//const dbPath string = "E:\\code\\GIthub\\Education\\CRM_System\\app\\Data\\Sql\\stud.db"

func Init() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("[db][conn] - Не удалось получить путь к exe: %v", err)
	}

	exeDir := filepath.Dir(exePath) // Вот это папка с exe

	dbPath := filepath.Join(exeDir, "Data", "Sql", "stud.db")

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Fatalf("[db][conn] - База данных не существует по пути: %s", dbPath)
	}

	var errOpen error
	DB, errOpen = sql.Open("sqlite", dbPath)
	if errOpen != nil {
		log.Fatalf("[db][conn] - Ошибка при подключении к БД: %v", errOpen)
	}

	if errPing := DB.Ping(); errPing != nil {
		log.Fatalf("[db][conn] - БД не отвечает: %v", errPing)
	}
}

func PingDB() error {
	if DB == nil {
		return fmt.Errorf("[db][conn] - База данных не инициализирована")
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("[db][conn] - не удается подключиться к базе данных: %v", err)
	}

	return nil
}

// func Createdb() {
// 	const query string = ""

// 	_, err := DB.Exec(query)

// 	if err != nil {
// 		log.Fatal("создание таблиц", err)
// 	}

// 	log.Println("Таблицы созданы")
// }
