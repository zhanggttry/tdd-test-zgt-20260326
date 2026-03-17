// 编写第1个测试用例
const assert = require('assert');

// 产品代码
class Dollar {
    //创建1个构造器
    constructor(amount) {
        this.amount = amount;
    }
    times(multiplier) {
        return new Dollar(this.amount * multiplier); // 返回1个表示乘积美元的对象
    }
}

// 测试代码
let fiver = new Dollar(5);
let tenner = fiver.times(2);
assert.strictEqual(tenner.amount, 10);