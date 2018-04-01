/*
NAME:			双向链表
AUTHOR:			bazinga
CREATETIME: 	2018-03-31
*/

package dllist

import "fmt"

// ElemType 元素类型
type ElemType int

// Node 结点类型
type Node struct {
	data       ElemType
	prev, next *Node
}

// Init 初始化循环链表
func Init(args ...ElemType) *Node {
	p := new(Node)
	for i, v := range args {
		p.Insert(i, v)
	}
	return p
}

// Traverse 遍历, 不包含头结点
func (p *Node) Traverse() {
	for p.next != nil {
		p = p.next
		fmt.Print(p.data, " ")
	}
	fmt.Println()
}

// Insert 在第i个位置插入元素e
func (p *Node) Insert(i int, e ElemType) {
	newNode := &Node{data: e}

	for j := 0; j < i; j++ {
		p = p.next
	}
	if p.next != nil {
		// [p] [newNode] [p.next]
		newNode.prev = p      //[p]<-[newNode] [p.next]
		p.next.prev = newNode //[p]<-[newNode]<-[p.next]
		newNode.next = p.next //[p]<-[newNode]<=>[p.next]
		p.next = newNode      //[p]<=>[newNode]<=>[p.next]
	} else { // 插入到最后
		// [p] [newNode]
		p.next = newNode
		newNode.prev = p
	}

}

// Delete 删除双向链表第i个结点
func (p *Node) Delete(i int) (e ElemType) {
	for j := 0; j < i; j++ {
		p = p.next
	}
	e = p.next.data

	if p.next.next == nil { // 要删除的是最后一个结点
		p.next = nil
	} else {
		p.next.next.prev = p // 先左后右的指针连接顺序
		p.next = p.next.next

		/* //先右后左应该是这样, 避免这种写法
		temp := p.next
		p.next = p.next.next
		temp.next.prev = p
		*/
	}
	return e
}
