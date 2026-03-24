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
		total = total + convert(m, currency)
	}
	return Money{Amount: total, Currency: currency}
}

// 编写换算货币的convert方法
func convert(money Money, currency string) float64 {
	eurToUsd := 1.2
	if money.Currency == currency {
		return money.Amount
	}
	return money.Amount * eurToUsd
}