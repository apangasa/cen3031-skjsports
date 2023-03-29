package main

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func getPlayerStats(player_name string) map[string]string {
	positions := getPositions(player_name)

	var stats map[string]string

	if strings.Contains(positions, "GK") {
		stats = getGoalkeeperStats(player_name, stats)
	}

	if strings.Contains(positions, "DF") {
		stats = getDefenderStats(player_name, stats)
	}

	if strings.Contains(positions, "MF") {
		stats = getMidfielderStats(player_name, stats)
	}

	if strings.Contains(positions, "FW") {
		stats = getForwardStats(player_name, stats)
	}

	return stats
}

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

func getGoalkeeperStats(player_name string, stats map[string]string) map[string]string {
	return stats
}

func getDefenderStats(player_name string, stats map[string]string) map[string]string {
	return stats
}

func getMidfielderStats(player_name string, stats map[string]string) map[string]string {
	return stats
}

func getForwardStats(player_name string, stats map[string]string) map[string]string {
	return stats
}
