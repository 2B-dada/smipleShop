package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID     int    `gorm:"primaryKey"`
	Name   string `gorm:"column:product_name"`
	Price  int    `gorm:"column:product_price"`
	Number int    `gorm:"column:product_number"`
}

func (Product) TableName() string {
	return "Product"
}

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(
		mysql.Open("test:test@tcp(127.0.0.1:3306)/gormdb?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}
}
