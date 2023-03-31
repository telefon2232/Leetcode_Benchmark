package leetcode

import (
	"fmt"
	"testing"
)

type question409 struct {
	para409
	ans409
}

// para 是参数
// one 代表第一个参数
type para409 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans409 struct {
	one int
}

func Benchmark_Problem409(b *testing.B) {

	qs := []question409{

		{
			para409{"abccccdd"},
			ans409{7},
		},
	}


	for _, q := range qs {
		_, p := q.ans409, q.para409
		fmt.Printf("【input】:%v       【output】:%v\n", p, longestPalindrome(p.one))
	}
}
