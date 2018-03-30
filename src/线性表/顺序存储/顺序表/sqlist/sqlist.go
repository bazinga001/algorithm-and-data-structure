/*
NAME:			顺序表
AUTHOR:			bazinga
CREATETIME: 	2018-03-28
*/

package sqlist

import (
	"errors"
	"fmt"
)

// ElemType 元素类型
type ElemType int

// SqList 定义
type SqList struct {
	data   []ElemType // 元素集合
	Length int        // 当前长度
	size   int        // 最大容量
}

// Init 初始化List, size表示线性表最大容量, args为初始表数据
func Init(size int, args ...ElemType) (s *SqList, err error) {
	if len(args) > size {
		err = errors.New("表数据长度大于最大容量")
		return nil, err
	}
	s = &SqList{make([]ElemType, size), 0, size}
	for i := 0; i < len(args); i++ {
		s.Insert(i, args[i])
	}
	return s, nil
}

// GetElem 获取第i个元素
func (s SqList) GetElem(i int) (e ElemType, err error) {
	if i < 0 || i > s.size {
		err = errors.New("索引越界")
		return 0, err
	}
	return s.data[i], nil
}

// Insert 第i个位置插入元素e
func (s *SqList) Insert(i int, e ElemType) (err error) {
	if i < 0 || i > s.size {
		err = errors.New("索引越界")
		return err
	}
	for j := s.Length - 1; j >= i; j-- {
		s.data[j+1] = s.data[j]
	}
	s.data[i] = e
	s.Length++
	return nil
}

// Delete 删除第i个位置的元素
func (s *SqList) Delete(i int) (err error) {
	if i < 0 || i > s.size {
		err = errors.New("索引越界")
		return err
	}
	for j := i; j < s.Length; j++ {
		s.data[j] = s.data[j+1]
	}
	s.Length--
	return nil
}

// Traverse 遍历list
func (s SqList) Traverse() {
	for i := 0; i < s.Length; i++ {
		fmt.Print(s.data[i], " ")
	}
	fmt.Println()
}
