package main

import (
	"regexp"
	"strings"
)

type JsonMap map[string]interface{}

func processArticle(articleObj Article) JsonMap {
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
		"title":   articleObj.Title,
		"content": contentList,
	}

	return res
}
