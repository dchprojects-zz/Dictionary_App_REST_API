package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabase() (*sql.DB, error) {

	host := "50.87.139.17"
	port := "3306"
	user := "dchproje_N0"
	password := "Future$Projects@2021"
	dbName := "dchproje_dbmdwords"

	db, err := sql.Open("mysql", user+":"+password+"@tcp"+"("+host+":"+port+")"+"/"+dbName)
	if err != nil {
		return nil, err
	}
	
	return db, nil
}
