package leetcode

import (
	"testing"
)

type question1691 struct {
	para1691
	ans1691
}

// para 是参数
// one 代表第一个参数
type para1691 struct {
	cuboids [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1691 struct {
	one int
}

func Benchmark_Problem1691(b *testing.B) {

	qs := []question1691{

		{
			para1691{[][]int{{50, 45, 20}, {95, 37, 53}, {45, 23, 12}}},
			ans1691{190},
		},

		{
			para1691{[][]int{{38, 25, 45}, {76, 35, 3}}},
			ans1691{76},
		},

		{
			para1691{[][]int{{7, 11, 17}, {7, 17, 11}, {11, 7, 17}, {11, 17, 7}, {17, 7, 11}, {17, 11, 7}}},
			ans1691{102},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1691, q.para1691
		(maxHeight(p.cuboids))
	}
}}}
