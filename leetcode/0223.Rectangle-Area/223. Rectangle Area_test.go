package leetcode

import (
	"testing"
)

type question223 struct {
	para223
	ans223
}

// para 是参数
// one 代表第一个参数
type para223 struct {
	A int
	B int
	C int
	D int
	E int
	F int
	G int
	H int
}

// ans 是答案
// one 代表第一个答案
type ans223 struct {
	one int
}

func Benchmark_Problem223(b *testing.B) {

	qs := []question223{

		{
			para223{-3, 0, 3, 4, 0, -1, 9, 2},
			ans223{45},
		},
	}

	for _, q := range qs {
		_, p := q.ans223, q.para223
		(computeArea(p.A, p.B, p.C, p.D, p.E, p.F, p.G, p.H))
	}
}
