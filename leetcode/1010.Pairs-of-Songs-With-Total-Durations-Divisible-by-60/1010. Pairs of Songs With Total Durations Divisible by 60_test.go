package leetcode

import (
	"fmt"
	"testing"
)

type question1010 struct {
	para1010
	ans1010
}

// para 是参数
// one 代表第一个参数
type para1010 struct {
	time []int
}

// ans 是答案
// one 代表第一个答案
type ans1010 struct {
	one int
}

func Benchmark_Problem1010(b *testing.B) {

	qs := []question1010{

		{
			para1010{[]int{30, 20, 150, 100, 40}},
			ans1010{3},
		},

		{
			para1010{[]int{60, 60, 60}},
			ans1010{3},
		},
	}


	for _, q := range qs {
		_, p := q.ans1010, q.para1010
		fmt.Printf("【input】:%v       【output】:%v\n", p, numPairsDivisibleBy60(p.time))
	}
}
