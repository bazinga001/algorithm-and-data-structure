/*
NAME:			逆波兰计算器
AUTHOR:			bazinga
CREATETIME: 	2018-04-02

NOTE:			目前支持浮点数运算
				负数需要在中缀表达式转为后缀表达式过程中处理 比如: "-11" --> "0 11 -"
*/

package main

import (
	lstack "02栈与队列/栈/链式存储/lstack2"
	"fmt"
	"strconv"
)

func main3() {
	// str := "3 2 +"
	// str := "32.5 2.5 + 5 *"
	str := "1 2 - 4 5 + *"

	stack := lstack.New()
	var result float64
	var floatValue float64
	var floatSlice []rune
	for i, val := range str {
		switch val {
		case '+':
			result = stack.Pop().(float64) + stack.Pop().(float64)
			stack.Push(result)
		case '-':
			result = -(stack.Pop()).(float64) + (stack.Pop()).(float64)
			stack.Push(result)
		case '*':
			result = (stack.Pop()).(float64) * (stack.Pop()).(float64)
			stack.Push(result)
		case '/':
			result = (stack.Pop()).(float64) + (stack.Pop()).(float64)
			stack.Push(result)
		case ' ':
			if !isOperator(str[i-1]) {
				fmt.Println("floatValue:", floatValue)
				stack.Push(floatValue)
			}
			floatSlice = nil
		default: // 当前值为数字或者小数点
			floatSlice = append(floatSlice, val)
			floatValue, _ = strconv.ParseFloat(string(floatSlice), 64)
		}
	}
	fmt.Println("result:", stack.Peek())
}

func isOperator(s byte) bool {
	return s == '+' || s == '-' || s == '*' || s == '/'
}
