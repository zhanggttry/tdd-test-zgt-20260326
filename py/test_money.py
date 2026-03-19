# 编写第一个TDD测试用例，2026年03月17日09:00:00
import unittest

# 产品代码

class Money:
    def __init__(self, amount, currency):
        self.amount = amount
        self.currency = currency

    def times(self, multiplier):
        return Money(self.amount * multiplier, self.currency)
    
# 测试用例代码1
class TestMoney(unittest.TestCase):
    def testMultiplicationInDollars(self):
        fiver = Money(5, "USD")
        tenner = fiver.times(2)
        self.assertEqual(10, tenner.amount)
        self.assertEqual("USD", tenner.currency)
        
    def testMultiplicationInEuros(self):
        tenEuros = Money(10, "EUR")
        twentyEuros = tenEuros.times(2)
        self.assertEqual(20, twentyEuros.amount)
        self.assertEqual("EUR", twentyEuros.currency)
        
if __name__ == '__main__':
    unittest.main()