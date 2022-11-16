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
		{"Description", "ValueOnSale", "ValueToBuy", "DateTimeReport"},
	}

	for _, re := range reportCurrencies {
		rowReportCurr := []string{re.Description, re.ValueOnSale, re.ValueToBuy, re.DateTime.Format("2006-01-02T15:04:05-0700")}
		currencyTable = append(currencyTable, rowReportCurr)
	}

	errWritingAll := writer.WriteAll(currencyTable)

	if errWritingAll != nil {
		log.Fatalf("Could not write inside file, err: %q", errWritingAll)
		return
	}
}
