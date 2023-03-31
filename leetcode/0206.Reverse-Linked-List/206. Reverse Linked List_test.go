package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question206 struct {
	para206
	ans206
}

// para 是参数
// one 代表第一个参数
type para206 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans206 struct {
	one []int
}

func Benchmark_Problem206(b *testing.B) {

	qs := []question206{

		{
			para206{[]int{1, 2, 3, 4, 5}},
			ans206{[]int{5, 4, 3, 2, 1}},
		},
	}

	for _, q := range qs {
		_, p := q.ans206, q.para206
		(structures.List2Ints(reverseList(structures.Ints2List(p.one))))
	}
}
