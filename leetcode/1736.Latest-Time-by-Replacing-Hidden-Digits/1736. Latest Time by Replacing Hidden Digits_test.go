package leetcode

import (
	"testing"
)

type question1736 struct {
	para1736
	ans1736
}

// para 是参数
// one 代表第一个参数
type para1736 struct {
	time string
}

// ans 是答案
// one 代表第一个答案
type ans1736 struct {
	one string
}

func Benchmark_Problem1736(b *testing.B) {

	qs := []question1736{

		{
			para1736{"2?:?0"},
			ans1736{"23:50"},
		},

		{
			para1736{"0?:3?"},
			ans1736{"09:39"},
		},

		{
			para1736{"1?:22"},
			ans1736{"19:22"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1736, q.para1736
				(maximumTime(p.time))
			}
		}
	}
}
