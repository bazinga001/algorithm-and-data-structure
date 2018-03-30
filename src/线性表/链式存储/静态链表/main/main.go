package main

import (
	"fmt"
	"线性表/链式存储/静态链表/stllist"
)

func main() {
	list := stllist.Init(10)
	list.Traverse()

	list.Insert(0, 88)
	list.Insert(0, 77)
	list.Insert(0, 66)
	list.Insert(0, 55)
	list.Insert(2, 44)
	list.Traverse()

	// list.TraverseValidList()
	fmt.Println("Deleted: ", list.Delete(4))
	fmt.Println("Deleted: ", list.Delete(3))
	fmt.Println("Deleted: ", list.Delete(2))
	fmt.Println("Deleted: ", list.Delete(1))
	fmt.Println("Deleted: ", list.Delete(0))
	// list.Traverse()
	list.TraverseValidList()

	fmt.Println("Insert: ")
	list.Insert(0, 100)
	list.TraverseValidList()

}
