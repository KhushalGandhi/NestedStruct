package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CallingTasks struct {
	ID        uint    `gorm:"primary key:autoIncrement" json:"id"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Email     string  `json:"email"`
	ClientID  int     `json:"-"`
	Address   Address `gorm:"foreignKey:ClientID;references:ID" json:"address"`
}

type Address struct {
	ID     uint   `gorm:"unique" json:"-"` // foreign key
	State  string `json:"state"`
	City   string `json:"city"`
	Street string `json:"street"`
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
	//db.Preload("Address").First(&CallingTasks)
	DB = db
}
