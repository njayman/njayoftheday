package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Namespace struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Titles struct {
	Canonical  string `json:"canonical"`
	Normalized string `json:"normalized"`
	Display    string `json:"display"`
}

type Image struct {
	Source string `json:"source"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type URLs struct {
	Page      string `json:"page"`
	Revisions string `json:"revisions"`
	Edit      string `json:"edit"`
	Talk      string `json:"talk"`
}

type ContentURLs struct {
	Desktop URLs `json:"desktop"`
	Mobile  URLs `json:"mobile"`
}

type Page struct {
	Type              string      `json:"type"`
	Title             string      `json:"title"`
	DisplayTitle      string      `json:"displaytitle"`
	Namespace         Namespace   `json:"namespace"`
	WikibaseItem      string      `json:"wikibase_item"`
	Titles            Titles      `json:"titles"`
	PageID            int         `json:"pageid"`
	Thumbnail         Image       `json:"thumbnail"`
	OriginalImage     Image       `json:"originalimage"`
	Lang              string      `json:"lang"`
	Dir               string      `json:"dir"`
	Revision          string      `json:"revision"`
	TID               string      `json:"tid"`
	Timestamp         string      `json:"timestamp"`
	Description       string      `json:"description"`
	DescriptionSource string      `json:"description_source"`
	Coordinates       Coordinates `json:"coordinates"`
	ContentURLs       ContentURLs `json:"content_urls"`
	Extract           string      `json:"extract"`
	ExtractHTML       string      `json:"extract_html"`
	NormalizedTitle   string      `json:"normalizedtitle"`
}

type Data struct {
	Text  string `json:"text"`
	Pages []Page `json:"pages"`
	Year  int    `json:"year"`
}

func main() {

	date := time.Now()

	formatedDate := date.Format("01/02")

	requestURL := fmt.Sprintf("https://api.wikimedia.org/feed/v1/wikipedia/en/onthisday/all/%v", formatedDate)

	fmt.Print(requestURL)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Print(res.Body)

	response := res.Body

	fmt.Print(response)

}
