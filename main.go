package main

import (
	"encoding/xml"
	"fmt"
	"github.com/bostigger/goole-trends-cli/controller"
	"github.com/bostigger/goole-trends-cli/model"
	"log"
)

const SEPARATOR = "--------------------------------------------------------"

func main() {
	wMsg := "Google Trends Cli Application"
	fmt.Printf("%s\n", wMsg)

	var rssFeed model.Rss

	data := controller.ReadGoogleTrends()
	err := xml.Unmarshal(data, &rssFeed)
	if err != nil {
		fmt.Println(err.Error())
	}

	for i := range rssFeed.Channel.Items {
		fmt.Println(SEPARATOR)
		fmt.Println("#", i+1)
		logItem(rssFeed.Channel.Items[i].Title, rssFeed.Channel.Items[i].Traffic, rssFeed.Channel.Items[i].Link)
		logNewsItem(rssFeed.Channel.Items[i].NewsItems)
	}

}

func logItem(title string, traffic string, url string) {
	log.Printf("Trend Title: %s Trend Traffic :%s  Trend Url: %s\n", title, traffic, url)
}

func logNewsItem(newsItems []model.NewsItem) {
	for _, newsItem := range newsItems {
		log.Println(newsItem.Title)
	}
}
