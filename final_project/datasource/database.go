package datasource

import (
	"fmt"

	"final_project/config"
	"final_project/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	conn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_DATABASE_NAME,
	)

	db, err := gorm.Open(postgres.Open(conn))
	if err != nil {
		fmt.Println("Error open connection to db", err)
		return nil, err
	}

	err = db.Debug().AutoMigrate(
		model.User{}, model.Photo{}, model.Comment{}, model.SocialMedia{},
	)
	if err != nil {
		fmt.Println("error on migration", err)
		return nil, err
	}

	return db, nil
}
