package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question2181 struct {
	para2181
	ans2181
}

// para 是参数
// one 代表第一个参数
type para2181 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans2181 struct {
	one []int
}

func Benchmark_Problem2181(b *testing.B) {

	qs := []question2181{

		{
			para2181{[]int{0, 3, 1, 0, 4, 5, 2, 0}},
			ans2181{[]int{4, 11}},
		},

		{
			para2181{[]int{0, 1, 0, 3, 0, 2, 2, 0}},
			ans2181{[]int{1, 3, 4}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans2181, q.para2181
				(structures.List2Ints(mergeNodes(structures.Ints2List(p.one))))
			}
		}
	}
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return newHead.Next
}
