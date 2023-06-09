package leetcode

import (
	"testing"
)

type question978 struct {
	para978
	ans978
}

// para 是参数
// one 代表第一个参数
type para978 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans978 struct {
	one int
}

func Benchmark_Problem978(b *testing.B) {

	qs := []question978{

		{
			para978{[]int{0, 1, 1, 0, 1, 0, 1, 1, 0, 0}},
			ans978{5},
		},

		{
			para978{[]int{9, 9}},
			ans978{1},
		},

		{
			para978{[]int{9, 4, 2, 10, 7, 8, 8, 1, 9}},
			ans978{5},
		},

		{
			para978{[]int{4, 8, 12, 16}},
			ans978{2},
		},

		{
			para978{[]int{100}},
			ans978{1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans978, q.para978
				(maxTurbulenceSize(p.one))
			}
		}
	}
}
