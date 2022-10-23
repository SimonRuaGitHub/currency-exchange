package interactions

import (
	scraping "currency-exchange-medellin/scraping"
	"fmt"

	"github.com/gocolly/colly"
)

const (
	currContainerTarget   = "div.row > div"
	currCardsTarget       = "div > div div"
	currDescriptionTarget = "span:nth-child(1)"
	toBuyTarget           = "span:nth-child(2)"
	onSaleTarget          = "span:nth-child(3)"
)

type ExchangeNutifinanzas struct {
	RequestExchange
}

type currencyNutifinanzas struct {
	description string
	valueOnSale float64
	valueToBuy  float64
}

func (reqExchange *ExchangeNutifinanzas) selectExchange() ResultExchange {
	fmt.Println("----------- Nutifinanzas Currency Exchange ----------------")

	fmt.Printf("Request Currency Exchange - Nutifinanzas: %s - Value: %f - OperationType: %s \n",
		reqExchange.Exchange.Currency, reqExchange.Exchange.Value, reqExchange.Exchange.OperationType)

	var scraper = scraping.BuildScraper(scrappingTimeout)

	currenciesNutifinanzas := make([]currencyNutifinanzas, 0)

	scrapCurrenciesNutifinanzas(scraper, &currenciesNutifinanzas)

	scraper.Visit(reqExchange.Url)

	return ResultExchange{}
}

func scrapCurrenciesNutifinanzas(scraper *colly.Collector, currenciesNutifinanzas *[]currencyNutifinanzas) {

	/*scraper.OnHTML("#currenciesContainer", func(cardsHtml *colly.HTMLElement) {
		fmt.Println("CardsHtml: ", cardsHtml)
		cardsHtml.ForEach("div", func(i int, currencyCard *colly.HTMLElement) {
			fmt.Println("CurrencyCard: ", currencyCard)
			valueToBuyStr := strings.Split(currencyCard.ChildText(toBuyTarget), "$")[1]
			valueOnSaleStr := strings.Split(currencyCard.ChildText(onSaleTarget), "$")[1]
			valueToBuy, _ := utils.FromStringToFloat(strings.Trim(valueToBuyStr, " "))
			valueOnSale, _ := utils.FromStringToFloat(strings.Trim(valueOnSaleStr, " "))

			currencyNutifinanzas := currencyNutifinanzas{
				description: currencyCard.ChildText(currDescriptionTarget),
				valueToBuy:  valueToBuy,
				valueOnSale: valueOnSale,
			}

			fmt.Println("Currency scraped: ", currencyNutifinanzas)

			*currenciesNutifinanzas = append(*currenciesNutifinanzas, currencyNutifinanzas)
		})
	})*/

	scraper.OnHTML("div#currenciesContainer", func(currenciesContainer *colly.HTMLElement) {
		fmt.Println("currencies container: ", currenciesContainer)
	})

}
