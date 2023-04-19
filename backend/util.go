package main

import (
	"crypto/rand"
	"regexp"
	"strings"
)

type JsonMap map[string]interface{}

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"

func shortID(length int) string {
	ll := len(chars)
	b := make([]byte, length)
	rand.Read(b) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%ll]
	}
	return string(b)
}

func getArticleContentsByID(id string, draft bool) JsonMap {
	var is_draft int
	if draft {
		is_draft = 1
	} else {
		is_draft = 0
	}

	articleObj := retrieveArticle(id, is_draft) // Get struct of Article by querying DB

	if articleObj == nil {
		return nil
	}

	result := processArticle(articleObj) // Dissect article into text and image content list
	return result
}

func wrapArticleContents(contentList []map[string]string) string {
	content := ""

	for i := 0; i < len(contentList); i++ {
		if contentList[i]["contentType"] == "text" {
			content += contentList[i]["text"]
		} else if contentList[i]["contentType"] == "img" {
			content += "<image " + contentList[i]["id"] + ">"
		} else {
			content += "" // unrecognized element
		}
	}

	return content
}

func processArticle(articleObj *Article) JsonMap {
	re := regexp.MustCompile(`<image [\w]+>`)
	content := articleObj.Content

	imageIndices := re.FindAllStringIndex(content, -1)

	var contentList []map[string]string

	j := 0
	for i := 0; i < len(imageIndices); i++ {
		nextImageStart := imageIndices[i][0]
		nextImageEnd := imageIndices[i][1]

		if nextImageStart > j {
			text := content[j:nextImageStart]

			contentMap := make(map[string]string)
			contentMap["contentType"] = "text"
			contentMap["text"] = text
			contentList = append(contentList, contentMap)
		}

		imageString := content[nextImageStart:nextImageEnd]

		imageID := strings.Split(strings.Split(imageString, "<image ")[1], ">")[0]

		contentMap := make(map[string]string)
		contentMap["contentType"] = "img"
		contentMap["id"] = imageID
		contentList = append(contentList, contentMap)

		j = nextImageEnd
	}

	text := content[j:]

	contentMap := make(map[string]string)
	contentMap["contentType"] = "text"
	contentMap["text"] = text
	contentList = append(contentList, contentMap)

	res := JsonMap{
		"title":        articleObj.Title,
		"content":      contentList,
		"author":       articleObj.Author,
		"author_email": articleObj.AuthorEmail,
	}

	return res
}

func getArticlesMatchingAuthorId(author_id string, draft bool) JsonMap {
	var is_draft int
	if draft {
		is_draft = 1
	} else {
		is_draft = 0
	}

	articles := retrieveAuthorArticles(author_id, is_draft)
	if articles == nil {
		return nil
	}

	res := formatArticles(articles)
	return res
}

func getArticlesMatchingSearch(search string, draft bool) JsonMap {
	var is_draft int
	if draft {
		is_draft = 1
	} else {
		is_draft = 0
	}

	articles := searchDatabaseForArticles(search, is_draft)
	if articles == nil {
		return nil
	}

	res := formatArticles(articles)
	return res
}

func formatArticles(articles []Article) JsonMap {
	var resultsList []map[string]string

	for i := 0; i < len(articles); i++ {
		titleMap := make(map[string]string)
		titleMap["id"] = articles[i].ID
		titleMap["title"] = articles[i].Title
		resultsList = append(resultsList, titleMap)
	}

	res := JsonMap{
		"results": resultsList,
	}

	return res
}
