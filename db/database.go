package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabase() (*sql.DB, error) {

	user := "d9d9vs9"
	password := "Vagina$2020"
	dbName := "dbword"

	db, err := sql.Open("mysql", user+":"+password+"@/"+dbName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
