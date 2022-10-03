package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CallingTasks struct {
	ID        uint   `gorm:"primary key:autoIncrement" json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Address   string `json:"address"`
	Email     string `json:"email"`
}

var DB *gorm.DB

func ConnecttoDatabase() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=postgres dbname=calling password=Gandhi@123 sslmode=disable",
	}))
	if err != nil {
		panic("Error:Failed to connect to database!")
	}

	db.AutoMigrate(&CallingTasks{})

	DB = db
}
