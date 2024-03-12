package models

type Items struct {
	ItemID      uint `gorm:"primarykey"`
	ItemCode    string
	Description string
	Quantity    uint
	OrderID     uint
}
