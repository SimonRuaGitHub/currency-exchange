package interactions

import (
	reports "currency-exchange-medellin/reports"
	"fmt"
	"time"
)

const rootFolder = "reports/"

var ReportPaths = map[string]string{
	"GC":        rootFolder + "globocambios.csv",
	"HC":        rootFolder + "homecambios.csv",
	"MM":        rootFolder + "moneymax.csv",
	"NF":        rootFolder + "nutifinanzas.csv",
	"UC":        rootFolder + "unicambios.csv",
	"BESTOFFER": rootFolder + "bestoffer.csv",
}

func mapToReportCurrency(currency Currency) reports.ReportCurrency {

	valueOnSaleStr := fmt.Sprintf("%f", currency.ValueOnSale)
	valueToBuyStr := fmt.Sprintf("%f", currency.ValueToBuy)

	return reports.ReportCurrency{
		ValueToBuy:  valueToBuyStr,
		ValueOnSale: valueOnSaleStr,
		Description: currency.Description,
		DateTime:    time.Now(),
	}
}

func MapToReportBestOffer(resultExchange ResultExchange) reports.ReportBestOffer {

	valueOperationStr := fmt.Sprintf("%f", resultExchange.Exchange.Value)
	valueConvertionStr := fmt.Sprintf("%f", resultExchange.ValueConvertion)

	return reports.ReportBestOffer{
		Name:            resultExchange.Name,
		CurrencyCode:    resultExchange.Exchange.CurrencyCode,
		OperationType:   resultExchange.Exchange.OperationType,
		ValueOperation:  valueOperationStr,
		ValueConversion: valueConvertionStr,
		DateTime:        time.Now(),
	}
}

func FromCurrenciesToReportCurrencies(currencies []Currency) []reports.ReportCurrency {
	reportCurrencies := make([]reports.ReportCurrency, 0)

	for _, currency := range currencies {
		reportCurrency := mapToReportCurrency(currency)
		reportCurrencies = append(reportCurrencies, reportCurrency)
	}

	return reportCurrencies
}
