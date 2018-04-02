/*
NAME:			链队列
AUTHOR:			bazinga
CREATETIME:		2018-04-03

NOTE:			head->[]->[].....[]->[tail]
				tail端添加结点, head端删除结点
*/

package lqueue

import "fmt"

type node struct {
	data interface{}
	next *node
}

// Queue 队列结构
type Queue struct {
	head, tail *node
}

// New 新建队列
func New() *Queue {
	return new(Queue)
}

// Traverse 队列遍历
func (q Queue) Traverse() {
	for q.head != nil {
		fmt.Print(q.head.data, " ")
		q.head = q.head.next
	}
	fmt.Println()
}

// EnQueue 入队列, 从队尾插入
func (q *Queue) EnQueue(value interface{}) {
	newNode := &node{data: value}
	if q.tail == nil { // 当前队列为空
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

// DeQueue 出队列, 从队首出队列
func (q *Queue) DeQueue() (value interface{}) {
	value = q.head.data
	q.head = q.head.next
	return value
}

// Peek 返回队首元素
func (q Queue) Peek() interface{} {
	return q.head.data
}
