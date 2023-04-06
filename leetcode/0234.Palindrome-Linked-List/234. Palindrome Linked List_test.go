package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question234 struct {
	para234
	ans234
}

// para 是参数
// one 代表第一个参数
type para234 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans234 struct {
	one bool
}

func Benchmark_Problem234(b *testing.B) {

	qs := []question234{

		{
			para234{[]int{1, 1, 2, 2, 3, 4, 4, 4}},
			ans234{false},
		},

		{
			para234{[]int{1, 1, 1, 1, 1, 1}},
			ans234{true},
		},

		{
			para234{[]int{1, 2, 2, 1, 3}},
			ans234{false},
		},

		{
			para234{[]int{1}},
			ans234{true},
		},

		{
			para234{[]int{}},
			ans234{true},
		},

		{
			para234{[]int{1, 2, 2, 2, 2, 1}},
			ans234{true},
		},

		{
			para234{[]int{1, 2, 2, 3, 3, 3, 3, 2, 2, 1}},
			ans234{true},
		},

		{
			para234{[]int{1, 2}},
			ans234{false},
		},

		{
			para234{[]int{1, 0, 1}},
			ans234{true},
		},

		{
			para234{[]int{1, 1, 2, 1}},
			ans234{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans234, q.para234
				(isPalindrome(structures.Ints2List(p.one)))
			}
		}
	}
}
