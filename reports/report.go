package reports

import (
	"encoding/csv"
	"log"
	"os"
)

type ReportCurrency struct {
	valueToBuy  string
	valueOnSale string
	description string
}

func reportCSV(filePath string, reportExchanges []ReportCurrency) {
	file, errCreateFile := os.Create(filePath)
	if errCreateFile != nil {
		log.Fatalf("Could not create file, err: %q", errCreateFile)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var currencyTable = [][]string{
		{"Description", "ValueOnSale", "ValueToBuy"},
	}

	for _, re := range reportExchanges {
		rowReportExc := []string{re.description, re.valueOnSale, re.valueToBuy}
		currencyTable = append(currencyTable, rowReportExc)
	}

	errWritingAll := writer.WriteAll(currencyTable)

	if errWritingAll != nil {
		log.Fatalf("Could not write inside file, err: %q", errWritingAll)
		return
	}
}
