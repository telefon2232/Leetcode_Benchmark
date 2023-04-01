package leetcode

import (
	"testing"
)

type question1332 struct {
	para1332
	ans1332
}

// para 是参数
// one 代表第一个参数
type para1332 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1332 struct {
	one int
}

func Benchmark_Problem1332(b *testing.B) {

	qs := []question1332{

		{
			para1332{"ababa"},
			ans1332{1},
		},

		{
			para1332{"abb"},
			ans1332{2},
		},

		{
			para1332{"baabb"},
			ans1332{2},
		},

		{
			para1332{""},
			ans1332{0},
		},

		{
			para1332{"bbaabaaa"},
			ans1332{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1332, q.para1332
		(removePalindromeSub(p.s))
	}
}}}
