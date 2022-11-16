package interactions

import (
	reports "currency-exchange-medellin/reports"
	"fmt"
	"time"
)

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

func FromCurrenciesToReportCurrencies(currencies []Currency) []reports.ReportCurrency {
	reportCurrencies := make([]reports.ReportCurrency, 0)

	for _, currency := range currencies {
		reportCurrency := mapToReportCurrency(currency)
		reportCurrencies = append(reportCurrencies, reportCurrency)
	}

	return reportCurrencies
}
