package leetcode

import (
	"testing"
)

type question630 struct {
	para630
	ans630
}

// para 是参数
// one 代表第一个参数
type para630 struct {
	courses [][]int
}

// ans 是答案
// one 代表第一个答案
type ans630 struct {
	one int
}

func Benchmark_Problem630(b *testing.B) {

	qs := []question630{

		{
			para630{[][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}},
			ans630{3},
		},
	}

	for _, q := range qs {
		_, p := q.ans630, q.para630
		(scheduleCourse(p.courses))
	}
}
