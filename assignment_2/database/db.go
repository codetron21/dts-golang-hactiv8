package database

import (
	"assignment_2/model"
	"errors"
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
	data := model.Order{}
	err = d.GetOrderById(&data, orderId)

	return data, err
}

func (d Database) GetOrderById(order *model.Order, orderId int) error {
	order.ID = orderId

	log.Println("get order by id:", order)

	err := d.db.Model(&model.Order{}).Preload("Items").Find(&order).Error
	if err != nil {
		fmt.Println("error get data by id", err)
		return errors.New("error get order by id")
	}

	return nil
}

func (d Database) GetOrders(orders *[]model.Order) error {
	err := d.db.Model(&model.Order{}).Preload("Items").Find(orders).Error
	if err != nil {
		fmt.Println("error get orders", err)
		return errors.New("error get orders")
	}

	log.Println("orders", orders)

	return nil
}

func (d Database) UpdateOrderById(newOrder *model.Order) error {
	log.Println("print new order:", newOrder)

	err := d.db.First(&model.Order{}, newOrder.ID).Error
	if err != nil {
		return errors.New("order with id not found")
	}

	err = d.db.Model(&model.Order{}).Where("order_id = ?", newOrder.ID).Updates(&newOrder).Error
	if err != nil {
		log.Println("error update:", err)
		return errors.New("error update order data")
	}

	return nil
}

func (d Database) DeleteOrderById(orderId int) error {
	err := d.db.First(&model.Order{}, orderId).Error
	if err != nil {
		return errors.New("order with id not found")
	}

	err = d.db.Delete(&model.Order{}, orderId).Error
	if err != nil {
		log.Println("error delete:", err)
		return errors.New("error delete data order")
	}

	return nil
}
