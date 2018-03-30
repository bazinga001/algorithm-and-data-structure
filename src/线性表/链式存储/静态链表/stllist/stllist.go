/*
NAME:			静态链表
AUTHOR:			bazinga
CREATETIME: 	2018-03-30

NOTE:			采用单例模式
REFERENCE :		http://blog.ralch.com/tutorial/design-patterns/golang-singleton/
*/

/*
1. 数组的第一个和最后一个元素不存放有效数据
2. 下标为0的元素的cur指向备用链表第一个结点的下标
3. 下标为(size-1)的元素的cur指向有效链表第一个结点的下标
*/

package stllist

import "fmt"

// ElemType 元素类型
type ElemType int

// Node 结点类型
type Node struct {
	data ElemType
	cur  int
}

// ListType 链表类型
type ListType []Node

// 单例, 所有的方法的receiver都是list
var list ListType

// 链表容量
var size int

// Init 初始化
func Init(s int) ListType {
	size = s
	list = make([]Node, size)
	for i := 0; i < size-1; i++ {
		list[i].cur = i + 1
	}
	list[size-1].cur = 0 // 规定当list[i].cur == 0 有效链表截至
	return list
}

// Traverse 显示所有静态链表数据
func (l ListType) Traverse() {
	for i, v := range l {
		fmt.Println("index:", i, "data:", v.data, "cur:", v.cur)
	}
	fmt.Println()
}

// Insert 向有效链表第i个位置后添加结点
// 没有对i位置的合法性进行判断
func (l ListType) Insert(i int, e ElemType) {
	// 申请空闲地址的下标
	index := malloc()
	l[index].data = e

	cur := size - 1
	for j := 0; j < i; j++ {
		cur = l[cur].cur
	}

	// core
	l[index].cur = l[cur].cur
	l[cur].cur = index

	// 当前有效链表长度为0
	if l[size-1].cur == 0 {
		l[size-1].cur = index
		l[index].cur = 0
	}
}

func malloc() (c int) {
	c = list[0].cur
	list[0].cur = list[list[0].cur].cur
	return c
}

// func (l ListType) Delete(i int) (e ElemType) {
// 	cur := size - 1
// 	if i == 0 {
// 		temp := l[cur].cur
// 		e = l[temp].data
// 		l[cur].cur = l[l[cur].cur].cur
// 		free(temp)
// 	} else {
// 		for j := 0; j < i; j++ {
// 			cur = l[cur].cur
// 		}
// 		temp := l[cur].cur
// 		e = l[temp].data
// 		l[cur].cur = l[l[cur].cur].cur
// 		free(temp)
// 	}

// 	return e
// }

// Delete 删除有效链表中第i个结点,并将结点空间释放到备份链表中
func (l ListType) Delete(i int) (e ElemType) {
	cur := size - 1
	for j := 0; j < i; j++ {
		cur = l[cur].cur
	}
	temp := l[cur].cur
	e = l[temp].data
	l[cur].cur = l[l[cur].cur].cur
	free(temp)

	return e
}

// free 将下标为i的空间释放到备份链表的首部
func free(i int) {
	list[i].cur = list[0].cur
	list[0].cur = i
}

// TraverseValidList 遍历有效链表段
func (l ListType) TraverseValidList() {
	i := l[size-1].cur
	if i == 0 {
		return
	}
	for {
		fmt.Println("index:", i, "data:", l[i].data, "cur:", l[i].cur)
		if l[i].cur == 0 {
			break
		}
		i = l[i].cur
	}
}
