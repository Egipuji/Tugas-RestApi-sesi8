package database

import (
	"tugas-restapi-sesi8/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=orders_by port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(models.Items{}, models.Orders{})
}

func GetDb() *gorm.DB {
	return db
}
