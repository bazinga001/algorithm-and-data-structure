package main

import (
	"02栈与队列/队列/链式存储/lqueue"
	"fmt"
)

func main() {
	lq := lqueue.New()
	lq.EnQueue(1)
	lq.EnQueue(2)
	lq.EnQueue(3)
	lq.EnQueue(4)
	lq.Traverse()

	fmt.Println("DeQueue: ", lq.DeQueue())
	fmt.Println("DeQueue: ", lq.DeQueue())
	fmt.Println("Peek:", lq.Peek())
	fmt.Println("DeQueue: ", lq.DeQueue())
	fmt.Println("DeQueue: ", lq.DeQueue())
}
