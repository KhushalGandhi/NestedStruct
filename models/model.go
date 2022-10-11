package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Info struct {
	ID uint `gorm:"primary key:autoIncrement" json:"id"`
	//Firstname string  `json:"firstname"`
	//Lastname  string  `json:"lastname"`
	//Email     string  `json:"email"`
	ClientID int     `json:"-"`
	PersonID int     `json:"-"`
	Person   Person  `gorm:"foreignkey:PersonID;references:ID" json:"person"`
	Address  Address `gorm:"foreignKey:ClientID;references:ID" json:"address"`
}

type Address struct {
	ID     uint   `gorm:"unique" json:"-" postgressql:"type:uint REFERENCES address(id) ON DELETE CASCADE"` // foreign key
	State  string `json:"state"`
	City   string `json:"city"`
	Street string `json:"street"`
}

type Person struct {
	ID        uint   `gorm:"unique" json:"-" postgressql:"type:uint REFERENCES person(id) ON DELETE CASCADE"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
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

	db.AutoMigrate(&Info{})
	//db.Preload("Address").First(&CallingTasks)
	DB = db
}
