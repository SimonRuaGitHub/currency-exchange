package interactions

import (
	scraping "currency-exchange-medellin/scraping"
	"currency-exchange-medellin/utils"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type ExchangeMaxmoney struct {
	RequestExchange
}

const (
	tableTarget          = "div.tab-content table tbody"
	rowTarget            = "tr"
	descriptionColTarget = "td:nth-child(1)"
	purchaseColTarget    = "td:nth-child(2)"
	onSaleColTarget      = "td:nth-child(3)"
)

func (reqExchange *ExchangeMaxmoney) selectExchange() ResultExchange {
	fmt.Println("----------- Moneymax Currency Exchange ----------------")

	fmt.Printf("Request Currency: %s - Value: %f - OperationType: %s \n",
		reqExchange.Exchange.CurrencyCode, reqExchange.Exchange.Value, reqExchange.Exchange.OperationType)

	var scraper = scraping.BuildCollyScrapper(scraping.DefaultTimeOutColly)

	currenciesMoneymax := make([]Currency, 0)

	scrapCurrenciesMoneymax(scraper, &currenciesMoneymax)

	scraper.Visit(reqExchange.Url)

	var resultExchange = CalculateConversion(currenciesMoneymax, &reqExchange.Exchange)

	fmt.Printf("Result Exchange Unicambios:\nCurrency: %s\nOperation Type: %s\nValue Operation: %f\nValue Convertion: %f\n",
		resultExchange.Exchange.CurrencyCode, resultExchange.OperationType, resultExchange.Value, resultExchange.ValueConvertion)

	return resultExchange
}

func scrapCurrenciesMoneymax(scraper *colly.Collector, currenciesMoneymax *[]Currency) {

	scraper.OnHTML(tableTarget, func(tableHtml *colly.HTMLElement) {

		fmt.Println("Table: ", tableTarget)

		tableHtml.ForEach(rowTarget, func(i int, row *colly.HTMLElement) {

			description := row.ChildText(descriptionColTarget)
			purchaseValueStr := strings.Replace(row.ChildText(purchaseColTarget), ".", "", -1)
			onSaleValueStr := strings.Replace(row.ChildText(onSaleColTarget), ".", "", -1)
			purchaseValue, _ := utils.FromStringToFloat(strings.Trim(purchaseValueStr, "$ "))
			onSaleValue, _ := utils.FromStringToFloat(strings.Trim(onSaleValueStr, "$ "))

			currencyMoneymax := Currency{
				Description: description,
				ValueToBuy:  onSaleValue,
				ValueOnSale: purchaseValue,
			}

			fmt.Println("Currency scraped: ", currencyMoneymax)

			*currenciesMoneymax = append(*currenciesMoneymax, currencyMoneymax)
		})
	})
}
