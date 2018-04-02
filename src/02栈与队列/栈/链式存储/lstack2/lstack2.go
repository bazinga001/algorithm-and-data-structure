/*
NAME:			单链表
AUTHOR:			bazinga
CREATETIME: 	2018-04-01

NOTE:			栈顶放到链表的头部, 方便进行插入删除操作
				github.com/golang-collections/collections/stack
				看了这里的源码后有启发, 重新定义类型, 重新实现.
				之前的线性表暂时不更改.
*/

package lstack

import "fmt"

type node struct {
	data interface{}
	next *node
}

// Stack 栈结构
type Stack struct {
	top *node
}

// New 新建一个栈
func New() *Stack {
	return new(Stack)
}

// Traverse 遍历栈
func (s Stack) Traverse() {
	// 如果用 recevier 是 (s *Stack) 下面的代码会改变top指针的位置
	// 可用一个临时变量替代 s.top

	for s.top != nil {
		fmt.Print(s.top.data, " ")
		s.top = s.top.next
	}
	fmt.Println()
}

// Push 入栈
func (s *Stack) Push(value interface{}) {
	newNode := &node{data: value}
	newNode.next = s.top
	s.top = newNode
}

// Pop 出栈
func (s *Stack) Pop() interface{} {
	value := s.top.data
	s.top = s.top.next
	return value
}

// Peek 返回栈顶元素
func (s Stack) Peek() interface{} {
	return s.top.data
}

// IsEmpty 判断栈是否为空
func (s Stack) IsEmpty() bool {
	return s.top == nil
}
