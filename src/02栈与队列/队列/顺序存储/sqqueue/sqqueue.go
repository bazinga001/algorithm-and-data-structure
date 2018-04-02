/*
NAME:			顺序(循环)队列
AUTHOR:			bazinga
CREATETIME:		2018-04-03

NOTE:			data[tail]------>--->-----data[head]
*/

package sqqueue

import "fmt"

// Queue 队列结构
type Queue struct {
	data       []interface{}
	head, tail int
	size       int
}

// New 新建一个循环队列
func New(size int) *Queue {
	return &Queue{
		data: make([]interface{}, size),
		head: 0,
		tail: 0,
		size: size,
	}
}

// Traverse 遍历队列
func (q Queue) Traverse() {
	for q.tail != q.head {
		fmt.Print(q.data[q.head], " ")
		q.head = (q.head - 1 + q.size) % q.size
	}
	fmt.Println()
}

// EnQueue 入队列
func (q *Queue) EnQueue(value interface{}) {
	q.data[q.tail] = value
	q.tail = (q.tail - 1 + q.size) % q.size
}

// DeQueue 出队列
func (q *Queue) DeQueue() (value interface{}) {
	value = q.data[q.head]
	q.head = (q.head - 1 + q.size) % q.size
	return value
}
