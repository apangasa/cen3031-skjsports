package main

import (
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

func retrieveArticle(article_id string) Article {
	var article Article

	db, err := gorm.Open(sqlite.Open("skjsports.db"), &gorm.Config{})
	// db, err := gorm.Open("sqlite3", "./skjsports.db")

	if err != nil {
		panic("failed to connect database")
	}

	db.First(&article, "id = ?", article_id)

	return article
}

func searchArticle(search string) []Article {
	var articles []Article

	db, err := gorm.Open(sqlite.Open("skjsports.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	//db.First(&articles, "title = ?", search)
	db.Raw("SELECT id, title FROM articles WHERE title LIKE '%" + search + "%'").Scan(&articles)

	return articles

}
