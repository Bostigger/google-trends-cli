package model

import (
	"encoding/xml"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title         string     `xml:"title"`
	Traffic       string     `xml:"approx_traffic"`
	Link          string     `xml:"link"`
	PublishedDate string     `xml:"pubDate"`
	PictureURL    string     `xml:"picture"`
	NewsItems     []NewsItem `xml:"news_item"`
}

type NewsItem struct {
	Title   string `xml:"news_item_title"`
	Snippet string `xml:"news_item_snippet"`
	Url     string `xml:"news_item_url"`
	Source  string `xml:"news_item_source"`
}
