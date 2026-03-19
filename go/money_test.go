package main

import (
	"testing"
)

// 测试代码1
func TestMultiplicationInDollars(t *testing.T) {
	fiver := Money{amount: 5, currency: "USD"}
	actualResult:= fiver.Times(2)
	expectedResult := Money{amount: 10, currency: "USD"}
	if expectedResult != actualResult {
		t.Errorf("期望 [%+v]， 实际： [%+v]", expectedResult, actualResult)
	}
}

// 测试代码2
func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := Money{amount: 10, currency: "EUR"}
	twentyEuros := tenEuros.Times(2)
	if twentyEuros.amount != 20 {
		t.Errorf("期望 20， 实际： [%f]", twentyEuros.amount)
	}
	if twentyEuros.currency != "EUR" {
		t.Errorf("期望 EUR, 实际： [%s]", twentyEuros.currency)
	}
}

// 测试代码3
func TestDivision(t *testing.T) {
	originalMoney := Money{amount: 4002, currency: "KRW"}
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := Money{amount: 1000.5, currency: "KRW"}

	if expectedMoneyAfterDivision != actualMoneyAfterDivision {
		t.Errorf("期望 %+v，实际 %+v",
	        expectedMoneyAfterDivision, actualMoneyAfterDivision)
	}
}

type Money struct {
	amount float64
	currency string
}

func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * float64(multiplier), currency: m.currency}
}

func (m Money) Divide(divisor int) Money {
	return Money{amount: m.amount / float64(divisor), currency: m.currency}
}