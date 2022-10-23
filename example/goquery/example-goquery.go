package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://nutifinanzas.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	// Find the review items
	s := doc.Find("div.card-info")

	fmt.Println("Currency card: ", s)
	/*.Each(func(i int, s *goquery.Selection) {

		fmt.Println("Currency card: ", s)
		// For each item found, get the title
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})*/
}

func main() {
	ExampleScrape()
}
