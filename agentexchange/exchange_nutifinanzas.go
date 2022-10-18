package interactions

type ExchangeNutifinanzas struct {
	RequestExchange
}

func (reqExchange *ExchangeNutifinanzas) selectExchange() ResultExchange {
	return ResultExchange{
		Exchange{"USD", 4135.43, "purshace"},
		0.0,
	}
}
