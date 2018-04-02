/*
NAME:			单链表
AUTHOR:			bazinga
CREATETIME: 	2018-04-01

NOTE:			栈顶放到链表的头部, 方便进行插入删除操作
*/

package lstack

import "fmt"

type ElemType int

type Node struct {
	data ElemType
	next *Node
}

func Init() *Node {
	// 值传递, p在整个生命期调用其方法后, 值都没有发生变化
	p := new(Node)
	return p
}

func (p *Node) Push(e ElemType) {
	newNode := &Node{data: e}
	newNode.data = e
	newNode.next = p.next
	p.next = newNode

	// 不能是这样,原因Init方法中
	// newNode.next = p.next
	// p.next = newNode
}

func (p *Node) Pop() (e ElemType) {
	e = p.next.data
	p.next = p.next.next
	return e
}

func (p *Node) Traverse() {
	p = p.next
	for p != nil {
		fmt.Print(p.data, " ")
		p = p.next
	}
}
