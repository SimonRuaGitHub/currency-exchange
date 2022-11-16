package interactions

import (
	reports "currency-exchange-medellin/reports"
	scraping "currency-exchange-medellin/scraping"
	utils "currency-exchange-medellin/utils"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

const (
	containerTarget   = "#SITE_PAGES_TRANSITION_GROUP"
	purchaseUSDtarget = "#comp-l721a1km > h2.font_2:nth-child(1) > span"
	onSaleUSDtarget   = "#comp-l5r49t4d > h2 > span > span > span > span > span"
	purchaseEURtarget = "#comp-l9t12o5w > h2 > span > span > span > span > span"
	onSaleEURtarget   = "#comp-l9t12o5w > h2 > span > span > span > span > span"
	purchaseMXNtarget = "#comp-ku74zp7c > h2 > span > span > span > span > span"
	onSaleMXNtarget   = "#comp-ku7550f3 > h2 > span > span > span > span > span"
)

type ExchangeHomeCambios struct {
	RequestExchange
}

func (reqExchange *ExchangeHomeCambios) selectExchange() ResultExchange {

	fmt.Println("----------- Home Cambios Currency Exchange ----------------")

	fmt.Printf("Request Currency Exchange - Homecambios: %s - Value: %f - OperationType: %s \n",
		reqExchange.Exchange.CurrencyCode, reqExchange.Exchange.Value, reqExchange.Exchange.OperationType)

	var scraper = scraping.BuildCollyScrapper(scraping.DefaultTimeOutColly)

	currenciesHc := make([]Currency, 0)

	scrapCurrenciesHomecambios(scraper, &currenciesHc)

	scraper.Visit(reqExchange.Url)

	reportCurrencies := FromCurrenciesToReportCurrencies(currenciesHc)

	reports.ReportCSV(ReportPaths["HC"], reportCurrencies)

	var resultExchange = CalculateConversion(currenciesHc, &reqExchange.RequestExchange)

	fmt.Printf("Result Exchange HomeCambios:\nCurrency: %s\nOperation Type: %s\nValue Operation: %f\nValue Convertion: %f\n",
		resultExchange.Exchange.CurrencyCode, resultExchange.OperationType, resultExchange.Value, resultExchange.ValueConvertion)

	return resultExchange
}

func scrapCurrenciesHomecambios(scraper *colly.Collector, currenciesHc *[]Currency) {

	scraper.OnHTML(containerTarget, func(container *colly.HTMLElement) {

		purchaseUSDstr := container.ChildText(purchaseUSDtarget)
		onSaleUSDstr := container.ChildText(onSaleUSDtarget)
		purchaseEURstr := container.ChildText(purchaseEURtarget)
		onSaleEURstr := container.ChildText(onSaleEURtarget)
		purchaseMXNstr := container.ChildText(purchaseMXNtarget)
		onSaleMXNstr := container.ChildText(onSaleMXNtarget)

		purchaseUSD, _ := utils.FromStringToFloat(strings.Trim(purchaseUSDstr, " "))
		onSaleUSD, _ := utils.FromStringToFloat(strings.Trim(onSaleUSDstr, " "))

		currencyUSD := Currency{
			Description: "USD",
			ValueToBuy:  onSaleUSD,
			ValueOnSale: purchaseUSD,
		}

		*currenciesHc = append(*currenciesHc, currencyUSD)

		purchaseEUR, _ := utils.FromStringToFloat(strings.Trim(purchaseEURstr, " "))
		onSaleEUR, _ := utils.FromStringToFloat(strings.Trim(onSaleEURstr, " "))

		currencyEUR := Currency{
			Description: "EUR",
			ValueToBuy:  onSaleEUR,
			ValueOnSale: purchaseEUR,
		}

		*currenciesHc = append(*currenciesHc, currencyEUR)

		purchaseMXN, _ := utils.FromStringToFloat(strings.Trim(purchaseMXNstr, " "))
		onSaleMXN, _ := utils.FromStringToFloat(strings.Trim(onSaleMXNstr, " "))

		currencyMXN := Currency{
			Description: "MXN",
			ValueToBuy:  onSaleMXN,
			ValueOnSale: purchaseMXN,
		}

		*currenciesHc = append(*currenciesHc, currencyMXN)

		for _, currency := range *currenciesHc {
			fmt.Println("Currency scraped: ", currency)
		}
	})
}
