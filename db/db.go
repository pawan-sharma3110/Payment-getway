package db

import (
	"database/sql"
	"fmt"
)

func DbIn() (db *sql.DB, err error) {

	connStr := "host=localhost port=5432 user=postgres dbname=Payament_GetWay sslmode=disable password=Pawan@2003"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("invalid connection string: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("falid to connect to database : %w", err)
	}
	return db, nil
}
