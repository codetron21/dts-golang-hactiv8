package model

type Item struct {
	ID          int    `gorm:"primaryKey;column:item_id;autoIncrement" json:"itemId,omitempty"`
	ItemCode    string `gorm:"not null;column:item_code" json:"itemCode"`
	Description string `gorm:"column:description" json:"description"`
	Quantity    int    `gorm:"column:quantity" json:"quantity"`
	OrderID     int    `gorm:"column:order_Id" json:"orderId"`
}
