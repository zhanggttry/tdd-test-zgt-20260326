package stocks

type Money struct {
	Amount   float64
	Currency string
}

func (m Money) Times(multiplier int) Money {
	return Money{Amount: m.Amount * float64(multiplier), Currency: m.Currency}
}

func (m Money) Divide(divisor int) Money {
	return Money{Amount: m.Amount / float64(divisor), Currency: m.Currency}
}

func NewMoney(amount float64, currency string) Money {
	return Money{Amount: amount, Currency: currency}
}