package scraping

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func BuildScraper(timeout time.Duration) *colly.Collector {
	scraper := colly.NewCollector()

	scraper.SetRequestTimeout(timeout)

	scraper.OnRequest(func(request *colly.Request) {
		fmt.Println("Visting web page: " + request.URL.String())
	})

	scraper.OnResponse(func(r *colly.Response) {
		fmt.Println("Got response from page")
	})

	scraper.OnError(func(r *colly.Response, e error) {
		fmt.Println("Error when visiting page: ", e)
	})

	scraper.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished web scraping")
	})

	return scraper
}
