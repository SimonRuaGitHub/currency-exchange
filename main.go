package main

import (
	interactions "currency-exchange-medellin/interactions"
	webmapping "currency-exchange-medellin/webmapping-med"
	"fmt"
)

func main() {

	fmt.Println("------ Currency conversion agent -----")

	for name, url := range webmapping.ExchangeHouses {

		switch name {

		case webmapping.Globocambios:
			reqExchange := interactions.ExchangeGlobocambios{
				RequestExchange: interactions.RequestExchange{
					Exchange: interactions.Exchange{Currency: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			interactions.SelectCurrencyExchange(&reqExchange)
		case webmapping.Moneymax:
			reqExchange := interactions.ExchangeMaxmoney{
				RequestExchange: interactions.RequestExchange{
					Exchange: interactions.Exchange{Currency: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			interactions.SelectCurrencyExchange(&reqExchange)

		case webmapping.Unicambios:
			reqExchange := interactions.ExchangeUnicambios{
				RequestExchange: interactions.RequestExchange{
					Exchange: interactions.Exchange{Currency: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			interactions.SelectCurrencyExchange(&reqExchange)

		case webmapping.Homecambios:
			reqExchange := interactions.ExchangeHomeCambios{
				RequestExchange: interactions.RequestExchange{
					Exchange: interactions.Exchange{Currency: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			interactions.SelectCurrencyExchange(&reqExchange)

		case webmapping.Nutifinanzas:
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
