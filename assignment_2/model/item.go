package model

type Item struct {
	ItemId      uint   `gorm:"primaryKey;column:item_id"`
	ItemCode    string `gorm:"not null;unique;column:item_code"`
	Description string `gorm:"column:description"`
	Quantity    int    `gorm:"column:quantity"`
	OrderId     uint   `gorm:"column:order_id"`
}
