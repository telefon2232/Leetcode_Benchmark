package leetcode

import (
	"testing"
)

type question726 struct {
	para726
	ans726
}

// para 是参数
// one 代表第一个参数
type para726 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans726 struct {
	one string
}

func Benchmark_Problem726(b *testing.B) {

	qs := []question726{

		{
			para726{"H200P"},
			ans726{"H200P"},
		},

		{
			para726{"H2O"},
			ans726{"H2O"},
		},

		{
			para726{"Mg(OH)2"},
			ans726{"H2MgO2"},
		},

		{
			para726{"K4(ON(SO3)2)2"},
			ans726{"K4N2O14S4"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans726, q.para726
				(countOfAtoms(p.one))
			}
		}
	}
}
