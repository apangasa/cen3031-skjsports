package main

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func webScrapeTeamStats(team_name string) map[string]string {
	return nil
}

func webScrapePlayerStats(player_name string) map[string]string {
	fmt.Println("Getting position...")

	formData := url.Values{}
	formData.Set("search", player_name)

	queryString := formData.Encode()
	fbrefURL := "https://fbref.com/en/search/search.fcgi?" + queryString

	positions := getPositions(fbrefURL)

	stats := make(map[string]string)

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

func getPositions(fbrefURL string) string {
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
	})

	c.OnHTML("p:has(strong:contains('National Team'))", func(e *colly.HTMLElement) {
		text := e.Text
		text = strings.Replace(text, "National Team:", "", 1)
		text = strings.Split(text, " ")[0]
		reg, err := regexp.Compile("[^a-zA-Z]+")
		if err != nil {
			fmt.Println("error")
		}

		text = reg.ReplaceAllString(text, " ")
		text = strings.TrimSpace(text)

		stats["Nation"] = text
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
	})

	c.OnHTML("p:has(strong:contains('National Team'))", func(e *colly.HTMLElement) {
		text := e.Text
		text = strings.Replace(text, "National Team:", "", 1)
		text = strings.Split(text, " ")[0]
		reg, err := regexp.Compile("[^a-zA-Z]+")
		if err != nil {
			fmt.Println("error")
		}

		text = reg.ReplaceAllString(text, " ")
		text = strings.TrimSpace(text)

		stats["Nation"] = text
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
	})

	c.OnHTML("p:has(strong:contains('National Team'))", func(e *colly.HTMLElement) {
		text := e.Text
		text = strings.Replace(text, "National Team:", "", 1)
		text = strings.Split(text, " ")[0]
		reg, err := regexp.Compile("[^a-zA-Z]+")
		if err != nil {
			fmt.Println("error")
		}

		text = reg.ReplaceAllString(text, " ")
		text = strings.TrimSpace(text)

		stats["Nation"] = text
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
	})

	c.OnHTML("p:has(strong:contains('National Team'))", func(e *colly.HTMLElement) {
		text := e.Text
		text = strings.Replace(text, "National Team:", "", 1)
		text = strings.Split(text, " ")[0]
		reg, err := regexp.Compile("[^a-zA-Z]+")
		if err != nil {
			fmt.Println("error")
		}

		text = reg.ReplaceAllString(text, " ")
		text = strings.TrimSpace(text)

		stats["Nation"] = text
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
