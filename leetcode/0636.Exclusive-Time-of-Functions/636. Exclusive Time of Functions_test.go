package leetcode

import (
	"testing"
)

type question636 struct {
	para636
	ans636
}

// para 是参数
// one 代表第一个参数
type para636 struct {
	n   int
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans636 struct {
	one []int
}

func Benchmark_Problem636(b *testing.B) {

	qs := []question636{

		{
			para636{2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:7", "1:end:7", "0:end:8"}},
			ans636{[]int{8, 1}},
		},

		{
			para636{2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"}},
			ans636{[]int{7, 1}},
		},

		{
			para636{2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}},
			ans636{[]int{3, 4}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans636, q.para636
				(exclusiveTime(p.n, p.one))
			}
		}
	}
}
