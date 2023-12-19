package datastore

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBUsername = "root"
	DBPassword = "Xonen@3616"
	DBHost     = "127.0.0.1"
	DBPort     = "3306"
	DBName     = "mysql_database"
)

func ConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName)
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", ConnectionString())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("Connected to the database")

	return db, nil
}
