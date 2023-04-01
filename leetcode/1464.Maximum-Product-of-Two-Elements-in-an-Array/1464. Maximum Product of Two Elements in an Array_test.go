package leetcode

import (
	"fmt"
	"testing"
)

type question1464 struct {
	para1464
	ans1464
}

// para 是参数
// one 代表第一个参数
type para1464 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1464 struct {
	one int
}

func Benchmark_Problem1464(b *testing.B) {

	qs := []question1464{

		{
			para1464{[]int{3, 4, 5, 2}},
			ans1464{12},
		},

		{
			para1464{[]int{1, 5, 4, 5}},
			ans1464{16},
		},

		{
			para1464{[]int{3, 7}},
			ans1464{12},
		},

		{
			para1464{[]int{1}},
			ans1464{0},
		},
	}


	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1464, q.para1464
		fmt.Printf("【input】:%v      ", p)
		fmt.Printf("【output】:%v      \n", maxProduct(p.nums))
	}
}}}
