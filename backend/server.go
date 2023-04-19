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

		id := r.URL.Query().Get("id")               // Access query param
		result := getArticleContentsByID(id, false) // retrieve article as list of content elements (text/image)

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

func getDraft(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("New GET request received for article retrieval.")

		// Set headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		id := r.URL.Query().Get("id")              // Access query param
		result := getArticleContentsByID(id, true) // retrieve article as list of content elements (text/image)

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

func getArticlesByAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("New GET request received for article retrieval by author.")

		// Set headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		id := r.URL.Query().Get("author_id")          // Access query param
		res := getArticlesMatchingAuthorId(id, false) // TODO replace with getting articles with matching article id

		jsonRes, err := json.Marshal(res)

		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Write(jsonRes)
	} else {
		fmt.Fprintf(w, "Unsupported request type.")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getDraftsByAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("New GET request received for article retrieval by author.")

		// Set headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		id := r.URL.Query().Get("author_id")         // Access query param
		res := getArticlesMatchingAuthorId(id, true) // TODO replace with getting articles with matching article id

		jsonRes, err := json.Marshal(res)

		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Write(jsonRes)
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
		res := getArticlesMatchingSearch(search, false)

		jsonRes, err := json.Marshal(res)

		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Write(jsonRes)
	} else {
		fmt.Fprintf(w, "Unsupported request type.")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getPlayerStats(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("New GET request received for player stats.")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		player_name := r.URL.Query().Get("player_name")
		res := webScrapePlayerStats(player_name)

		jsonRes, err := json.Marshal(res)

		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Write(jsonRes)

	} else {
		fmt.Fprintf(w, "Unsupported request type.")
	}
}

func getTeamStats(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("New GET request received for player stats.")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		team_name := r.URL.Query().Get("team_name")
		res := webScrapeTeamStats(team_name)

		jsonRes, err := json.Marshal(res)

		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Write(jsonRes)

	} else {
		fmt.Fprintf(w, "Unsupported request type.")
	}
}

func createDraft(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("New GET request received for player stats.")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)

		var bodyMap JsonMap
		json.Unmarshal(body, &bodyMap)

		content := wrapArticleContents(bodyMap["content"].([]map[string]string))

		article_id := addDraftToDatabase(bodyMap["title"].(string), content, bodyMap["author_email"].(string))

		res := JsonMap{
			"article_id": article_id,
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

func editDraft(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("New GET request received for player stats.")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)

		var bodyMap JsonMap
		json.Unmarshal(body, &bodyMap)

		content := wrapArticleContents(bodyMap["content"].([]map[string]string))

		success := updateDraftInDatabase(bodyMap["article_id"].(string), content)

		if success {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest) // id does not exist
		}

	} else {
		fmt.Fprintf(w, "Unsupported request type.")
	}
}

func publishDraft(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("New GET request received for player stats.")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)

		var bodyMap JsonMap
		json.Unmarshal(body, &bodyMap)

		success := convertDraftToArticle(bodyMap["article_id"].(string))

		if success {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest) // id does not exist
		}

	} else {
		fmt.Fprintf(w, "Unsupported request type.")
	}
}

func main() {
	http.HandleFunc("/", defaultRoute)
	http.HandleFunc("/article", getArticle)
	http.HandleFunc("/articles", getArticlesByAuthor)
	http.HandleFunc("/search", getSearchResults)
	http.HandleFunc("/subscribe", subscribe)
	http.HandleFunc("/unsubscribe", unsubscribe)
	http.HandleFunc("/stats/player", getPlayerStats)
	http.HandleFunc("/stats/team", getTeamStats)

	http.HandleFunc("/create-draft", createDraft)
	http.HandleFunc("/edit-draft", editDraft)
	http.HandleFunc("/publish-draft", publishDraft)

	http.HandleFunc("/draft", getArticle)
	http.HandleFunc("/drafts", getArticlesByAuthor)

	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/authenticate", Auth)
	http.HandleFunc("/renew", Renew)
	http.HandleFunc("/addWriter", addWriter)
	http.HandleFunc("/logout", Logout)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
