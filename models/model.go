package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Address struct {
	IDClient int    `gorm:"unique" json:"-"` // foreign key
	State    string `json:"state"`
	City     string `json:"city"`
	Gli      int    `json:"gli"`
}

type CallingTasks struct {
	ClientID  int     `json:"-"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Details   Address `gorm:"foreignKey:ClientID;references:IDClient" json:"details"`
	Email     string  `json:"email"`
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
