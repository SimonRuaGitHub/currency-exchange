package interactions

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
	Currency      string
	Value         float64
	OperationType string
}

func SelectCurrencyExchange(agent AgentCurrencyExchange) ResultExchange {
	return agent.selectExchange()
}
