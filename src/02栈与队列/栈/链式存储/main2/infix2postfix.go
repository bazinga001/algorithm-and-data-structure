/*
NAME:			中缀表达式转为后缀表达式
AUTHOR:			bazinga
CREATETIME: 	2018-04-02

NOTE:			(1-2)*(4+5) --> 1 2 - 4 5 + *
				中缀表达式中不包含空格
				对于负数, 应该写成表达式的形式 即: -12 应该写成 0-12
*/

package main

import (
	lstack "02栈与队列/栈/链式存储/lstack2"
	"fmt"
	"strconv"
)

// a 的优先级是否大于等于 b
func isPriority(a, b rune) bool {
	if a == '*' || a == '/' {
		return true
	}
	return false
}

func main() {
	str := "(12-2)*(4+5)"
	symstack := lstack.New()
	for i, v := range str {
		switch v {
		case '(':
			symstack.Push('(')
		case ')':
			for symstack.Peek() != '(' {
				fmt.Print(string(symstack.Pop().(rune)), " ")
			}
			symstack.Pop() // '(' 出栈
		case '+':
			if symstack.IsEmpty() {
				symstack.Push('+')
				continue
			}
			for isPriority(symstack.Peek().(rune), '+') {
				symstack.Pop()
			}
			symstack.Push('+')
		case '-':
			if symstack.IsEmpty() {
				symstack.Push('-')
				continue
			}
			for isPriority(symstack.Peek().(rune), '-') {
				symstack.Pop()
			}
			symstack.Push('-')
		case '*':
			if symstack.IsEmpty() {
				symstack.Push('*')
				continue
			}
			for isPriority(symstack.Peek().(rune), '*') {
				symstack.Pop()
			}
			symstack.Push('*')
		case '/':
			if symstack.IsEmpty() {
				symstack.Push('/')
				continue
			}
			for isPriority(symstack.Peek().(rune), '/') {
				symstack.Pop()
			}
			symstack.Push('/')
		// case ' ':
		default:
			val, _ := strconv.Atoi(string(v))
			fmt.Print(val)
			if isOperatorOrBrackets(str[i+1]) {
				fmt.Print(" ")
			}
		}
	}
	// 打印最后一个符号
	fmt.Print(string(symstack.Pop().(rune)), " ")
}

func isOperatorOrBrackets(s byte) bool {
	return s == '+' || s == '-' || s == '*' ||
		s == '/' || s == '(' || s == ')'
}
