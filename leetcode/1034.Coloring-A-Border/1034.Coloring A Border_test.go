package leetcode

import (
	"fmt"
	"testing"
)

type question1034 struct {
	para1034
	ans1034
}

// para 是参数
type para1034 struct {
	grid  [][]int
	row   int
	col   int
	color int
}

// ans 是答案
type ans1034 struct {
	ans [][]int
}

func Benchmark_Problem1034(b *testing.B) {

	qs := []question1034{

		{
			para1034{[][]int{{1, 1}, {1, 2}}, 0, 0, 3},
			ans1034{[][]int{{3, 3}, {3, 2}}},
		},

		{
			para1034{[][]int{{1, 2, 2}, {2, 3, 2}}, 0, 1, 3},
			ans1034{[][]int{{1, 3, 3}, {2, 3, 3}}},
		},

		{
			para1034{[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 1, 1, 2},
			ans1034{[][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}},
		},
	}


	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1034, q.para1034
		fmt.Printf("【input】:%v    【output】:%v\n", p, colorBorder(p.grid, p.row, p.col, p.color))
	}
}}}
