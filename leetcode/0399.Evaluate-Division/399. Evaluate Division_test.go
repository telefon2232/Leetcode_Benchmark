package leetcode

import (
	"testing"
)

type question399 struct {
	para399
	ans399
}

// para 是参数
// one 代表第一个参数
type para399 struct {
	e [][]string
	v []float64
	q [][]string
}

// ans 是答案
// one 代表第一个答案
type ans399 struct {
	one []float64
}

func Benchmark_Problem399(b *testing.B) {

	qs := []question399{

		{
			para399{[][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}},
			ans399{[]float64{6.0, 0.5, -1.0, 1.0, -1.0}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans399, q.para399
		(calcEquation(p.e, p.v, p.q))
	}
}}}
