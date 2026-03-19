# 编写第一个TDD测试用例，2026年03月17日09:00:00
import unittest

# 产品代码

class Money:
    def __init__(self, amount, currency):
        self.amount = amount
        self.currency = currency
    
    # 定义方法1
    def times(self, multiplier):
        return Money(self.amount * multiplier, self.currency)
    
    # 定义产品方法2
    def divide(self, divisor):
        return Money(self.amount / divisor, self.currency)
# 测试用例代码1
class TestMoney(unittest.TestCase):
    
    # 测试代码1
    def testMultiplicationInDollars(self):
        fiver = Money(5, "USD")
        tenner = fiver.times(2)
        self.assertEqual(10, tenner.amount)
        self.assertEqual("USD", tenner.currency)
    
    # 测试代码2
    def testMultiplicationInEuros(self):
        tenEuros = Money(10, "EUR")
        twentyEuros = tenEuros.times(2)
        self.assertEqual(20, twentyEuros.amount)
        self.assertEqual("EUR", twentyEuros.currency)
        
    # 测试代码3
    def testDivision(self):
        originalMoney = Money(4002, "KRW")
        actualMoneyAfterDivision = originalMoney.divide(4)
        expectedMoneyAfterDivision = Money(1000.5, "KRW")
        self.assertEqual(expectedMoneyAfterDivision.amount,
                         actualMoneyAfterDivision.amount)
        self.assertEqual(expectedMoneyAfterDivision.currency,
                         actualMoneyAfterDivision.currency)    
        
if __name__ == '__main__':
    unittest.main()