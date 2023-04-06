package leetcode

import (
	"testing"
)

type question492 struct {
	para492
	ans492
}

// area 是参数
type para492 struct {
	area int
}

// ans 是答案
type ans492 struct {
	ans []int
}

func Benchmark_Problem492(b *testing.B) {

	qs := []question492{

		{
			para492{4},
			ans492{[]int{2, 2}},
		},

		{
			para492{37},
			ans492{[]int{37, 1}},
		},

		{
			para492{122122},
			ans492{[]int{427, 286}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans492, q.para492
				(constructRectangle(p.area))
			}
		}
	}
}
