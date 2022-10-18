package interactions

import (
	utils "currency-exchange-medellin/utils"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

const scrappingTimeout = 120 * time.Second

const (
	currenciesRows    = "tr"
	currenciesColumns = "td"
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
	fmt.Printf("Request Currency: %s - Value: %f - OperationType: %s \n",
		reqExchange.Exchange.Currency, reqExchange.Exchange.Value, reqExchange.Exchange.OperationType)

	var scrapper = buildScrapper(scrappingTimeout)

	currenciesUnicambios := make([]currencyUnicambios, 0)

	scrapCurrenciesInfo(scrapper, &currenciesUnicambios)

	scrapper.Visit(reqExchange.Url)

	var resultExchange = calculateConversion(currenciesUnicambios, reqExchange)

	fmt.Println("Result Exchange Unicambios: ", resultExchange)

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

func buildScrapper(timeout time.Duration) *colly.Collector {
	scrapper := colly.NewCollector()

	scrapper.SetRequestTimeout(timeout)

	scrapper.OnRequest(func(request *colly.Request) {
		fmt.Println("Visting web page: " + request.URL.String())
	})

	scrapper.OnResponse(func(r *colly.Response) {
		fmt.Println("Got response from page")
	})

	scrapper.OnError(func(r *colly.Response, e error) {
		fmt.Println("Error when visiting page: ", e)
	})

	scrapper.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished scrapping page")
	})

	return scrapper
}

func scrapCurrenciesInfo(scrapper *colly.Collector, currenciesUnicambios *[]currencyUnicambios) {

	for tableSide, table := range currenciesTables {

		fmt.Println("Scrapping currencies from table: ", tableSide)

		scrapper.OnHTML(table, func(tableHtml *colly.HTMLElement) {
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

					fmt.Println("Currency scrapped: ", currencyUnicambio)

					*currenciesUnicambios = append(*currenciesUnicambios, currencyUnicambio)
				}
			})
		})
	}
}
