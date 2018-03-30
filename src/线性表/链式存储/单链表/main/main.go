package main

import (
	"fmt"
	"线性表/链式存储/单链表/sllist"
)

func main() {
	p := sllist.Init()
	p.Insert(0, 2)
	p.Traverse()
	p.Insert(1, 4)
	p.Traverse()
	p.Insert(0, 8)
	p.Traverse()
	p.Insert(3, 16)
	p.Traverse()
	p.Insert(4, 32)
	p.Traverse()
	p.Insert(3, 64)
	p.Traverse()
	p.Insert(2, 128)
	p.Traverse()
	// p.Delete(3)
	// p.Traverse()
	res, _ := p.Delete(0)
	p.Traverse()
	fmt.Println("Delete: ", res)

	res, _ = p.Delete(5)
	p.Traverse()
	fmt.Println("Delete: ", res)

}
