/*
NAME:			循环链表
AUTHOR:			bazinga
CREATETIME: 	2018-03-29
*/

package cirlist

import "fmt"

// ElemType 元素类型
type ElemType int

// Node 结点类型
type Node struct {
	data ElemType
	next *Node
}

// 需要维护头尾指针
var head, tail *Node

// Init 初始化循环链表, i(i>1)表示结点个数
func Init(i int) *Node {
	head = &Node{ElemType(1), nil}
	tail = head
	for j := 0; j < i-1; j++ {
		tail.next = &Node{ElemType(j + 2), nil}
		tail = tail.next
	}
	tail.next = head
	return head
}

// Traverse 遍历循环链表
func (p Node) Traverse() {
	cur := tail.next
	for cur != tail {
		fmt.Print(cur.data, " ")
		cur = cur.next
	}
	fmt.Print(tail.data)
	fmt.Println()
}

// Delete 删除从头指针位置开始的第i个位置的结点
func (p Node) Delete(i int) (e ElemType) {
	cur := tail
	for j := 0; j < i; j++ {
		cur = cur.next
	}
	e = cur.next.data
	cur.next = cur.next.next
	return e
}

// GoAhead 头尾指针都向前移动一个单位
func (p Node) GoAhead() {
	tail = tail.next
	head = head.next
}
