package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Items       []RSSItem `xml:"item"`
	} `xml:"channel"`
}
type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubdate"`
}

func urlTofeed(url string) (*RSSFeed, error) {
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteData, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	rssFeed := RSSFeed{}
	errParse := xml.Unmarshal(byteData, &rssFeed)
	if errParse != nil {
		return nil, errParse
	}
	return &rssFeed, nil
}
