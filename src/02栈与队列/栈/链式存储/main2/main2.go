package main

import (
	lstack "02栈与队列/栈/链式存储/lstack2"
	"fmt"
)

func main2() {
	s := lstack.New()
	fmt.Println(s.IsEmpty())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	s.Traverse()

	// fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Peek())
	fmt.Println(s.Peek())

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())

}
