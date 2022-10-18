package interactions

type ExchangeMaxmoney struct {
	RequestExchange
}

func (reqExchange *ExchangeMaxmoney) selectExchange() ResultExchange {
	return ResultExchange{
		Exchange{"USD", 4135.43, "purshace"},
	}
}
