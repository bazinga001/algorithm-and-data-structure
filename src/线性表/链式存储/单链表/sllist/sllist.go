/*
NAME:			单链表
AUTHOR:			bazinga
CREATETIME: 	2018-03-29
*/

package sllist

import (
	"errors"
	"fmt"
)

// ElemType 元素类型
type ElemType int

// Node 结点类型
type Node struct {
	data ElemType
	next *Node
}

// Init 初始化单链表
func Init() (p *Node) {
	// p = &Node{0, nil} // or
	p = new(Node)
	return p
}

// GetElem 获取第i个元素
func (p *Node) GetElem(i int) (ElemType, error) {
	var j int
	for j := 0; j < i && p.next != nil; j++ {
		p = p.next
	}
	if j < i && p.next == nil {
		return 0, errors.New("越界")
	}
	return p.data, nil
}

// Insert 在第i个位置后插入元素e
func (p *Node) Insert(i int, e ElemType) error {
	var j int
	for j = 0; j < i && p.next != nil; j++ {
		p = p.next
	}
	if j < i && p.next == nil {
		return errors.New("越界")
	}
	// newNode := &Node{e, p.next} // or
	newNode := new(Node)
	newNode.data = e
	newNode.next = p.next

	p.next = newNode
	return nil
}

// Delete 删除第i个位置的元素
func (p *Node) Delete(i int) (ElemType, error) {
	var j int
	var pre *Node // 保存当前位置前一个结点的指针
	// 对pre必须赋值, 循环必须执行至少一次(j <= i)
	for j = 0; j <= i && p.next != nil; j++ {
		pre = p
		p = p.next
	}
	if j < i && p.next == nil {
		return 0, errors.New("越界")
	} else if j == i && p.next == nil { // 删除最后一个位置的元素
		res := p.data
		pre.next = nil
		return res, nil
	}
	res := pre.next.data
	pre.next = pre.next.next
	return res, nil
}

// Traverse 遍历
func (p *Node) Traverse() {
	for p.next != nil {
		p = p.next
		fmt.Print(p.data, " ")
	}
	fmt.Println()
}
