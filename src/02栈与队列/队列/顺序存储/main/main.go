package main

import (
	"02栈与队列/队列/顺序存储/sqqueue"
	"fmt"
)

func main() {
	sqq := sqqueue.New(10)
	sqq.EnQueue(1)
	sqq.EnQueue(2)
	sqq.EnQueue(3)
	sqq.EnQueue(4)
	sqq.EnQueue(5)
	sqq.EnQueue(6)
	sqq.EnQueue(7)

	sqq.Traverse()

	fmt.Println("DeQueue: ", sqq.DeQueue())
	fmt.Println("DeQueue: ", sqq.DeQueue())
	fmt.Println("DeQueue: ", sqq.DeQueue())
	fmt.Println("DeQueue: ", sqq.DeQueue())

	sqq.EnQueue(8)
	sqq.EnQueue(9)

	fmt.Println("DeQueue: ", sqq.DeQueue())
	fmt.Println("DeQueue: ", sqq.DeQueue())

}
