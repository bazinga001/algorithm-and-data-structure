package main

import "02栈与队列/栈/链式存储/lstack"

func main() {
	ls := lstack.Init()
	ls.Push(1)
	ls.Push(2)
	ls.Push(3)
	ls.Push(4)
	ls.Push(1)
	ls.Push(1)
	ls.Push(1)

	ls.Pop()
	ls.Pop()
	ls.Pop()
	ls.Pop()

	ls.Traverse()
}
