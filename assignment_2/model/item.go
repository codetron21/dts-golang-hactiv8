package model

type Item struct {
	ItemId      int    `gorm:"primaryKey;column:item_id;autoIncrement" json:"item_id,omitempty"`
	ItemCode    string `gorm:"not null;column:item_code" json:"item_code"`
	Description string `gorm:"column:description" json:"description"`
	Quantity    int    `gorm:"column:quantity" json:"quantity"`
	OrderId     int    `gorm:"column:order_id" json:"order_id"`
}
