package main

import (
	"01线性表/顺序存储/顺序表/sqlist"
	"fmt"
)

func main() {
	s, _ := sqlist.Init(10, 100, 200)
	s.Traverse()
	s.Insert(0, 2)
	s.Traverse()
	s.Insert(1, 4)
	s.Traverse()
	s.Insert(2, 8)
	s.Traverse()
	s.Insert(3, 16)
	s.Traverse()
	s.Insert(4, 32)
	s.Traverse()
	s.Insert(5, 64)
	s.Traverse()
	s.Insert(6, 128)
	s.Traverse()
	s.Delete(4)
	s.Traverse()
	fmt.Println(s.Length)
}
