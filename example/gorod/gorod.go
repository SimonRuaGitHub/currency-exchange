package main

import (
	"github.com/go-rod/rod"
)

func main() {
	browser := rod.New().MustConnect().NoDefaultDevice()
	page := browser.MustPage("https://www.wikipedia.org/").MustWindowFullscreen()

	page.MustElement("#searchInput").MustInput("earth")

	page.MustWaitLoad().MustScreenshot("a.png")
}
