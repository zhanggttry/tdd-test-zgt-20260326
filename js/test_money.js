// 编写第1个测试用例
const assert = require('assert');
const Money = require('./money');
const Portfolio = require('./portfolio');

// 测试代码1
let fiveDollars = new Money(5, "USD");
let tenDollars = new Money(10, "USD");
assert.deepStrictEqual(fiveDollars.times(2), tenDollars);

// 测试代码2
let tenEuros = new Money(10, "EUR");
let twentyEuros = new Money(20, "EUR");
assert.strictEqual(twentyEuros.amount, 20);
assert.strictEqual(twentyEuros.currency, "EUR");

// 测试代码3，针对除法
let originalMoney = new Money(4002, "KRW")
let actualMoneyAfterDivision = originalMoney.divide(4)
let expectedMoneyAfterDivision = new Money(1000.5, "KRW")
assert.deepStrictEqual(actualMoneyAfterDivision, expectedMoneyAfterDivision)

// 测试代码3
let fifteenDollars = new Money(15, "USD");
let portfolio = new Portfolio();
portfolio.add(fiveDollars, tenDollars);
assert.deepStrictEqual(portfolio.evaluate("USD"), fifteenDollars);