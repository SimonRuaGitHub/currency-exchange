package main

import (
	exchangemed "currency-exchange-medellin/exchange-med"
	interactions "currency-exchange-medellin/interactions"
	"fmt"
)

func main() {

	fmt.Println("------ Currency Conversion Agent in Action -----")

	for name, url := range exchangemed.ExchangeHouses {

		switch name {

		case exchangemed.Globocambios:
			reqExchange := interactions.ExchangeGlobocambios{
				RequestExchange: interactions.RequestExchange{
					Exchange: interactions.Exchange{Currency: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			interactions.SelectCurrencyExchange(&reqExchange)
		case exchangemed.Moneymax:
			reqExchange := interactions.ExchangeMaxmoney{
				RequestExchange: interactions.RequestExchange{
					Exchange: interactions.Exchange{Currency: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			interactions.SelectCurrencyExchange(&reqExchange)

		case exchangemed.Unicambios:
			reqExchange := interactions.ExchangeUnicambios{
				RequestExchange: interactions.RequestExchange{
					Exchange: interactions.Exchange{Currency: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			interactions.SelectCurrencyExchange(&reqExchange)

		case exchangemed.Homecambios:
			reqExchange := interactions.ExchangeHomeCambios{
				RequestExchange: interactions.RequestExchange{
					Exchange: interactions.Exchange{Currency: "USD", Value: 242, OperationType: "purchase"},
					Url:      url,
				},
			}

			interactions.SelectCurrencyExchange(&reqExchange)

		case exchangemed.Nutifinanzas:
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
