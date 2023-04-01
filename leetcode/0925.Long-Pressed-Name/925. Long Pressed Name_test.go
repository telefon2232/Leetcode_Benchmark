package leetcode

import (
	"testing"
)

type question925 struct {
	para925
	ans925
}

// para 是参数
// one 代表第一个参数
type para925 struct {
	name  string
	typed string
}

// ans 是答案
// one 代表第一个答案
type ans925 struct {
	one bool
}

func Benchmark_Problem925(b *testing.B) {

	qs := []question925{

		{
			para925{"alex", "aaleex"},
			ans925{true},
		},

		{
			para925{"alex", "alexxr"},
			ans925{false},
		},

		{
			para925{"alex", "alexxxxr"},
			ans925{false},
		},

		{
			para925{"alex", "alexxxxx"},
			ans925{true},
		},

		{
			para925{"saeed", "ssaaedd"},
			ans925{false},
		},

		{
			para925{"leelee", "lleeelee"},
			ans925{true},
		},

		{
			para925{"laiden", "laiden"},
			ans925{true},
		},

		{
			para925{"kikcxmvzi", "kiikcxxmmvvzz"},
			ans925{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans925, q.para925
		(isLongPressedName(p.name, p.typed))
	}
}}}
