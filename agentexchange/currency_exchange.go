package interactions

import (
	"fmt"
	"strings"
)

type AgentCurrencyExchange interface {
	selectExchange() ResultExchange
}

type RequestExchange struct {
	Exchange
	Url string
}

type ResultExchange struct {
	Exchange
	ValueConvertion float64
}

type Exchange struct {
	CurrencyCode  string
	Value         float64
	OperationType string
}

type Currency struct {
	Description string
	ValueOnSale float64
	ValueToBuy  float64
}

func SelectCurrencyExchange(agent AgentCurrencyExchange) ResultExchange {
	return agent.selectExchange()
}

func CalculateConversion(currenciesInfo []Currency, reqExchange *Exchange) ResultExchange {
	var valueConvertion float64 = 0.0
	var valueOperation float64 = 0.0

	for _, currencyInfo := range currenciesInfo {
		if strings.Contains(currencyInfo.Description, reqExchange.CurrencyCode) {
			fmt.Println("Found currency: ", currencyInfo.Description)

			if reqExchange.OperationType == "purchase" {
				valueConvertion = currencyInfo.ValueToBuy * reqExchange.Value
				valueOperation = currencyInfo.ValueToBuy
			} else {
				valueConvertion = currencyInfo.ValueOnSale * reqExchange.Value
				valueOperation = currencyInfo.ValueOnSale
			}
			break
		}
	}

	return ResultExchange{
		Exchange{reqExchange.CurrencyCode, valueOperation, reqExchange.OperationType},
		valueConvertion,
	}
}

func HomologateCurrencyDespByCode(currencies []Currency, currencyCodeMap map[string]string) []Currency {

	newCurrencies := make([]Currency, 0)

	for code, description := range currencyCodeMap {
		for _, currency := range currencies {

			if currency.Description == description {
				newCurrency := Currency{
					Description: code,
					ValueOnSale: currency.ValueOnSale,
					ValueToBuy:  currency.ValueToBuy,
				}

				newCurrencies = append(newCurrencies, newCurrency)

				break
			}
		}
	}

	return newCurrencies
}
