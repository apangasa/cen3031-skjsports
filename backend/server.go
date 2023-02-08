package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type JsonMap map[string]interface{}
type Slice []interface{}

func default_route(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("New GET request received.")
		fmt.Fprintf(w, "Success!")
	} else {
		fmt.Fprintf(w, "Unsupported request type.")
	}
}

func article_retrieval(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("New GET request received for article retrieval.")

		w.Header().Set("Content-Type", "application/json")

		id := r.URL.Query().Get("id")
		article := retrieveArticle(id)

		re := regexp.MustCompile(`<image [\w]+>`)
		//content := re.Split(article.Content, -1)
		content := article.Content

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
			imageID := strings.Split(strings.Split(imageString, "<image ")[0], ">")[0]

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
			"title":   article.Title,
			"content": contentList,
		}

		fmt.Println(contentList)

		jsonRes, err := json.Marshal(res)

		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Write(jsonRes)
	} else {
		fmt.Fprintf(w, "Unsupported request type.")
	}
}

func main() {
	http.HandleFunc("/", default_route)
	http.HandleFunc("/article", article_retrieval)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
