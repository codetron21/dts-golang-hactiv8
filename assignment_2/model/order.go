package model

import "time"

type Order struct {
	OrderId      int       `gorm:"primaryKey;column:order_id;autoIncrement" json:"order_id,omitempty"`
	CustomerName string    `gorm:"column:customer_name" json:"customer_name"`
	OrderedAt    time.Time `gorm:"column:ordered_at" json:"order_at"`
	Items        []Item    `gorm:"constraint:OnDelete:CASCADE" json:"items"`
}
