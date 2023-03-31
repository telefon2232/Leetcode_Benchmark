package leetcode

import (
	"testing"
)

type question410 struct {
	para410
	ans410
}

// para 是参数
// one 代表第一个参数
type para410 struct {
	nums []int
	m    int
}

// ans 是答案
// one 代表第一个答案
type ans410 struct {
	one int
}

func Benchmark_Problem410(b *testing.B) {

	qs := []question410{

		{
			para410{[]int{7, 2, 5, 10, 8}, 2},
			ans410{18},
		},
	}

	for _, q := range qs {
		_, p := q.ans410, q.para410
		(splitArray(p.nums, p.m))
	}
}
