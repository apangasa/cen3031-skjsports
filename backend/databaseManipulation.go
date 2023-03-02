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

func retrieveArticle(article_id string) *Article {
	article := new(Article)

	db, err := gorm.Open(sqlite.Open("skjsports.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	result := db.First(&article, "id = ?", article_id)

	if result.Error == gorm.ErrRecordNotFound {
		return nil // if record not found
	}

	return article
}

func searchArticle(search string) []Article {
	var articles []Article

	if search == "" {
		return nil
	}

	for i := 0; i < len(search); i++ {
		if search[i] != ' ' {
			break
		}
		if i == len(search)-1 {
			return nil
		}
	}

	db, err := gorm.Open(sqlite.Open("skjsports.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	//db.First(&articles, "title = ?", search)
	db.Raw("SELECT id, title FROM articles WHERE title LIKE '%" + search + "%'").Scan(&articles)

	return articles

}

func addSubscriber(email string) {
	// TODO insert into Subscribers table
	// consider what attributes should be stored (e.g. first name, last name, email)
}
