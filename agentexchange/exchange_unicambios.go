package interactions

import (
	scraping "currency-exchange-medellin/scraping"
	utils "currency-exchange-medellin/utils"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

const scrappingTimeout = 120 * time.Second

const (
	currenciesRows = "tr"
)

var currenciesTables = map[string]string{
	"tableLeftSide":  "section.premium-tabs-section table#supsystic-table-12 tbody",
	"tableRightSide": "section.premium-tabs-section table#supsystic-table-11 tbody",
}

type ExchangeUnicambios struct {
	RequestExchange
}

type currencyUnicambios struct {
	description string
	valueOnSale float64
	valueToBuy  float64
}

func (reqExchange *ExchangeUnicambios) selectExchange() ResultExchange {

	fmt.Println("----------- Unicambios Currency Exchange ----------------")

	fmt.Printf("Request Currency: %s - Value: %f - OperationType: %s \n",
		reqExchange.Exchange.Currency, reqExchange.Exchange.Value, reqExchange.Exchange.OperationType)

	var scraper = scraping.BuildCollyScrapper(scrappingTimeout)

	currenciesUnicambios := make([]currencyUnicambios, 0)

	scrapCurrenciesUnicambios(scraper, &currenciesUnicambios)

	scraper.Visit(reqExchange.Url)

	var resultExchange = calculateConversion(currenciesUnicambios, reqExchange)

	fmt.Printf("Result Exchange Unicambios:\nCurrency: %s\nOperation Type: %s\nValue Operation: %f\nValue Convertion: %f\n",
		resultExchange.Exchange.Currency, resultExchange.OperationType, resultExchange.Value, resultExchange.ValueConvertion)

	return resultExchange
}

func calculateConversion(currenciesInfo []currencyUnicambios, reqExchange *ExchangeUnicambios) ResultExchange {
	var valueConvertion float64 = 0.0
	var valueOperation float64 = 0.0

	for _, currencyInfo := range currenciesInfo {
		if strings.Contains(currencyInfo.description, reqExchange.Exchange.Currency) {
			fmt.Println("Found currency: ", currencyInfo.description)

			if reqExchange.Exchange.OperationType == "purchase" {
				valueConvertion = currencyInfo.valueToBuy * reqExchange.Exchange.Value
				valueOperation = currencyInfo.valueToBuy
			} else {
				valueConvertion = currencyInfo.valueOnSale * reqExchange.Exchange.Value
				valueOperation = currencyInfo.valueOnSale
			}
			break
		}
	}

	return ResultExchange{
		Exchange{reqExchange.Currency, valueOperation, reqExchange.OperationType},
		valueConvertion,
	}
}

func scrapCurrenciesUnicambios(scraper *colly.Collector, currenciesUnicambios *[]currencyUnicambios) {

	for tableSide, table := range currenciesTables {

		fmt.Println("Scraping currencies from table: ", tableSide)

		scraper.OnHTML(table, func(tableHtml *colly.HTMLElement) {
			fmt.Println("Table: ", table)
			tableHtml.ForEach(currenciesRows, func(i int, row *colly.HTMLElement) {
				if row.ChildText("td:nth-child(2)") != "" {

					targetValueToBuy := strings.Replace(row.ChildText("td:nth-child(3)"), ".", "", -1)
					targetValueOnSale := strings.Replace(row.ChildText("td:nth-child(4)"), ".", "", -1)
					valueToBuy, _ := utils.FromStringToFloat(targetValueToBuy)
					valueOnSale, _ := utils.FromStringToFloat(targetValueOnSale)

					currencyUnicambio := currencyUnicambios{
						description: row.ChildText("td:nth-child(2)"),
						valueToBuy:  valueToBuy,
						valueOnSale: valueOnSale,
					}

					fmt.Println("Currency scraped: ", currencyUnicambio)

					*currenciesUnicambios = append(*currenciesUnicambios, currencyUnicambio)
				}
			})
		})
	}
}
