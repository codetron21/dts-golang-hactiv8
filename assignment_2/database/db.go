package database

import (
	"assignment_2/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func Start() (Database, error) {
	const host = "localhost"
	const port = 5432
	const username = "postgres"
	const password = ""
	const dbName = "hacktiv-golang"

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbName)

	db, err := gorm.Open(postgres.Open(conn))
	if err != nil {
		fmt.Println("Error open connection to db", err)
		return Database{}, err
	}

	err = db.Debug().AutoMigrate(model.Order{}, model.Item{})

	if err != nil {
		fmt.Println("error on migration", err)
		return Database{}, err
	}

	return Database{
		db: db,
	}, nil
}
