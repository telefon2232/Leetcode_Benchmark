package leetcode

import (
	"testing"
)

type question748 struct {
	para748
	ans748
}

// para 是参数
// one 代表第一个参数
type para748 struct {
	c string
	w []string
}

// ans 是答案
// one 代表第一个答案
type ans748 struct {
	one string
}

func Benchmark_Problem748(b *testing.B) {

	qs := []question748{

		{
			para748{"1s3 PSt", []string{"step", "steps", "stripe", "stepple"}},
			ans748{"steps"},
		},

		{
			para748{"1s3 456", []string{"looks", "pest", "stew", "show"}},
			ans748{"pest"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans748, q.para748
		(shortestCompletingWord(p.c, p.w))
	}
}}}
