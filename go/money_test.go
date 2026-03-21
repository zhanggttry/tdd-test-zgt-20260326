package main

import (
	"testing"
)

// 提取的函数
func assertEqual(t *testing.T, expected Money, actual Money) {
	if expected != actual {
		t.Errorf("期望 %+v, 实际 %+v", expected, actual)
	}
}

// 测试代码1
func TestMultiplicationInDollars(t *testing.T) {
	fiver := Money{amount: 5, currency: "USD"}
	actualResult:= fiver.Times(2)
	expectedResult := Money{amount: 10, currency: "USD"}
	assertEqual(t, expectedResult, actualResult)
}

// 测试代码2
func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := Money{amount: 10, currency: "EUR"}
	actualResult := tenEuros.Times(2)

	expectedResult := Money{amount: 20, currency: "EUR"}
	assertEqual(t, expectedResult, actualResult)
}

// 测试代码3
func TestDivision(t *testing.T) {
	originalMoney := Money{amount: 4002, currency: "KRW"}
	actualResult := originalMoney.Divide(4)

	expectedResult := Money{amount: 1000.5, currency: "KRW"}
	assertEqual(t, expectedResult, actualResult)
}

// 测试代码5-测试加法
func TestAddition(t *testing.T) {
	var portfolio Portfolio
	var portfolioInDollars Money
	
	fiveDollars := Money{amount: 5, currency: "USD"}
	tenDollars := Money{amount: 10, currency: "USD"}
	fifteenDollars := Money{amount: 15, currency: "USD"}

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars = portfolio.Evaluate("USD")

	assertEqual(t, fifteenDollars, portfolioInDollars)
}

