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

var currencyCodeNutiMap = map[string]string{
	"USD":     "US Dólar",
	"EUR":     "Euro",
	"EUR-LOW": "Euro 200 y 500",
	"ARS":     "Peso Argentino",
	"AUD":     "Dólar australiano",
	"BRL":     "real brasileño",
	"CAD":     "Dólar canadiense",
	"CHF":     "Franco suizo",
	"CLP":     "Peso Chileno",
	"CNY":     "Yuan Chino",
	"GBP":     "Libra Esterlina",
	"JPY":     "Yen japonés",
	"MXN":     "Peso mexicano",
	"PEN":     "Nuevo Sol peruano",
}

func (reqExchange *ExchangeNutifinanzas) selectExchange() ResultExchange {
	fmt.Println("----------- Nutifinanzas Currency Exchange ----------------")

	fmt.Printf("Request Currency Exchange - Nutifinanzas: %s - Value: %f - OperationType: %s \n",
		reqExchange.Exchange.CurrencyCode, reqExchange.Exchange.Value, reqExchange.Exchange.OperationType)

	var scraper = scraping.BuildGoRodScrapper(reqExchange.Url)

	currenciesNutifinanzas := make([]Currency, 0)

	scrapCurrenciesNutifinanzas(scraper, &currenciesNutifinanzas)

	currenciesNutifinanzas = HomologateCurrencyDespByCode(currenciesNutifinanzas, currencyCodeNutiMap)

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

		purchaseValue, _ := utils.FromStringToFloat(strings.Trim(valueToBuyStr, " "))
		onSaleValue, _ := utils.FromStringToFloat(strings.Trim(valueOnSaleStr, " "))

		currencyNutifinanzas := Currency{
			description: strings.Trim(description, " "),
			valueToBuy:  onSaleValue,
			valueOnSale: purchaseValue,
		}

		*currenciesNutifinanzas = append(*currenciesNutifinanzas, currencyNutifinanzas)

		fmt.Printf("Description: %s\nValue To Buy: %f\nValue On Sale: %f\n", currencyNutifinanzas.description, currencyNutifinanzas.valueToBuy, currencyNutifinanzas.valueOnSale)
	}

}
