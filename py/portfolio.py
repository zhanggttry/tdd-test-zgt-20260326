import functools
import operator
from money import Money
    
# 产品代码，定义Porfolio类
class Portfolio:
    def __init__(self):
        self.moneys = []
    def add(self, *moneys):
        self.moneys.extend(moneys)
    
    def evaluate(self, currency):
        total = functools.reduce(
            operator.add, map(lambda m: m.amount, self.moneys), 0)
        return Money(total, currency)
    