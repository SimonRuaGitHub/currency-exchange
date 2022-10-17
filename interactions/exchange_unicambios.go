package interactions

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

const (
	currenciesTable   = "section.premium-tabs-section table#supsystic-table-12 tbody"
	currenciesRows    = "tr"
	currenciesColumns = "td"
)

type ExchangeUnicambios struct {
	RequestExchange
}

type currencyUnicambios struct {
	description string
	valueOnSale float64
	valueToBuy  float64
}

func (reqExchange *ExchangeUnicambios) selectExchange() ResultExchange {
	fmt.Printf("Request Currency: %s - Value: %f - OperationType: %s",
		reqExchange.Exchange.Currency, reqExchange.Exchange.Value, reqExchange.Exchange.OperationType)

	scrapper := colly.NewCollector()

	scrapper.SetRequestTimeout(120 * time.Second)

	scrapper.OnRequest(func(request *colly.Request) {
		fmt.Println("Visting web page: " + request.URL.String())
	})

	scrapper.OnResponse(func(r *colly.Response) {
		fmt.Println("Got response from page")
	})

	currenciesUnicambios := make([]currencyUnicambios, 0)

	scrapper.OnHTML(currenciesTable, func(tableHtml *colly.HTMLElement) {
		tableHtml.ForEach(currenciesRows, func(i int, row *colly.HTMLElement) {
			if row.ChildText("td:nth-child(2)") != "" {

				valueToBuy, _ := fromStringToFloat("td:nth-child(3)")
				valueOnSale, _ := fromStringToFloat("td:nth-child(4)")

				currencyUnicambio := currencyUnicambios{
					description: row.ChildText("td:nth-child(2)"),
					valueToBuy:  valueToBuy,
					valueOnSale: valueOnSale,
				}

				fmt.Println("Currency scrapped: ", currencyUnicambio)

				currenciesUnicambios = append(currenciesUnicambios, currencyUnicambio)
			}
		})
	})

	scrapper.OnError(func(r *colly.Response, e error) {
		fmt.Println("Error when visiting page: ", e)
	})

	scrapper.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished scrapping page")
	})

	scrapper.Visit(reqExchange.Url)

	return ResultExchange{
		Exchange{"USD", 4135.43, "purshace"},
	}
}

func fromStringToFloat(value string) (parsedValue float64, err error) {

	floatValue, err := strconv.ParseFloat(value, 64)

	if err != nil {
		fmt.Errorf("Error when parsing from string to float64", err)
		return 0.0, err
	} else {
		return floatValue, nil
	}
}
