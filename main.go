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

	for name, url := range exchangemed.ExchangeHouses {

		switch name {

		case exchangemed.Globocambios:
			reqExchange := agent.ExchangeGlobocambios{
				RequestExchange: agent.RequestExchange{
					Exchange: agent.Exchange{CurrencyCode: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
				SiteName: "Aeropuerto Internacional José María Córdova",
			}

			wg.Add(1)

			go func() {
				defer wg.Done()
				agent.SelectCurrencyExchange(&reqExchange)
			}()

		case exchangemed.Moneymax:
			reqExchange := agent.ExchangeMaxmoney{
				RequestExchange: agent.RequestExchange{
					Exchange: agent.Exchange{CurrencyCode: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			wg.Add(1)

			go func() {
				defer wg.Done()
				agent.SelectCurrencyExchange(&reqExchange)
			}()

		case exchangemed.Unicambios:
			reqExchange := agent.ExchangeUnicambios{
				RequestExchange: agent.RequestExchange{
					Exchange: agent.Exchange{CurrencyCode: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			wg.Add(1)

			go func() {
				defer wg.Done()
				agent.SelectCurrencyExchange(&reqExchange)
			}()

		case exchangemed.Homecambios:
			reqExchange := agent.ExchangeHomeCambios{
				RequestExchange: agent.RequestExchange{
					Exchange: agent.Exchange{CurrencyCode: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			wg.Add(1)

			go func() {
				defer wg.Done()
				agent.SelectCurrencyExchange(&reqExchange)
			}()

		case exchangemed.Nutifinanzas:
			reqExchange := agent.ExchangeNutifinanzas{
				RequestExchange: agent.RequestExchange{
					Exchange: agent.Exchange{CurrencyCode: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			wg.Add(1)

			go func() {
				defer wg.Done()
				agent.SelectCurrencyExchange(&reqExchange)
			}()

		}
	}

	wg.Wait()

}
