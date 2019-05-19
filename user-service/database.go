package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

const (
	DB_DEFAULT_HOST = "172.17.0.1"
	DB_DEFAULT_USER = "shippy"
	DB_DEFAULT_PASSWORD = "shippy"
	DB_DEFAULT_DBNAME = "shippy"
)

func CreateConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	if host == "" {
		host = DB_DEFAULT_HOST
	}

	if user == "" {
		user = DB_DEFAULT_USER
	}

	if password == "" {
		password = DB_DEFAULT_PASSWORD
	}

	if dbName == "" {
		dbName = DB_DEFAULT_DBNAME
	}

	return gorm.Open(
				"postgres",
				fmt.Sprintf(
					"host=%s user=%s dbname=%s sslmode=disable password=%s",
					host, user, dbName, password))
}

