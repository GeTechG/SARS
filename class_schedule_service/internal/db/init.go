package db

import (
	"database/sql"
	"git.it-college.ru/i21s617/SARS/class_schedule_service/internal/sqlc"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var queries *sqlc.Queries
var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		return err
	}

	queries = sqlc.New(db)

	return nil
}

func GetQueries() *sqlc.Queries {
	return queries
}

func CloseDB() {
	_ = db.Close()
}
