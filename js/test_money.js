// 编写第1个测试用例
const assert = require('assert');

// 产品代码
class Money {
    constructor(amount, currency) {
        this.amount = amount;
        this.currency = currency;
    }

    times(multiplier) {
        return new Money(this.amount * multiplier, this.currency);
    }

    // 定义除法
    divide(divisor) {
        return new Money(this.amount / divisor, this.currency);
    }
}

// 测试代码1
let fiver = new Money(5, "USD");
let tenner = fiver.times(2);
assert.strictEqual(tenner.amount, 10);
assert.strictEqual(tenner.currency, "USD");

// 测试代码2
let tenEuros = new Money(10, "EUR");
let twentyEuros = tenEuros.times(2);
assert.strictEqual(twentyEuros.amount, 20);
assert.strictEqual(twentyEuros.currency, "EUR");

// 测试代码3，针对除法
let originalMoney = new Money(4002, "KRW")
let actualMoneyAfterDivision = originalMoney.divide(4)
let expectedMoneyAfterDivision = new Money(1000.5, "KRW")
assert.deepStrictEqual(actualMoneyAfterDivision, expectedMoneyAfterDivision)