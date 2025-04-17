package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

const dbPath string = "E:/code/Go/Wails/CRM_System/Backend/Data/Sql/stud.db"

// DB — глобальная переменная для работы с базой данных
var DB *sql.DB

func Init() {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Fatalf("База данных не существует по пути: %s", dbPath)
	}

	// Открываем соединение с БД
	var err error
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("Ошибка при подключении к БД: %v", err)
	}

	// Проверяем, доступна ли база данных
	if err = DB.Ping(); err != nil {
		log.Fatalf("БД не отвечает: %v", err)
	}

}

func PingDB() error {
	if DB == nil {
		return fmt.Errorf("База данных не инициализирована")
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("не удается подключиться к базе данных: %v", err)
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
