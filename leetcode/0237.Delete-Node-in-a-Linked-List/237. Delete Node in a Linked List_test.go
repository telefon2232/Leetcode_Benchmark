package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question237 struct {
	para237
	ans237
}

// para 是参数
// one 代表第一个参数
type para237 struct {
	one []int
	n   int
}

// ans 是答案
// one 代表第一个答案
type ans237 struct {
	one []int
}

func Benchmark_Problem237(b *testing.B) {

	qs := []question237{

		{
			para237{[]int{1, 2, 3, 4, 5}, 1},
			ans237{[]int{2, 3, 4, 5}},
		},

		{
			para237{[]int{1, 2, 3, 4, 5}, 2},
			ans237{[]int{1, 3, 4, 5}},
		},

		{
			para237{[]int{1, 1, 1, 1, 1}, 1},
			ans237{[]int{}},
		},

		{
			para237{[]int{1, 2, 3, 4, 5}, 5},
			ans237{[]int{1, 2, 3, 4}},
		},

		{
			para237{[]int{}, 5},
			ans237{[]int{}},
		},

		{
			para237{[]int{1, 2, 3, 4, 5}, 10},
			ans237{[]int{1, 2, 3, 4, 5}},
		},

		{
			para237{[]int{1}, 1},
			ans237{[]int{}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans237, q.para237
				(structures.List2Ints(removeElements(structures.Ints2List(p.one), p.n)))
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
