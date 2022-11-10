package reports

import (
	agent "currency-exchange-medellin/agentexchange"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type ReportCurrency struct {
	valueToBuy  string
	valueOnSale string
	description string
	dateTime    time.Time
}

func mapToReportCurrency(currency agent.Currency) ReportCurrency {

	valueOnSaleStr := fmt.Sprintf("%f", currency.ValueOnSale)
	valueToBuyStr := fmt.Sprintf("%f", currency.ValueToBuy)

	return ReportCurrency{
		valueToBuy:  valueToBuyStr,
		valueOnSale: valueOnSaleStr,
		description: currency.Description,
		dateTime:    time.Now(),
	}
}

func fromCurrenciesToReportCurrencies(currencies []agent.Currency) []ReportCurrency {
	reportCurrencies := make([]ReportCurrency, 0)

	for _, currency := range currencies {
		reportCurrency := mapToReportCurrency(currency)
		reportCurrencies = append(reportCurrencies, reportCurrency)
	}

	return reportCurrencies
}

func reportCSV(filePath string, currencies []agent.Currency) {

	reportCurrencies := fromCurrenciesToReportCurrencies(currencies)

	file, errCreateFile := os.Create(filePath)
	if errCreateFile != nil {
		log.Fatalf("Could not create file, err: %q", errCreateFile)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var currencyTable = [][]string{
		{"Description", "ValueOnSale", "ValueToBuy", "Date time Report"},
	}

	for _, re := range reportCurrencies {
		rowReportCurr := []string{re.description, re.valueOnSale, re.valueToBuy, re.dateTime.GoString()}
		currencyTable = append(currencyTable, rowReportCurr)
	}

	errWritingAll := writer.WriteAll(currencyTable)

	if errWritingAll != nil {
		log.Fatalf("Could not write inside file, err: %q", errWritingAll)
		return
	}
}
