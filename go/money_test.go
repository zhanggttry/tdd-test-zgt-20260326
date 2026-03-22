package main

import (
	"testing"
	s "tdd/stocks"
)

// 提取的函数
func assertEqual(t *testing.T, expected s.Money, actual s.Money) {
	if expected != actual {
		t.Errorf("期望 %+v, 实际 %+v", expected, actual)
	}
}

// 测试代码1
func TestMultiplicationInDollars(t *testing.T) {
	fiver := s.Money{Amount: 5, Currency: "USD"}
	actualResult:= fiver.Times(2)
	expectedResult := s.Money{Amount: 10, Currency: "USD"}
	assertEqual(t, expectedResult, actualResult)
}

// 测试代码2
func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := s.Money{Amount: 10, Currency: "EUR"}
	actualResult := tenEuros.Times(2)

	expectedResult := s.Money{Amount: 20, Currency: "EUR"}
	assertEqual(t, expectedResult, actualResult)
}

// 测试代码3
func TestDivision(t *testing.T) {
	originalMoney := s.Money{Amount: 4002, Currency: "KRW"}
	actualResult := originalMoney.Divide(4)

	expectedResult := s.Money{Amount: 1000.5, Currency: "KRW"}
	assertEqual(t, expectedResult, actualResult)
}

// 测试代码 5-测试加法
func TestAddition(t *testing.T) {
	var portfolio s.Portfolio
	var portfolioInDollars s.Money
	
	fiveDollars := s.Money{Amount: 5, Currency: "USD"}
	tenDollars := s.Money{Amount: 10, Currency: "USD"}
	fifteenDollars := s.Money{Amount: 15, Currency: "USD"}

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars = portfolio.Evaluate("USD")

	assertEqual(t, fifteenDollars, portfolioInDollars)
}

