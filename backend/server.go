package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Slice []interface{}

func default_route(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("New GET request received.")
		fmt.Fprintf(w, "Success!")
	} else {
		fmt.Fprintf(w, "Unsupported request type.")
	}
}

func get_article(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("New GET request received for article retrieval.")

		// Set headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		id := r.URL.Query().Get("id")        // Access query param
		articleObj := retrieveArticle(id)    // Get struct of Article by querying DB
		result := processArticle(articleObj) // Dissect article into text and image content list
		jsonRes, err := json.Marshal(result)

		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
	} else {
		fmt.Fprintf(w, "Unsupported request type.")
	}
}

func article_search(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("New GET request received for article search.")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		search := r.URL.Query().Get("search")
		articles := searchArticle(search)

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
	http.HandleFunc("/article", get_article)
	http.HandleFunc("/search", article_search)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
