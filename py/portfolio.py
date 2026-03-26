import functools
import operator
from money import Money
    
# 产品代码，定义Porfolio类
class Portfolio:
    def __init__(self):
        self.moneys = []
        # 定义1个私有变量
        self._eur_to_usd = 1.2
        
    # 编写__convert方法
    def __convert(self,aMoney, aCurrency):
        exchangeRates = {'EUR->USD': 1.2, 'USD->KRW': 1100} # 加一个字典
        if aMoney.currency == aCurrency:
            return aMoney.amount
        else:
            key = aMoney.currency + '->' + aCurrency
            return aMoney.amount * exchangeRates[key]
        
    def add(self, *moneys):
        self.moneys.extend(moneys)
    
    def evaluate(self, currency):
        total = functools.reduce(
            operator.add, map(lambda m: self.__convert(m, currency), self.moneys), 0)
        return Money(total, currency)
    