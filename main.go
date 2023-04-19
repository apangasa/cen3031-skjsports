package main

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func webScrapeTeamStats(team_name string) map[string]string {
	formData := url.Values{}
	formData.Set("search", team_name)

	queryString := formData.Encode()
	fbrefURL := "https://fbref.com/en/search/search.fcgi?" + queryString

	stats := make(map[string]string)

	fmt.Println("Getting statistics for ", team_name)

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)

		if r.Request.URL.Path == "/en/search/search.fcgi" {
			searchDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(r.Body))
			if err != nil {
				log.Println("Failed to parse search page response: ", err)
				return
			}

			resultLink := searchDoc.Find(".search-item-name a[href]").First()
			if resultLink.Length() == 0 {
				log.Println("No results found for search term")
				return
			}

			resultURL, exists := resultLink.Attr("href")
			if !exists {
				log.Println("Failed to extract result URL")
				return
			}

			c.Visit("https://fbref.com" + resultURL)
		}
	})

	c.OnHTML("p:contains('Record:'):not(:contains('Home Record:'), :contains('Away Record:'))", func(e *colly.HTMLElement) {
		text := e.Text
		text = strings.Replace(text, "Record:", "", 1)
		text = strings.Split(text, ",")[0]
		text = strings.TrimSpace(text)
		var arr [3]string
		temp := ""
		index := 0

		for i := 0; i < len(text); i++ {
			if text[i] == 45 {
				arr[index] = temp
				temp = ""
				index++
				continue
			}
			temp += string(text[i])
			if i == len(text)-1 {
				arr[index] = temp
			}
		}

		wins, err1 := strconv.Atoi(arr[0])
		losses, err2 := strconv.Atoi(arr[1])
		draws, err3 := strconv.Atoi(arr[2])

		if err1 != nil {
			panic(err1)
		}

		if err2 != nil {
			panic(err2)
		}

		if err3 != nil {
			panic(err3)
		}

		games := wins + losses + draws
		gamesStr := strconv.Itoa(games)

		stats["Total games played"] = gamesStr
		stats["Wins"] = arr[0]
		stats["Losses"] = arr[1]
		stats["Draws"] = arr[2]

	})

	c.OnHTML("table#stats_keeper_12 > tfoot", func(e *colly.HTMLElement) {
		var text string
		var text2 string
		temp := true

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			if temp {
				text = el.ChildText("td:nth-child(9)")
				temp = false
			} else {
				text2 = el.ChildText("td:nth-child(9)")
				temp = true
			}
		})

		if temp {
			stats["Goals scored"] = text2
			stats["Goals conceced"] = text
		} else {
			stats["Goals conceded"] = text
			stats["Goals scored"] = text2
		}

	})

	c.Visit(fbrefURL)

	return stats
}

func webScrapePlayerStats(player_name string) map[string]string {
	fmt.Println("Getting position...")

	formData := url.Values{}
	formData.Set("search", player_name)

	queryString := formData.Encode()
	fbrefURL := "https://fbref.com/en/search/search.fcgi?" + queryString

	stats := make(map[string]string)
	stats_pointer := &stats

	positions := getPositions(fbrefURL, stats_pointer)
	temp := ""

	//iterate through positions string, append to a final string until the character
	//is a space, or a 32
	for i := 0; i < len(positions); i++ {
		if positions[i] == 32 {
			break
		}
		temp += string(positions[i])
	}

	positions = temp

	fmt.Println("Getting statistics for ", player_name)

	if strings.Contains(positions, "GK") {
		stats = getGoalkeeperStats(fbrefURL, stats)
	}

	if strings.Contains(positions, "DF") {
		stats = getDefenderStats(fbrefURL, stats)
	}

	if strings.Contains(positions, "MF") {
		stats = getMidfielderStats(fbrefURL, stats)
	}

	if strings.Contains(positions, "FW") {
		stats = getForwardStats(fbrefURL, stats)
	}

	return stats
}

func getPositions(fbrefURL string, stats *map[string]string) string {
	c := colly.NewCollector()

	selector := "p:has(strong:contains('Position'))"

	var positions string

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)

		if r.Request.URL.Path == "/en/search/search.fcgi" {
			searchDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(r.Body))
			if err != nil {
				log.Println("Failed to parse search page response: ", err)
				return
			}

			resultLink := searchDoc.Find(".search-item-name a[href]").First()
			if resultLink.Length() == 0 {
				log.Println("No results found for search term")
				return
			}

			resultURL, exists := resultLink.Attr("href")
			if !exists {
				log.Println("Failed to extract result URL")
				return
			}

			c.Visit("https://fbref.com" + resultURL)
		}
	})

	c.OnHTML(selector, func(e *colly.HTMLElement) {
		text := e.Text
		text = strings.Replace(text, "Position: ", "", 1)
		text = strings.Split(text, "Footed")[0]
		reg, err := regexp.Compile("[^a-zA-Z]+")
		if err != nil {
			panic(err)
		}
		text = reg.ReplaceAllString(text, " ")
		text = strings.TrimSpace(text)

		positions = text
	})

	c.OnHTML("p:has(strong:contains('National Team'))", func(e *colly.HTMLElement) {
		text := e.Text
		//text = strings.Replace(text, "National Team:", "", 1)
		text = strings.Split(text, ":")[1]
		arr := strings.Split(text, " ")
		text = ""
		for i := 0; i < len(arr); i++ {
			if i == len(arr)-1 {
				text = text[0 : len(text)-1]
				break
			}
			text += arr[i]
			text += " "
		}

		reg, err := regexp.Compile("[^a-zA-Z]+")
		if err != nil {
			fmt.Println("error")
		}

		text = reg.ReplaceAllString(text, " ")
		text = strings.TrimSpace(text)

		(*stats)["Nation"] = text
	})

	c.Visit(fbrefURL)

	return positions
}

