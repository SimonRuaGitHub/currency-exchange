package main

import (
	interactions "currency-exchange-medellin/interactions"
	mappingmed "currency-exchange-medellin/webmapping-med"
	"fmt"
)

func main() {

	fmt.Println("------ Currency Conversion Agent in Action -----")

	for name, url := range mappingmed.ExchangeHouses {

		switch name {

		case mappingmed.Globocambios:
			reqExchange := interactions.ExchangeGlobocambios{
				RequestExchange: interactions.RequestExchange{
					Exchange: interactions.Exchange{Currency: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			interactions.SelectCurrencyExchange(&reqExchange)
		case mappingmed.Moneymax:
			reqExchange := interactions.ExchangeMaxmoney{
				RequestExchange: interactions.RequestExchange{
					Exchange: interactions.Exchange{Currency: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			interactions.SelectCurrencyExchange(&reqExchange)

		case mappingmed.Unicambios:
			reqExchange := interactions.ExchangeUnicambios{
				RequestExchange: interactions.RequestExchange{
					Exchange: interactions.Exchange{Currency: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			interactions.SelectCurrencyExchange(&reqExchange)

		case mappingmed.Homecambios:
			reqExchange := interactions.ExchangeHomeCambios{
				RequestExchange: interactions.RequestExchange{
					Exchange: interactions.Exchange{Currency: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			interactions.SelectCurrencyExchange(&reqExchange)

		case mappingmed.Nutifinanzas:
			reqExchange := interactions.ExchangeNutifinanzas{
				RequestExchange: interactions.RequestExchange{
					Exchange: interactions.Exchange{Currency: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			interactions.SelectCurrencyExchange(&reqExchange)

		}
	}
}
