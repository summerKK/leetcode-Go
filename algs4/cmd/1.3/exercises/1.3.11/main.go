package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/summerKK/leetcode-Go/algs4/cmd/1.3"
	"github.com/summerKK/leetcode-Go/algs4/utils"
)

/**
1.3.11
编写一段程序 EvaluatePostfix ，从标准输入中得到一个后序表达式，
求值并打印结果(将上一题的程序中得到的输出用管道传递给这一段程序可以得到和 Evaluate 相同的行为)。
*/
func main() {
	// ( ( 2 * 3 ) / ( 2 - 1 ) ) + ( 3 * ( 4 - 1 ) )
	s0 := "2 3 * 2 1 - / 3 4 1 - * +"
	result := EvaluatePostfix(s0)
	fmt.Printf("%s     result:%f\n", s0, result)
}

func EvaluatePostfix(s string) float64 {
	collection := strings.Split(s, " ")
	stack0 := &__3.MyStack{}
	stack0.Init(0)
	for _, s := range collection {
		if isDigital(s) {
			stack0.Push(utils.StrTo(s).MustFloat64())
		} else {
			v0 := stack0.Pop().(float64)
			v1 := stack0.Pop().(float64)
			var res float64 = 0
			switch s {
			case "+":
				res = v1 + v0
			case "-":
				res = v1 - v0
			case "*":
				res = v1 * v0
			case "/":
				res = v1 / v0
			}
			stack0.Push(res)
		}
	}

	return stack0.Pop().(float64)
}

func isDigital(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}

	return true
}
