package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	dsn := "root:password@tcp(127.0.0.1:3306)/todo_db"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to the database!")
	}

	if err = DB.Ping(); err != nil {
		panic("Database is unreachable!")
	}

	fmt.Println("Database connected successfully!")
}
