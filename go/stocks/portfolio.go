package stocks

// 实现Portfolio类型
type Portfolio []Money

func (p Portfolio) Add(money Money) Portfolio {
	p = append(p, money)
	return p
}

func (p Portfolio) Evaluate(currency string) Money {
	total := 0.0
	for _, m := range p {
		total = total + m.Amount
	}
	return Money{Amount: total, Currency: currency}
}