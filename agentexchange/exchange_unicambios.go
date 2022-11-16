package interactions

import (
	reports "currency-exchange-medellin/reports"
	scraping "currency-exchange-medellin/scraping"
	utils "currency-exchange-medellin/utils"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

const currenciesRows = "tr"

var currenciesTables = map[string]string{
	"tableLeftSide":  "section.premium-tabs-section table#supsystic-table-12 tbody",
	"tableRightSide": "section.premium-tabs-section table#supsystic-table-11 tbody",
}

type ExchangeUnicambios struct {
	RequestExchange
}

func (reqExchange *ExchangeUnicambios) selectExchange() ResultExchange {

	fmt.Println("----------- Unicambios Currency Exchange ----------------")

	fmt.Printf("Request Currency: %s - Value: %f - OperationType: %s \n",
		reqExchange.Exchange.CurrencyCode, reqExchange.Exchange.Value, reqExchange.Exchange.OperationType)

	var scraper = scraping.BuildCollyScrapper(scraping.DefaultTimeOutColly)

	currenciesUnicambios := make([]Currency, 0)

	scrapCurrenciesUnicambios(scraper, &currenciesUnicambios)

	scraper.Visit(reqExchange.Url)

	reportCurrencies := FromCurrenciesToReportCurrencies(currenciesUnicambios)

	reports.ReportCSV(reportPaths["MM"], reportCurrencies)

	var resultExchange = CalculateConversion(currenciesUnicambios, &reqExchange.Exchange)

	fmt.Printf("Result Exchange Unicambios:\nCurrency: %s\nOperation Type: %s\nValue Operation: %f\nValue Convertion: %f\n",
		resultExchange.Exchange.CurrencyCode, resultExchange.OperationType, resultExchange.Value, resultExchange.ValueConvertion)

	return resultExchange
}

func scrapCurrenciesUnicambios(scraper *colly.Collector, currenciesUnicambios *[]Currency) {

	for tableSide, table := range currenciesTables {

		fmt.Println("Scraping currencies from table: ", tableSide)

		scraper.OnHTML(table, func(tableHtml *colly.HTMLElement) {
			fmt.Println("Table: ", table)
			tableHtml.ForEach(currenciesRows, func(i int, row *colly.HTMLElement) {
				if row.ChildText("td:nth-child(2)") != "" {

					targetValueToBuy := strings.Replace(row.ChildText("td:nth-child(3)"), ".", "", -1)
					targetValueOnSale := strings.Replace(row.ChildText("td:nth-child(4)"), ".", "", -1)
					purchaseValue, _ := utils.FromStringToFloat(targetValueToBuy)
					onSaleValue, _ := utils.FromStringToFloat(targetValueOnSale)

					currencyUnicambio := Currency{
						Description: row.ChildText("td:nth-child(2)"),
						ValueToBuy:  onSaleValue,
						ValueOnSale: purchaseValue,
					}

					fmt.Println("Currency scraped: ", currencyUnicambio)

					*currenciesUnicambios = append(*currenciesUnicambios, currencyUnicambio)
				}
			})
		})
	}
}
