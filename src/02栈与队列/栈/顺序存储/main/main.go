package main

import (
	"02栈与队列/栈/顺序存储/sqstack"
	"fmt"
)

func main1() {
	st := sqstack.Init(10)
	st.Push(1)
	st.Push(2)
	st.Push(3)
	// st.Push(4)
	// st.Push(5)
	// st.Push(6)
	// st.Push(7)
	// st.Push(8)
	st.Traverse()
	fmt.Println("Pop:", st.Pop())
	fmt.Println("Pop:", st.Pop())
	fmt.Println("Pop:", st.Pop())
	st.Traverse()
}

// 十进制转二进制
func main() {
	st := sqstack.Init(20)
	var num sqstack.ElemType = 256
	for num != 0 {
		st.Push(num % 2)
		num = num / 2
	}
	for !st.IsEmpty() {
		fmt.Print(st.Pop(), " ")
	}
	fmt.Println()
}
