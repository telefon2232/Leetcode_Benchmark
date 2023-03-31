package leetcode

import (
	"fmt"
	"testing"
)

type question1054 struct {
	para1054
	ans1054
}

// para 是参数
// one 代表第一个参数
type para1054 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1054 struct {
	one []int
}

func Benchmark_Problem1054(b *testing.B) {

	qs := []question1054{

		{
			para1054{[]int{1, 1, 1, 2, 2, 2}},
			ans1054{[]int{2, 1, 2, 1, 2, 1}},
		},

		{
			para1054{[]int{1, 1, 1, 1, 2, 2, 3, 3}},
			ans1054{[]int{1, 3, 1, 3, 2, 1, 2, 1}},
		},
	}


	for _, q := range qs {
		_, p := q.ans1054, q.para1054
		fmt.Printf("【input】:%v       【output】:%v\n", p, rearrangeBarcodes(p.one))
	}
}