func getGoalkeeperStats(fbrefURL string, stats map[string]string) map[string]string {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)

		if r.Request.URL.Path == "/en/search/search.fcgi" {
			searchDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(r.Body))
			if err != nil {
				log.Println("Failed to parse search page response: ", err)
				return
			}

			resultLink := searchDoc.Find(".search-item-name a[href]").First()
			if resultLink.Length() == 0 {
				log.Println("No results found for search term")
				return
			}

			resultURL, exists := resultLink.Attr("href")
			if !exists {
				log.Println("Failed to extract result URL")
				return
			}

			c.Visit("https://fbref.com" + resultURL)
		}
	})

	c.OnHTML("table#stats_keeper_dom_lg > tbody", func(e *colly.HTMLElement) {
		var text string

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(7)")
		})

		stats["Matches Played"] = text

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(19)")
		})

		stats["Total Clean Sheets"] = text

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(15)")
		})

		stats["Save Percentage"] = text

	})

	c.Visit(fbrefURL)

	return stats
}

func getDefenderStats(fbrefURL string, stats map[string]string) map[string]string {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)

		if r.Request.URL.Path == "/en/search/search.fcgi" {
			searchDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(r.Body))
			if err != nil {
				log.Println("Failed to parse search page response: ", err)
				return
			}

			resultLink := searchDoc.Find(".search-item-name a[href]").First()
			if resultLink.Length() == 0 {
				log.Println("No results found for search term")
				return
			}

			resultURL, exists := resultLink.Attr("href")
			if !exists {
				log.Println("Failed to extract result URL")
				return
			}

			c.Visit("https://fbref.com" + resultURL)
		}
	})

	c.OnHTML("table#stats_standard_dom_lg > tbody", func(e *colly.HTMLElement) {
		var text string

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(7)")
		})

		stats["Matches Played"] = text

	})

	c.OnHTML("table#stats_defense_dom_lg > tbody", func(e *colly.HTMLElement) {
		var text string

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(8)")
		})

		stats["Total Tackles"] = text

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(20)")
		})

		stats["Total Interceptions"] = text

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(15)")
		})

		stats["Tackle Percentage"] = text

	})

	c.Visit(fbrefURL)

	return stats
}

func getMidfielderStats(fbrefURL string, stats map[string]string) map[string]string {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)

		if r.Request.URL.Path == "/en/search/search.fcgi" {
			searchDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(r.Body))
			if err != nil {
				log.Println("Failed to parse search page response: ", err)
				return
			}

			resultLink := searchDoc.Find(".search-item-name a[href]").First()
			if resultLink.Length() == 0 {
				log.Println("No results found for search term")
				return
			}

			resultURL, exists := resultLink.Attr("href")
			if !exists {
				log.Println("Failed to extract result URL")
				return
			}

			c.Visit("https://fbref.com" + resultURL)
		}
	})

	c.OnHTML("table#stats_standard_dom_lg > tbody", func(e *colly.HTMLElement) {
		var text string

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(7)")
		})

		stats["Matches Played"] = text

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(11)")
		})

		stats["Total Goals"] = text

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(12)")
		})

		stats["Total Assists"] = text

	})

	c.OnHTML("table#stats_passing_dom_lg > tbody", func(e *colly.HTMLElement) {
		var text string

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(10)")
		})

		stats["Completion Percentage"] = text

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(24)")
		})

		stats["Total xA"] = text

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(7)")
		})

		stats["Matches Played"] = text
	})

	c.OnHTML("table#stats_defense_dom_lg > tbody", func(e *colly.HTMLElement) {
		var text string

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(8)")
		})

		stats["Total Tackles"] = text

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(20)")
		})

		stats["Total Interceptions"] = text
	})

	c.Visit(fbrefURL)

	return stats
}

func getForwardStats(fbrefURL string, stats map[string]string) map[string]string {

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)

		if r.Request.URL.Path == "/en/search/search.fcgi" {
			searchDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(r.Body))
			if err != nil {
				log.Println("Failed to parse search page response: ", err)
				return
			}

			resultLink := searchDoc.Find(".search-item-name a[href]").First()
			if resultLink.Length() == 0 {
				log.Println("No results found for search term")
				return
			}

			resultURL, exists := resultLink.Attr("href")
			if !exists {
				log.Println("Failed to extract result URL")
				return
			}

			c.Visit("https://fbref.com" + resultURL)
		}
	})

	c.OnHTML("table#stats_standard_dom_lg > tbody", func(e *colly.HTMLElement) {
		var text string

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(7)")
		})

		stats["Matches Played"] = text

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(11)")
		})

		stats["Total Goals"] = text

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(19)")
		})

		stats["Total xG"] = text

	})

	c.OnHTML("table#stats_passing_dom_lg > tbody", func(e *colly.HTMLElement) {
		var text string

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			text = el.ChildText("td:nth-child(24)")
		})

		stats["Total xA"] = text

	})

	c.Visit(fbrefURL)

	return stats
}

func main() {
	name := "itumeleng khune"
	data := make(map[string]string)
	data = webScrapePlayerStats(name)

	for key, val := range data {
		fmt.Printf("Key: %s, Value: %s\n", key, val)
	}
}
