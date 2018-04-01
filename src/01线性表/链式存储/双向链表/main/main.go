package main

import (
	"01线性表/链式存储/双向链表/dllist"
	"fmt"
)

func main() {
	dl := dllist.Init(0)
	// dl.Traverse()
	dl.Insert(1, 10)
	dl.Insert(2, 20)
	dl.Insert(0, 30)
	dl.Insert(0, 40)
	dl.Insert(0, 50)
	// dl.Traverse()
	dl.Insert(4, 60)
	dl.Traverse()

	fmt.Println(dl.Delete(0))
	fmt.Println(dl.Delete(0))
	fmt.Println(dl.Delete(0))
	fmt.Println(dl.Delete(0))
	fmt.Println(dl.Delete(0))
	fmt.Println(dl.Delete(0))
	fmt.Println(dl.Delete(0))
	// // dl.Traverse()
}
