import functools
import operator

# 产品代码
class Money:
    def __init__(self, amount, currency):
        self.amount = amount
        self.currency = currency
    
    # 覆写__eq__方法
    def __eq__(self, other):
        return self.amount == other.amount and self.currency == other.currency
    
    # 覆写__str__方法
    def __str__(self):
        return f"{self.currency} {self.amount:0.2f}"
    
    # 定义方法1
    def times(self, multiplier):
        return Money(self.amount * multiplier, self.currency)
    
    # 定义产品方法2
    def divide(self, divisor):
        return Money(self.amount / divisor, self.currency)