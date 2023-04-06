package leetcode

import (
	"testing"
)

type question851 struct {
	para851
	ans851
}

// para 是参数
// one 代表第一个参数
type para851 struct {
	richer [][]int
	quiet  []int
}

// ans 是答案
// one 代表第一个答案
type ans851 struct {
	one []int
}

func Benchmark_Problem851(b *testing.B) {

	qs := []question851{

		{
			para851{[][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}}, []int{3, 2, 5, 4, 6, 1, 7, 0}},
			ans851{[]int{5, 5, 2, 5, 4, 5, 6, 7}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans851, q.para851
				(loudAndRich(p.richer, p.quiet))
			}
		}
	}
}
