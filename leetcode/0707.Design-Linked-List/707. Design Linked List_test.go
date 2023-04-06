package leetcode

import (
	"fmt"
	"testing"
)

func Benchmark_Problem707(b *testing.B) {
	obj := Constructor()
	(MList2Ints(&obj))

	(MList2Ints(&obj))
	obj.AddAtHead(38)
	(MList2Ints(&obj))
	obj.AddAtHead(45)
	(MList2Ints(&obj))
	obj.DeleteAtIndex(2)
	(MList2Ints(&obj))
	obj.AddAtIndex(1, 24)
	(MList2Ints(&obj))
	obj.AddAtTail(36)
	(MList2Ints(&obj))
	obj.AddAtIndex(3, 72)
	obj.AddAtTail(76)
	(MList2Ints(&obj))
	obj.AddAtHead(7)
	(MList2Ints(&obj))
	obj.AddAtHead(36)
	(MList2Ints(&obj))
	obj.AddAtHead(34)
	(MList2Ints(&obj))
	obj.AddAtTail(91)
	(MList2Ints(&obj))

	obj1 := Constructor()
	(MList2Ints(&obj1))
	param2 := obj1.Get(0)
	fmt.Println(param2)
	(MList2Ints(&obj1))
	obj1.AddAtIndex(1, 2)
	(MList2Ints(&obj1))
	param2 = obj1.Get(0)
	(MList2Ints(&obj1))
	param2 = obj1.Get(1)
	(MList2Ints(&obj1))
	obj1.AddAtIndex(0, 1)
	(MList2Ints(&obj1))
	param2 = obj1.Get(0)
	(MList2Ints(&obj1))
	param2 = obj1.Get(1)
	(MList2Ints(&obj1))
}

func MList2Ints(head *MyLinkedList) []int {
	res := []int{}
	cur := head.head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}
