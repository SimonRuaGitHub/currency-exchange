package main

import (
	"fmt"

	"github.com/go-rod/rod"
)

func main() {
	browser := rod.New().MustConnect().NoDefaultDevice()
	page := browser.MustPage("https://nutifinanzas.com").MustWindowFullscreen()

	page.MustElement("#currenciesContainer div.card-info").ScrollIntoView()

	cards := page.MustElements("#currenciesContainer div.card-info")

	for _, card := range cards {
		description, _ := card.MustElement("span:nth-child(1)").Text()
		valueToBuy, _ := card.MustElement("span:nth-child(3)").Text()
		valueOnSale, _ := card.MustElement("span:nth-child(5)").Text()

		fmt.Printf("Description: %s\nValue To Buy: %s\nValue On Sale: %s\n", description, valueToBuy, valueOnSale)
	}

	//page.MustWaitLoad().MustScreenshot("a.png")
}
