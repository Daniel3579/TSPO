package db

import (
	"database/sql"
	"fmt"
	"os"
	"tspo/dtos"
	// _ "github.com/lib/pq"
)

var db *sql.DB

// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

func ConnectDB(env string) error {
	var connStr string = os.Getenv(env)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("Ошибка при открытии базы данных: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Не удалось пингануть бд: %w", err)
	}

	return nil
}

func CloseDB() {
	db.Close()
}

// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

func Insert(req *dtos.Insertable) {

}
