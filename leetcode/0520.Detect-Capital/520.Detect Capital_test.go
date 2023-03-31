package leetcode

import (
	"fmt"
	"testing"
)

type question520 struct {
	para520
	ans520
}

// para 是参数
type para520 struct {
	word string
}

// ans 是答案
type ans520 struct {
	ans bool
}

func Benchmark_Problem520(b *testing.B) {

	qs := []question520{

		{
			para520{"USA"},
			ans520{true},
		},

		{
			para520{"FlaG"},
			ans520{false},
		},
	}


	for _, q := range qs {
		_, p := q.ans520, q.para520
		fmt.Printf("【input】:%v       【output】:%v\n", p, detectCapitalUse(p.word))
	}
}
