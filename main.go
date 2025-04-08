package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"fmt"
)

// User model
type User struct {
	ID uint `gorm:"primaryKey"` // auto-increment
	Name string
	Email string `gorm:"unique"`    // 👈 this makes the column UNIQUE
}

func main() {

	db, err := gorm.Open(sqlite.Open("db/gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// auto migration - creates table if not exists
	db.AutoMigrate(&User{})

	// create or insert a value to user tables
	// db.Create(&User{Name: "Amal", Email: "amal@gmail.com"})
	// db.Create(&User{Name: "Anaswara", Email: "anaswara@gmail.com"})
	// db.Create(&User{Name: "test", Email: "test@gmail.com"})

	// Read or display one value from table
	var user User
	db.First(&user, 1) // ID = 1
	fmt.Println("User Found:", user)

	// Update
	db.Model(&User{}).Where("id = ?", 1).Update("Email", "update@gmail.com")
	fmt.Println("Updated Data Found:", user) // ID = 1 look on line 33

	// Delete
	res := db.Delete(&User{}, 3)
	println(res)
}