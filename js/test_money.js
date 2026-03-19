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