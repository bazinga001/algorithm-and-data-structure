/*
NAME:			顺序存储栈
AUTHOR:			bazinga
CREATETIME: 	2018-03-31

NOTE:			初始分配最大容量, 后续不再扩容
				top指向栈顶的下一个元素
*/

package sqstack

import "fmt"

// ElemType 元素类型
type ElemType int

// SqStack 顺序栈数据
type SqStack struct {
	data []ElemType // 数据域
	top  int        // 栈顶指针
	base int        // 栈底指针
	size int        // 栈的容量
}

// Init 初始化顺序栈
func Init(size int) *SqStack {
	s := new(SqStack)
	s.data = make([]ElemType, size)
	s.top = 0
	s.base = 0
	s.size = size
	return s
}

// Traverse 栈的遍历
func (s *SqStack) Traverse() {
	cur := s.base
	for cur != s.top {
		fmt.Print(s.data[cur], " ")
		cur++
	}
	fmt.Println()
}

// Push 进栈
func (s *SqStack) Push(e ElemType) {
	s.data[s.top] = e
	s.top++
}

// Pop 出栈
func (s *SqStack) Pop() (e ElemType) {
	e = s.data[s.top-1]
	s.top--
	return e
}

// IsEempty 判断栈是否为空
func (s SqStack) IsEmpty() bool {
	return s.top == s.base
}
