# 编写第一个TDD测试用例，2026年03月17日09:00:00
import unittest

class Dollar:
    # 编写初始化器
    def __init__(self, amount):
      self.amount = amount
      
    def times(self, multiplier):
        return Dollar(10)
    
class TestMoney(unittest.TestCase):
    def testMultiplication(self):
        fiver = Dollar(5)
        tenner = fiver.times(2)
        self.assertEqual(10, tenner.amount)
        
if __name__ == '__main__':
    unittest.main()