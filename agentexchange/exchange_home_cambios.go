package interactions

type ExchangeHomeCambios struct {
	RequestExchange
}

func (reqExchange *ExchangeHomeCambios) selectExchange() ResultExchange {
	return ResultExchange{
		Exchange{"USD", 4135.43, "purshace"},
		0.0,
	}
}
