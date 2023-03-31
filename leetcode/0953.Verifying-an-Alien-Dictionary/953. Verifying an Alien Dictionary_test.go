package leetcode

import (
	"fmt"
	"testing"
)

type question953 struct {
	para953
	ans953
}

// para 是参数
// one 代表第一个参数
type para953 struct {
	one []string
	two string
}

// ans 是答案
// one 代表第一个答案
type ans953 struct {
	one bool
}

func Benchmark_Problem953(b *testing.B) {

	qs := []question953{
		{
			para953{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"},
			ans953{true},
		},

		{
			para953{[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"},
			ans953{false},
		},

		{
			para953{[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"},
			ans953{false},
		},
	}


	for _, q := range qs {
		_, p := q.ans953, q.para953
		fmt.Printf("【input】:%v       【output】:%v\n", p, isAlienSorted(p.one, p.two))
	}
}
