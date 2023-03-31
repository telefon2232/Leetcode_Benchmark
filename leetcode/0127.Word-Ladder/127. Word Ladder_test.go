package leetcode

import (
	"fmt"
	"testing"
)

type question127 struct {
	para127
	ans127
}

// para 是参数
// one 代表第一个参数
type para127 struct {
	b string
	e string
	w []string
}

// ans 是答案
// one 代表第一个答案
type ans127 struct {
	one int
}

func Benchmark_Problem127(b *testing.B) {

	qs := []question127{
		{
			para127{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}},
			ans127{5},
		},

		{
			para127{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}},
			ans127{0},
		},
	}


	for _, q := range qs {
		_, p := q.ans127, q.para127
		fmt.Printf("【input】:%v       【output】:%v\n", p, ladderLength(p.b, p.e, p.w))
	}
}
