package models

import "time"

type Orders struct {
	OrderID      uint `gorm:"primarykey"`
	CustomerName string
	OrderedAt    time.Time
	Items        []Items `gorm:"foreignKey:OrderID"`
}
