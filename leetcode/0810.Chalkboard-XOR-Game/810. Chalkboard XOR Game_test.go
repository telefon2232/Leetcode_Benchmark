package leetcode

import (
	"fmt"
	"testing"
)

type question810 struct {
	para810
	ans810
}

// para 是参数
// one 代表第一个参数
type para810 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans810 struct {
	one bool
}

func Benchmark_Problem810(b *testing.B) {

	qs := []question810{

		{
			para810{[]int{1, 1, 2}},
			ans810{false},
		},
	}


	for _, q := range qs {
		_, p := q.ans810, q.para810
		fmt.Printf("【input】:%v       【output】:%v\n", p, xorGame(p.nums))
	}
}
