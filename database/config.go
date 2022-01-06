package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Open() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/migratedb")
	if err != nil {
		fmt.Println("Error in opening DB", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error in pinging bd", err)
	}
	fmt.Println("Connected to DB ")
	return db
}

// File Path
const filePath string = "file://F:/projects/go/sql-migrate-cli/migrations"
