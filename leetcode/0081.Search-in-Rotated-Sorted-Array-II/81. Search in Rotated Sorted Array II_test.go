package leetcode

import (
	"fmt"
	"testing"
)

type question81 struct {
	para81
	ans81
}

// para 是参数
// one 代表第一个参数
type para81 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans81 struct {
	one bool
}

func Benchmark_Problem81(b *testing.B) {

	qs := []question81{

		{
			para81{[]int{2, 5, 6, 0, 0, 1, 2}, 0},
			ans81{true},
		},

		{
			para81{[]int{2, 5, 6, 0, 0, 1, 2}, 3},
			ans81{false},
		},
	}


	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans81, q.para81
		fmt.Printf("【input】:%v    【output】:%v\n", p, search(p.nums, p.target))
	}
}}}
