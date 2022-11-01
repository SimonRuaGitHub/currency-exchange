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
	description string
	valueOnSale float64
	valueToBuy  float64
}

func SelectCurrencyExchange(agent AgentCurrencyExchange) ResultExchange {
	return agent.selectExchange()
}

func CalculateConversion(currenciesInfo []Currency, reqExchange *Exchange) ResultExchange {
	var valueConvertion float64 = 0.0
	var valueOperation float64 = 0.0

	for _, currencyInfo := range currenciesInfo {
		if strings.Contains(currencyInfo.description, reqExchange.CurrencyCode) {
			fmt.Println("Found currency: ", currencyInfo.description)

			if reqExchange.OperationType == "purchase" {
				valueConvertion = currencyInfo.valueToBuy * reqExchange.Value
				valueOperation = currencyInfo.valueToBuy
			} else {
				valueConvertion = currencyInfo.valueOnSale * reqExchange.Value
				valueOperation = currencyInfo.valueOnSale
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

			if currency.description == description {
				newCurrency := Currency{
					description: code,
					valueOnSale: currency.valueOnSale,
					valueToBuy:  currency.valueToBuy,
				}

				newCurrencies = append(newCurrencies, newCurrency)

				break
			}
		}
	}

	return newCurrencies
}
