package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Slice []interface{}

func defaultRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("New GET request received.")
		fmt.Fprintf(w, "Success!")
	} else {
		fmt.Fprintf(w, "Unsupported request type.")
	}
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("New GET request received for article retrieval.")

		// Set headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		id := r.URL.Query().Get("id")        // Access query param
		result := getArticleContentsByID(id) // retrieve article as list of content elements (text/image)

		if result == nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			jsonRes, err := json.Marshal(result)

			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonRes)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(jsonRes)
			}
		}
	} else {
		fmt.Fprintf(w, "Unsupported request type.")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func subscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("New POST request received for new subscriber.")

		// Set headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)

		bodyMap := make(map[string]string)
		json.Unmarshal(body, &bodyMap)

		success := addSubscriber(bodyMap["email"], bodyMap["first_name"], bodyMap["last_name"])

		if success {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotModified) // subscriber already exists
		}

	} else {
		fmt.Fprintf(w, "Unsupported request type.")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func unsubscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("New POST request received for removing subscriber.")

		// Set headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)

		bodyMap := make(map[string]string)
		json.Unmarshal(body, &bodyMap)

		success := removeSubscriber(bodyMap["email"])

		if success {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotModified) // user wasn't subscribed under given email
		}

	} else {
		fmt.Fprintf(w, "Unsupported request type.")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getSearchResults(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("New GET request received for article search.")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		search := r.URL.Query().Get("search")
		res := getArticlesMatchingSearch(search)

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
	http.HandleFunc("/", defaultRoute)
	http.HandleFunc("/article", getArticle)
	http.HandleFunc("/search", getSearchResults)
	http.HandleFunc("/subscribe", subscribe)
	http.HandleFunc("/unsubscribe", unsubscribe)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
