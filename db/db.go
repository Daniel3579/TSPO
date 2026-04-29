package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/lib/pq"
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

func Insert[T Databaseable](req []T) ([]T, error) {
	if len(req) == 0 {
		return nil, fmt.Errorf("Ошибка вставки. Пустой Databaseable(")
	}

	tableName := req[0].TableName()

	columnsArr := req[0].Columns()
	columns := strings.Join(columnsArr, ", ")

	placeholdersArr := make([]string, len(columnsArr))
	for i := range placeholdersArr {
		placeholdersArr[i] = fmt.Sprintf("$%d", i+1)
	}
	placeholders := strings.Join(placeholdersArr, ", ")

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING *;",
		pq.QuoteIdentifier(tableName),
		columns,
		placeholders,
	)

	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("Ошибка при начале транзакции: %w", err)
	}
	defer tx.Rollback()

	for i := range req {
		insertableValues := req[i].InsertableValues()
		selectableValues := req[i].SelectableValues()

		err := tx.QueryRow(query, insertableValues...).Scan(selectableValues...)
		if err != nil {
			return nil, fmt.Errorf("Не удалось записать в бд: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("Ошибка при коммите транзакции: %w", err)
	}

	return req, nil
}

func Select[T Requestable[K], K Databaseable](req T) ([]K, error) {
	tableName := req.TableName()
	params, err := req.Params()
	if err != nil {
		return nil, fmt.Errorf("Ошибка запроса: %w", err)
	}

	query := fmt.Sprintf("Select * from %s %s;",
		pq.QuoteIdentifier(tableName),
		params,
	)

	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("Ошибка при начале транзакции: %w", err)
	}
	defer tx.Rollback()

	var res []K = []K{}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить задачи: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		row := req.GetDatabaseable()
		selectableValues := row.SelectableValues()

		if err := rows.Scan(selectableValues...); err != nil {
			return nil, fmt.Errorf("Не удалось просканировать строку: %w", err)
		}
		res = append(res, row)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("Ошибка при коммите транзакции: %w", err)
	}

	return res, nil
}
