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

func (reqExchange *ExchangeNutifinanzas) selectExchange() ResultExchange {
	fmt.Println("----------- Nutifinanzas Currency Exchange ----------------")

	fmt.Printf("Request Currency Exchange - Nutifinanzas: %s - Value: %f - OperationType: %s \n",
		reqExchange.Exchange.CurrencyCode, reqExchange.Exchange.Value, reqExchange.Exchange.OperationType)

	var scraper = scraping.BuildGoRodScrapper(reqExchange.Url)

	currenciesNutifinanzas := make([]Currency, 0)

	scrapCurrenciesNutifinanzas(scraper, &currenciesNutifinanzas)

	var resultExchange = CalculateConversion(currenciesNutifinanzas, &reqExchange.Exchange)

	fmt.Printf("Result Exchange Unicambios:\nCurrency: %s\nOperation Type: %s\nValue Operation: %f\nValue Convertion: %f\n",
		resultExchange.Exchange.CurrencyCode, resultExchange.OperationType, resultExchange.Value, resultExchange.ValueConvertion)

	return resultExchange
}

func scrapCurrenciesNutifinanzas(scraper *rod.Page, currenciesNutifinanzas *[]Currency) {

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

		currencyNutifinanzas := Currency{
			description: description,
			valueToBuy:  valueToBuy,
			valueOnSale: valueOnSale,
		}

		*currenciesNutifinanzas = append(*currenciesNutifinanzas, currencyNutifinanzas)

		fmt.Printf("Description: %s\nValue To Buy: %f\nValue On Sale: %f\n", currencyNutifinanzas.description, currencyNutifinanzas.valueToBuy, currencyNutifinanzas.valueOnSale)
	}

}
