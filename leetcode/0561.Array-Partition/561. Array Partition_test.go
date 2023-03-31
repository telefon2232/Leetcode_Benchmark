package leetcode

import (
	"fmt"
	"testing"
)

type question561 struct {
	para561
	ans561
}

// para 是参数
// one 代表第一个参数
type para561 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans561 struct {
	one int
}

func Benchmark_Problem561(b *testing.B) {

	qs := []question561{

		{
			para561{[]int{}},
			ans561{0},
		},

		{
			para561{[]int{1, 4, 3, 2}},
			ans561{4},
		},

		// 如需多个测试，可以复制上方元素。
	}


	for _, q := range qs {
		_, p := q.ans561, q.para561
		fmt.Printf("【input】:%v       【output】:%v\n", p, arrayPairSum(p.nums))
	}
}
