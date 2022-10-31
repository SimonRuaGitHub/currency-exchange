package scraping

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/gocolly/colly"
)

const DefaultTimeOutColly = 120 * time.Second

func BuildCollyScrapper(timeout time.Duration) *colly.Collector {
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

func BuildGoRodScrapper(url string) *rod.Page {
	browser := rod.New().MustConnect().NoDefaultDevice()
	page := browser.MustPage(url).MustWindowFullscreen()

	return page
}
