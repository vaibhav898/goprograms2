package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	
	res, err := http.Get("https://stackoverflow.com/questions/tagged/javascript")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	
	doc.Find(".question-summary .summary").Each(func(i int, s *goquery.Selection) {
		title := s.Find("H3").Text()
		fmt.Println(i, title)
	})
}

func main() {
	ExampleScrape()
}
