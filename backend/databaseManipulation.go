package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Title     string
	Content   string
	TestEntry int
}

func retrieveArticle() {
	var articles []Article

	db, err := gorm.Open(sqlite.Open("skjsports.db"), &gorm.Config{})
	// db, err := gorm.Open("sqlite3", "./skjsports.db")

	if err != nil {
		panic("failed to connect database")
	}

	db.Find(&articles)

	fmt.Println(len(articles))

	fmt.Println(articles[0])
}
