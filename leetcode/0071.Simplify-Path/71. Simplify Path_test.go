package leetcode

import (
	"testing"
)

type question71 struct {
	para71
	ans71
}

// para 是参数
// one 代表第一个参数
type para71 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans71 struct {
	one string
}

func Benchmark_Problem71(b *testing.B) {

	qs := []question71{

		{
			para71{"/.hidden"},
			ans71{"/.hidden"},
		},

		{
			para71{"/..hidden"},
			ans71{"/..hidden"},
		},

		{
			para71{"/abc/..."},
			ans71{"/abc/..."},
		},

		{
			para71{"/home/"},
			ans71{"/home"},
		},

		{
			para71{"/..."},
			ans71{"/..."},
		},

		{
			para71{"/../"},
			ans71{"/"},
		},

		{
			para71{"/home//foo/"},
			ans71{"/home/foo"},
		},

		{
			para71{"/a/./b/../../c/"},
			ans71{"/c"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans71, q.para71
		(simplifyPath(p.s))
	}
}}}
