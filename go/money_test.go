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
func TestMultiplication(t *testing.T) {
	tenEuros := s.Money{Amount: 10, Currency: "EUR"}
	actualResult := tenEuros.Times(2)

	expectedResult := s.Money{Amount: 20, Currency: "EUR"}
	assertEqual(t, expectedResult, actualResult)
}

// 测试代码2
func TestDivision(t *testing.T) {
	originalMoney := s.Money{Amount: 4002, Currency: "KRW"}
	actualResult := originalMoney.Divide(4)

	expectedResult := s.Money{Amount: 1000.5, Currency: "KRW"}
	assertEqual(t, expectedResult, actualResult)
}

// 测试代码3-测试加法
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

// 测试代码4-测试美元与欧元相加
func TestAdditionOfDollarsAndEuros(t *testing.T) {
	var portfolio s.Portfolio

	fiveDollars := s.NewMoney(5, "USD")
	tenEuros := s.NewMoney(10, "EUR")
	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)

	expectedValue := s.NewMoney(17, "USD")
	actualValue := portfolio.Evaluate("USD")

	assertEqual(t, expectedValue, actualValue)
}

func TestAdditionOfDollarsAndWons(t *testing.T) {
	var portfolio s.Portfolio

	oneDollar := s.NewMoney(1, "USD")
	elevenHundredWon := s.NewMoney(1100, "KRW")

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(elevenHundredWon)

	expectedValue := s.NewMoney(2200, "KRW")
	actualValue := portfolio.Evaluate("KRW")

	assertEqual(t, expectedValue, actualValue)
}
