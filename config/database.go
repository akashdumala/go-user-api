package config

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConnectDB() error {

	connStr := "postgres://postgres:Akash4151@localhost:5432/usersdb"

	db, err := sql.Open("pgx", connStr)

	if err != nil {
		return err
	}

	err = db.Ping()

	if err != nil {
		return err
	}

	DB = db

	fmt.Println("Database Connected Successfully")

	return nil
}
