package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
)

type news struct {
	Subject     string `json:"home-title"`
	Description string `json:"home-desc"`
	Link        string `json:"story-link"`
}

func writeJSON(data []news) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create JSON file")
		return
	}
	_ = ioutil.WriteFile("news.json", file, 0644)
}

func main() {
	allNews := make([]news, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("thehackernews.com"), // Add an Array made from CSV that has multiple domains.
	)
	collector.OnHTML(".body-post", func(element *colly.HTMLElement) {
		newsSubject := element.ChildText("h2.home-title")
		newsDescription := element.ChildText("div.home-desc")
		newsLink := element.ChildAttr("a", "href")
		news := news{
			Subject:     newsSubject,
			Description: newsDescription,
			Link:        newsLink,
		}
		allNews = append(allNews, news)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})
	err := collector.Visit("https://thehackernews.com/")
	if err != nil {
		log.Println("Could not visit the listed website.")
	}

	writeJSON(allNews)

	// Todo Add domains via CSV to Array Variable
	// Todo Add switch function to switch between multiple domains
	// Todo Change from exporting to JSON to CSV
}
