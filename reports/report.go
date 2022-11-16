package reports

import (
	"encoding/csv"
	"log"
	"os"
	"time"
)

type ReportCurrency struct {
	ValueToBuy  string
	ValueOnSale string
	Description string
	DateTime    time.Time
}

type ReportBestOffer struct {
	Name            string
	CurrencyCode    string
	ValueOperation  string
	OperationType   string
	ValueConversion string
	DateTime        time.Time
}

func ReportCSV(filePath string, reportCurrencies []ReportCurrency) {

	file, errCreateFile := os.Create(filePath)
	if errCreateFile != nil {
		log.Fatalf("Could not create file, err: %q", errCreateFile)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var currencyTable = [][]string{
		{"Description", "OnSalePrice", "PurchasePrice", "DateTimeReport"},
	}

	for _, re := range reportCurrencies {
		rowReportCurr := []string{re.Description, re.ValueToBuy, re.ValueOnSale, re.DateTime.Format("2006-01-02T15:04:05-0700")}
		currencyTable = append(currencyTable, rowReportCurr)
	}

	errWritingAll := writer.WriteAll(currencyTable)

	if errWritingAll != nil {
		log.Fatalf("Could not write inside file, err: %q", errWritingAll)
		return
	}
}

func ReportCSVBestOffer(filePath string, rbestOffer ReportBestOffer) {
	file, errCreateFile := os.Create(filePath)
	if errCreateFile != nil {
		log.Fatalf("Could not create file, err: %q", errCreateFile)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var bestOfferTable = [][]string{
		{"ExchangeHouseName", "CurrencyCode", "OperationType", "ValueOperation", "ValueConversion", "DateTime"},
	}

	rowReportBestOffer := []string{rbestOffer.Name, rbestOffer.CurrencyCode, rbestOffer.OperationType, rbestOffer.ValueOperation, rbestOffer.ValueConversion, rbestOffer.DateTime.Format("2006-01-02T15:04:05-0700")}
	bestOfferTable = append(bestOfferTable, rowReportBestOffer)

	errWritingAll := writer.WriteAll(bestOfferTable)

	if errWritingAll != nil {
		log.Fatalf("Could not write inside file, err: %q", errWritingAll)
		return
	}
}
