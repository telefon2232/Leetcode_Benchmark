package leetcode

import (
	"fmt"
	"testing"
)

type question959 struct {
	para959
	ans959
}

// para 是参数
// one 代表第一个参数
type para959 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans959 struct {
	one int
}

func Benchmark_Problem959(b *testing.B) {

	qs := []question959{
		{
			para959{[]string{" /", "/ "}},
			ans959{2},
		},

		{
			para959{[]string{" /", "  "}},
			ans959{1},
		},

		{
			para959{[]string{"\\/", "/\\"}},
			ans959{4},
		},

		{
			para959{[]string{"/\\", "\\/"}},
			ans959{5},
		},

		{
			para959{[]string{"//", "/ "}},
			ans959{3},
		},
	}


	for _, q := range qs {
		_, p := q.ans959, q.para959
		fmt.Printf("【input】:%v       【output】:%v\n", p, regionsBySlashes(p.one))
	}
}
