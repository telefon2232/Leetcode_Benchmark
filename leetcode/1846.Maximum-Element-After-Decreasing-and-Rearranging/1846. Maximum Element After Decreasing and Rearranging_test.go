package leetcode

import (
	"fmt"
	"testing"
)

type question1846 struct {
	para1846
	ans1846
}

// para 是参数
// one 代表第一个参数
type para1846 struct {
	arr []int
}

// ans 是答案
// one 代表第一个答案
type ans1846 struct {
	one int
}

func Benchmark_Problem1846(b *testing.B) {

	qs := []question1846{

		{
			para1846{[]int{2, 2, 1, 2, 1}},
			ans1846{2},
		},

		{
			para1846{[]int{100, 1, 1000}},
			ans1846{3},
		},

		{
			para1846{[]int{1, 2, 3, 4, 5}},
			ans1846{5},
		},
	}


	for _, q := range qs {
		_, p := q.ans1846, q.para1846
		fmt.Printf("【input】:%v       【output】:%v\n", p, maximumElementAfterDecrementingAndRearranging(p.arr))
	}
}
