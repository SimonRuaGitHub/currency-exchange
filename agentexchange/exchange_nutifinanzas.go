package interactions

import (
	scraping "currency-exchange-medellin/scraping"
	"currency-exchange-medellin/utils"
	"strings"

	"fmt"

	"github.com/go-rod/rod"
)

const (
	currenciesTarget  = "#currenciesContainer div.card-info"
	descriptionTarget = "span:nth-child(1)"
	toBuyTarget       = "span:nth-child(3)"
	onSaleTarget      = "span:nth-child(5)"
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

	var scraper = scraping.BuildGoRodScrapper(reqExchange.Url)

	currenciesNutifinanzas := make([]currencyNutifinanzas, 0)

	scrapCurrenciesNutifinanzas(scraper, &currenciesNutifinanzas)

	return ResultExchange{}
}

func scrapCurrenciesNutifinanzas(scraper *rod.Page, currenciesNutifinanzas *[]currencyNutifinanzas) {

	scraper.MustElement(currenciesTarget).ScrollIntoView()

	cards := scraper.MustElements(currenciesTarget)

	for _, card := range cards {
		description, _ := card.MustElement(descriptionTarget).Text()
		toBuy, _ := card.MustElement(toBuyTarget).Text()
		onSale, _ := card.MustElement(onSaleTarget).Text()

		valueToBuyStr := strings.Split(toBuy, "$")[1]
		valueOnSaleStr := strings.Split(onSale, "$")[1]

		valueToBuy, _ := utils.FromStringToFloat(strings.Trim(valueToBuyStr, " "))
		valueOnSale, _ := utils.FromStringToFloat(strings.Trim(valueOnSaleStr, " "))

		currencyNutifinanzas := currencyNutifinanzas{
			description: description,
			valueToBuy:  valueToBuy,
			valueOnSale: valueOnSale,
		}

		*currenciesNutifinanzas = append(*currenciesNutifinanzas, currencyNutifinanzas)

		fmt.Printf("Description: %s\nValue To Buy: %f\nValue On Sale: %f\n", currencyNutifinanzas.description, currencyNutifinanzas.valueToBuy, currencyNutifinanzas.valueOnSale)
	}

}
