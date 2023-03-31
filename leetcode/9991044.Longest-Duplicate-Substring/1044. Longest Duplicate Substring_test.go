package leetcode

import (
	"fmt"
	"testing"
)

type question1044 struct {
	para1044
	ans1044
}

// para 是参数
// one 代表第一个参数
type para1044 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans1044 struct {
	one string
}

func Benchmark_Problem1044(b *testing.B) {

	qs := []question1044{
		{
			para1044{"banana"},
			ans1044{"ana"},
		},

		{
			para1044{"abcd"},
			ans1044{""},
		},
	}


	for _, q := range qs {
		_, p := q.ans1044, q.para1044
		fmt.Printf("【input】:%v       【output】:%v\n", p, longestDupSubstring(p.one))
	}
}
