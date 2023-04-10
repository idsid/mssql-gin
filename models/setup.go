package models

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnnectDatabase() {
	dsn := "sqlserver://b1admin:1dsSky.123@94.237.79.44?database=KASAN"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// err = db.AutoMigrate(&Product{})
	// if err != nil {
	// 	panic(err)
	// }

	DB = db
}
