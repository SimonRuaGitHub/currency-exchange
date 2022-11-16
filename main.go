package main

import (
	agent "currency-exchange-medellin/agentexchange"
	exchangemed "currency-exchange-medellin/exchangemed"
	"fmt"
	"sync"
)

func main() {

	fmt.Println("------ Currency Conversion Agent in Action -----")

	var wg sync.WaitGroup
	resultExchanges := make([]agent.ResultExchange, 0)

	for name, url := range exchangemed.ExchangeHouses {

		switch name {

		case exchangemed.Globocambios:
			reqExchange := agent.ExchangeGlobocambios{
				RequestExchange: agent.RequestExchange{
					Exchange: agent.Exchange{CurrencyCode: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
					Name:     name,
				},
				SiteName: "Aeropuerto Internacional José María Córdova",
			}

			wg.Add(1)

			go func() {
				defer wg.Done()
				resultExchange := agent.SelectCurrencyExchange(&reqExchange)
				resultExchanges = append(resultExchanges, resultExchange)
			}()

		case exchangemed.Moneymax:
			reqExchange := agent.ExchangeMaxmoney{
				RequestExchange: agent.RequestExchange{
					Exchange: agent.Exchange{CurrencyCode: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
					Name:     name,
				},
			}

			wg.Add(1)

			go func() {
				defer wg.Done()
				resultExchange := agent.SelectCurrencyExchange(&reqExchange)
				resultExchanges = append(resultExchanges, resultExchange)
			}()

		case exchangemed.Unicambios:
			reqExchange := agent.ExchangeUnicambios{
				RequestExchange: agent.RequestExchange{
					Exchange: agent.Exchange{CurrencyCode: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
					Name:     name,
				},
			}

			wg.Add(1)

			go func() {
				defer wg.Done()
				resultExchange := agent.SelectCurrencyExchange(&reqExchange)
				resultExchanges = append(resultExchanges, resultExchange)
			}()

		case exchangemed.Homecambios:
			reqExchange := agent.ExchangeHomeCambios{
				RequestExchange: agent.RequestExchange{
					Exchange: agent.Exchange{CurrencyCode: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
					Name:     name,
				},
			}

			wg.Add(1)

			go func() {
				defer wg.Done()
				resultExchange := agent.SelectCurrencyExchange(&reqExchange)
				resultExchanges = append(resultExchanges, resultExchange)
			}()

		case exchangemed.Nutifinanzas:
			reqExchange := agent.ExchangeNutifinanzas{
				RequestExchange: agent.RequestExchange{
					Exchange: agent.Exchange{CurrencyCode: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
					Name:     name,
				},
			}

			wg.Add(1)

			go func() {
				defer wg.Done()
				resultExchange := agent.SelectCurrencyExchange(&reqExchange)
				resultExchanges = append(resultExchanges, resultExchange)
			}()
		}
	}

	wg.Wait()

	fmt.Println("------ The best exchange houses offer -----")

	resultExchange := bestPrice(resultExchanges)

	fmt.Printf("Exchange house: %s:\nCurrency: %s\nOperation Type: %s\nValue Operation: %f\nValue Convertion: %f\n",
		resultExchange.Name, resultExchange.Exchange.CurrencyCode, resultExchange.OperationType, resultExchange.Value, resultExchange.ValueConvertion)
}

func bestPrice(resultExchanges []agent.ResultExchange) agent.ResultExchange {

	minPrice := resultExchanges[0].Exchange.Value
	j := 0

	for i := 1; i < len(resultExchanges); i++ {

		if minPrice > resultExchanges[i].Exchange.Value {
			minPrice = resultExchanges[i].Exchange.Value
			j = i
		}
	}

	return resultExchanges[j]
}
