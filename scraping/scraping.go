package scraping

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
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

func BuildGoRodScrapperOnDebug(url string, enableHeadless bool, enableTrace bool, slowMotion time.Duration) *rod.Page {
	launcher := launcher.New().Headless(enableHeadless).StartURL(url)

	urlOnLauncher := launcher.MustLaunch()

	// Trace shows verbose debug information for each action executed
	// Slowmotion is a debug related function that waits 2 seconds between
	// each action, making it easier to inspect what your code is doing.
	browser := rod.New().
		ControlURL(urlOnLauncher).
		Trace(enableTrace).
		SlowMotion(slowMotion).
		MustConnect()

	page := browser.MustPage(url).MustWindowFullscreen()

	return page
}
