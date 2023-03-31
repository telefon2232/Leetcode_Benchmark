package leetcode

import (
	"fmt"
	"testing"
)

type question434 struct {
	para434
	ans434
}

// s 是参数
type para434 struct {
	s string
}

// ans 是答案
type ans434 struct {
	ans int
}

func Benchmark_Problem434(b *testing.B) {

	qs := []question434{

		{
			para434{"Hello, my name is John"},
			ans434{5},
		},

		{
			para434{"Hello"},
			ans434{1},
		},

		{
			para434{"love live! mu'sic forever"},
			ans434{4},
		},

		{
			para434{""},
			ans434{0},
		},
	}


	for _, q := range qs {
		_, p := q.ans434, q.para434
		fmt.Printf("【input】:%v       【output】:%v\n", p, countSegments(p.s))
	}
}
