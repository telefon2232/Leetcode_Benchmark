package leetcode

import (
	"fmt"
	"testing"
)

type question17 struct {
	para17
	ans17
}

// para 是参数
// one 代表第一个参数
type para17 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans17 struct {
	one []string
}

func Benchmark_Problem17(b *testing.B) {

	qs := []question17{

		{
			para17{"23"},
			ans17{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		},
	}


	for _, q := range qs {
		_, p := q.ans17, q.para17
		fmt.Printf("【input】:%v       【output】:%v\n", p, letterCombinations(p.s))
	}
}
