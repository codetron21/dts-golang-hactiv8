package database

import (
	"assignment_2/model"
	"fmt"
	"log"

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
	const password = "postgres"
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

/*
 * CRUD Orders
 */

func (d Database) CreateOrder(order model.Order) (model.Order, error) {
	err := d.db.Create(&order).Error

	if err != nil {
		return model.Order{}, err
	}

	orderId := order.ID

	createResult, err := d.GetOrderById(orderId)

	return createResult, err
}

func (d Database) GetOrderById(orderId int) (model.Order, error) {
	var order = model.Order{}
	order.ID = orderId
	err := d.db.Model(&model.Order{}).Preload("Items").Find(&order).Error
	log.Println("get order by id:", order)
	return order, err
}

func (d Database) GetOrders() ([]model.Order, error) {
	var orders []model.Order
	err := d.db.Model(&model.Order{}).Preload("Items").Find(&orders).Error
	log.Println("get orders:", orders)
	return orders, err
}

func (d Database) UpdateOrderById(orderId int, newOrder *model.Order) (model.Order, error) {
	newOrder.ID = orderId

	findResult := d.db.First(&model.Order{}, orderId)
	err := findResult.Error
	if err != nil {
		return model.Order{}, err
	}

	err = d.db.Save(&newOrder).Error

	log.Println("print new order (database):", newOrder)

	return *newOrder, err
}

func (d Database) DeleteOrderById(orderId int) (int, error) {
	findResult := d.db.First(&model.Order{}, orderId)
	err := findResult.Error
	if err != nil {
		return -1, err
	}

	result := d.db.Delete(&model.Order{}, orderId)

	log.Println("delete result:", result)

	return orderId, result.Error
}
