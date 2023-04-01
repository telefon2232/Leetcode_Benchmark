package leetcode

import (
	"testing"
)

type question771 struct {
	para771
	ans771
}

// para 是参数
// one 代表第一个参数
type para771 struct {
	one string
	two string
}

// ans 是答案
// one 代表第一个答案
type ans771 struct {
	one int
}

func Benchmark_Problem771(b *testing.B) {

	qs := []question771{
		{
			para771{"aA", "aAAbbbb"},
			ans771{3},
		},

		{
			para771{"z", "ZZ"},
			ans771{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans771, q.para771
		(numJewelsInStones(p.one, p.two))
	}
}}}
