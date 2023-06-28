package model

import (
	"encoding/xml"
	"time"
)

type Rss struct {
	XMLName xml.Name
	Channel *Channel
}

type Channel struct {
	Title       string
	Description string
	Link        string
	Items       []Item
}

type Item struct {
	Title         string
	Traffic       string
	Link          string
	PublishedDate time.Time
	PictureURL    string
	NewsItems     []NewsItem
}

type NewsItem struct {
	Title   string
	Snippet string
	Url     string
	Source  string
}
