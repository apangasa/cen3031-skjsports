package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Title       string
	Content     string
	Author      string
	AuthorEmail string
	IsDraft     int
}

type Subscriber struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	FirstName string
	LastName  string
}

type Author struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Author      string
	AuthorEmail string
}

func retrieveArticle(article_id string) *Article {
	article := new(Article)

	db, err := gorm.Open(sqlite.Open("skjsports.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	result := db.Where("is_draft = ?", "0").First(&article, "id = ?", article_id)

	if result.Error == gorm.ErrRecordNotFound {
		return nil // if record not found
	}

	return article
}

func retrieveAuthorArticles(author_id string) []Article {
	var articles []Article
	author := new(Author)

	db, err := gorm.Open(sqlite.Open("skjsports.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	result := db.First(&author, "id = ?", author_id)

	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}

	db.Where("author_email = ? AND is_draft = ?", author.AuthorEmail, "0").Find(&articles)

	return articles
}

func searchDatabaseForArticles(search string) []Article {
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
	db.Raw("SELECT id, title FROM articles WHERE is_draft = '0' AND title LIKE '%" + search + "%'").Scan(&articles)

	return articles

}

func addSubscriber(email string, first string, last string) bool {
	subscriber := Subscriber{ID: email, FirstName: first, LastName: last}

	db, err := gorm.Open(sqlite.Open("skjsports.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	result := db.Create(&subscriber)

	if result.RowsAffected != 0 {
		return true
	} else {
		result := db.Unscoped().Model(&Subscriber{ID: email}).Where("deleted_at IS NOT NULL").Update("deleted_at", nil)

		if result.RowsAffected != 0 {
			db.Unscoped().Model(&Subscriber{ID: email}).Updates(Subscriber{FirstName: first, LastName: last})
			return true
		} else {
			return false
		}
	}
}

func removeSubscriber(email string) bool {
	db, err := gorm.Open(sqlite.Open("skjsports.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	result := db.Delete(&Subscriber{ID: email})
	// result := db.Where("id = ?", email).Delete(&Subscriber{})

	if result.RowsAffected != 0 {
		return true
	} else {
		return false
	}
}

func retrieveSubscriber(email string) *Subscriber {
	subscriber := new(Subscriber)

	db, err := gorm.Open(sqlite.Open("skjsports.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	result := db.First(&subscriber, "id = ?", email)

	if result.Error == gorm.ErrRecordNotFound {
		return nil // if record not found
	}

	return subscriber
}

func addDraftToDatabase(title string, content string, author_email string) string {
	db, err := gorm.Open(sqlite.Open("skjsports.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	article_id := shortID(16)
	test_article := Article{ID: article_id}
	result := db.First(&test_article)
	for result.Error != gorm.ErrRecordNotFound {
		article_id := shortID(16)
		test_article := Article{ID: article_id}
		result = db.First(&test_article)
	}

	author := new(Author)
	db.Where("author_email = ?", author_email).First(&author)

	draft := Article{ID: article_id, Title: title, Content: content, Author: author.Author, AuthorEmail: author_email}

	result = db.Create(&draft)

	return article_id
}
