package main

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func getPositions(player_name string) string {
	fmt.Println("Getting position...")
	c := colly.NewCollector()

	formData := url.Values{}
	formData.Set("search", player_name)

	queryString := formData.Encode()
	fbrefURL := "https://fbref.com/en/search/search.fcgi?" + queryString

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

	c.Visit(fbrefURL)

	return positions
}
