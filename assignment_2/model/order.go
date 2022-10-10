package model

import "time"

type Order struct {
	OrderId      uint      `gorm:"primaryKey;column:order_id"`
	CustomerName string    `gorm:"column:customer_name"`
	OrderedAt    time.Time `gorm:"column:ordered_at"`
	Items        []Item
}
