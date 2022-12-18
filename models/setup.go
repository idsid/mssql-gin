package models

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnnectDatabase() {
	dsn := "sqlserver://Amirudev:amiru@DESKTOP-EN0V4OG?database=inventory"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}

	DB = db
}
