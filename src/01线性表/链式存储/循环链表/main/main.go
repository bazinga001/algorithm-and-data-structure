/*
解决约瑟夫环问题, 41人, Joseph给自己和朋友安排的两个位置是31,16
*/

package main

import (
	"01线性表/链式存储/循环链表/cirlist"
	"fmt"
)

func main() {
	c := cirlist.Init(41)
	// c.Traverse()
	// fmt.Println(c.Delete(0))
	// fmt.Println(c.Delete(0))
	// fmt.Println(c.Delete(0))
	// fmt.Println(c.Delete(0))
	// fmt.Println(c.Delete(0))
	// fmt.Println(c.Delete(0))
	// c.Traverse()

	for j := 0; j < 41; j++ {
		for i := 0; i < 2; i++ {
			c.GoAhead()
		}
		fmt.Println(c.Delete(0))
	}
}
