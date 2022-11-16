package interactions

import (
	reports "currency-exchange-medellin/reports"
	scraping "currency-exchange-medellin/scraping"
	utils "currency-exchange-medellin/utils"
	"fmt"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

const (
	formTarget                = "#iframe"
	acceptCookiesTarget       = "#gdpr-cookie-accept-all"
	currencyDescriptionTarget = "#select2-_ReservasCalculoPreciosPortlet_demandSelect-container"
	searchTextCurrencyTarget  = "input.select2-search__field"
	rateDemandTarget          = "#rateDemand"
)

type ExchangeGlobocambios struct {
	RequestExchange
	SiteName string
}

var currencyCodeGloboMap = map[string]string{
	"USD": "Dólar USA",
	"EUR": "Euro",
	"GBP": "Libra Esterlina",
	"MXN": "Peso Mexicano",
	"CHF": "Franco Suizo",
	"CNY": "Yuan China",
	"KRW": "Corea, República de - Won Surcoreano",
	"MAD": "Dirham Marroquí",
	"AUD": "Dólar australiano",
	"NZD": "Dólar Neozelandés",
	"INR": "India - Rupia India",
	"ILS": "Israel - Nuevo Séquel",
	"JMD": "Jamaica - Dólar de Jamaica",
	"NIO": "Nicaragua - Córdoba Nicaragua",
	"DOP": "República Dominicana - Peso Dominicano",
	"TTD": "Trinidad y Tobago - Dólar de Trinidad y Tobago",
	"UYU": "Uruguay - Peso Uruguayo",
}

func (reqExchange *ExchangeGlobocambios) selectExchange() ResultExchange {

	fmt.Println("----------- Globocambios Currency Exchange ----------------")

	fmt.Printf("Request Currency Exchange - Globocambios: %s - Value: %f - OperationType: %s \n",
		reqExchange.Exchange.CurrencyCode, reqExchange.Exchange.Value, reqExchange.Exchange.OperationType)

	var scraper = scraping.BuildGoRodScrapper(reqExchange.Url)

	currencyGlobocambios := scrapCurrencyGlobocambios(scraper, reqExchange.Exchange.CurrencyCode)

	reportCurrencies := FromCurrenciesToReportCurrencies([]Currency{currencyGlobocambios})

	reports.ReportCSV(reportPaths["GC"], reportCurrencies)

	var resultExchange = CalculateConversion([]Currency{currencyGlobocambios}, &reqExchange.Exchange)

	fmt.Printf("Result Exchange Unicambios:\nCurrency: %s\nOperation Type: %s\nValue Operation: %f\nValue Convertion: %f\n",
		resultExchange.Exchange.CurrencyCode, resultExchange.OperationType, resultExchange.Value, resultExchange.ValueConvertion)

	return resultExchange
}

func scrapCurrencyGlobocambios(scraper *rod.Page, currencyCode string) Currency {

	scraper.MustElement(acceptCookiesTarget).MustClick()
	scraper.MustElement(formTarget).ScrollIntoView()

	var valueOperation float64 = 0

	scraper.MustElement(currencyDescriptionTarget).MustClick()
	scraper.MustElement(searchTextCurrencyTarget).MustInput(currencyCodeGloboMap[currencyCode])
	scraper.KeyActions().Press(input.Enter).MustDo()

	time.Sleep(6 * time.Second)

	rateDemandStr, _ := scraper.MustElement(rateDemandTarget).Text() //1 USD = 4870,000214 COP
	fmt.Println("rate demand: ", rateDemandStr)
	rateDemandArr := strings.Split(rateDemandStr, " ") //["1","USD","4870,000214","COP"]

	valueOperationStr := rateDemandArr[3] //valueOperation
	valueOperation, _ = utils.FromStringToFloat(strings.Replace(valueOperationStr, ",", ".", -1))

	fmt.Println("value operation: ", valueOperation)

	return Currency{
		ValueToBuy:  valueOperation,
		Description: currencyCode,
	}
}
