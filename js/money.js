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
module.exports = Money;