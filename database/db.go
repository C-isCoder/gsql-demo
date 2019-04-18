package database

import (
	"database/sql"
	"fmt"
)

func GetDB() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "db_employee"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		fmt.Printf("open db error:%v\n", err.Error())
	}
	return db
}
