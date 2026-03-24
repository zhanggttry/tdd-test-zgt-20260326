// 编写第1个测试用例
const assert = require('assert');
const Money = require('./money');
const Portfolio = require('./portfolio');
//将测试规整到1个类里
class MoneyTest{
    testMultiplication() {
        // 测试代码1，针对乘法
        let tenEuros = new Money(10, "EUR");
        let twentyEuros = new Money(20, "EUR");
        assert.deepStrictEqual(tenEuros.times(2), twentyEuros);

    }
    testDivision(){
        // 测试代码2，针对除法
        let originalMoney = new Money(4002, "KRW")
        let actualMoneyAfterDivision = originalMoney.divide(4)
        let expectedMoneyAfterDivision = new Money(1000.5, "KRW")
        assert.deepStrictEqual(actualMoneyAfterDivision, expectedMoneyAfterDivision)
    }
    testAddition() {
    // 测试代码3，针对汇率转换
    let fiveDollars = new Money(5, "USD");
    let tenDollars = new Money(10, "USD")
    let fifteenDollars = new Money(15, "USD")
    let portfolio = new Portfolio();
    portfolio.add(fiveDollars, tenDollars);
    assert.deepStrictEqual(portfolio.evaluate("USD"), fifteenDollars);
    }

    getAllTestMethods() {
        let moneyPrototype = MoneyTest.prototype; // 获取MoneyTest类的原型对象
        let allProps = Object.getOwnPropertyNames(moneyPrototype); // 获取原型对象的所有属性名
        let testMethods = allProps.filter(p => {
            return typeof moneyPrototype[p] === 'function' && p.startsWith('test'); // 过滤出以test开头的方法
        });
        return testMethods;
    }

    runAllTests() {
        let testMethods = this.getAllTestMethods(); // 获取所有的测试方法名
        testMethods.forEach(m => {
            let method = Reflect.get(this, m); // 获取该方法的method对象
            Reflect.apply(method, this, []); // 在this对象上调用method所表示的测试方法
        })
        this.testMultiplication();
        this.testDivision();
        this.testAddition();
    }
}

//执行所有测试
new MoneyTest().runAllTests();


