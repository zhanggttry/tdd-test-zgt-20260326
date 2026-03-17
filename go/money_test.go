package main

import (
	"testing"
)

// 测试代码
func TestMultiplication(t *testing.T) {
	// 创建一个金额为5的 Dollar 对象
	fiver := Dollar{
		amount: 5,
	}
	// 调用 Times 方法将金额乘以2，预期结果为10
	tenner := fiver.Times(2)
	// 验证计算结果是否符合预期，如果不符合则输出错误信息
	if tenner.amount != 10 {
		t.Errorf("期望10，计算得到: [%d]", tenner.amount)
	}
}

type Dollar struct {
	amount int
}

func (d Dollar) Times(multiplier int) Dollar {
	return Dollar{amount: d.amount * multiplier}
}
