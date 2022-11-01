package model

import "time"

type Order struct {
	ID           int       `gorm:"primaryKey;column:order_id;autoIncrement" json:"orderId,omitempty"`
	CustomerName string    `gorm:"column:customer_name" json:"customerName"`
	OrderedAt    time.Time `gorm:"column:ordered_at" json:"orderAt"`
	Items        []Item    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
}
