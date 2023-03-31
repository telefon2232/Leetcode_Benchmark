package leetcode

import (
	"fmt"
	"testing"
)

type question78 struct {
	para78
	ans78
}

// para 是参数
// one 代表第一个参数
type para78 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans78 struct {
	one [][]int
}

func Benchmark_Problem78(b *testing.B) {

	qs := []question78{

		{
			para78{[]int{}},
			ans78{[][]int{{}}},
		},

		{
			para78{[]int{1, 2, 3}},
			ans78{[][]int{{}, {1}, {2}, {3}, {1, 2}, {2, 3}, {1, 3}, {1, 2, 3}}},
		},
	}


	for _, q := range qs {
		_, p := q.ans78, q.para78
		fmt.Printf("【input】:%v       【output】:%v\n", p, subsets(p.one))
	}
}
