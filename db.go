package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbName   = "baby_list"
)

func SuccessMessage() {
	fmt.Println("Successfully connected !")
}

func InitDB() (*sql.DB, error) {

	postgresInfo := fmt.Sprintf(
		"host=%s port=%d user=%s " + "dbname=%s sslmode=disable",
		host, port, user, dbName)

	db, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		return nil, fmt.Errorf("connection to %s failed", dbName)
	} else {
		SuccessMessage()
	}

	return db, nil

}

