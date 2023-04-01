package leetcode

import (
	"testing"
)

type question752 struct {
	para752
	ans752
}

// para 是参数
// one 代表第一个参数
type para752 struct {
	deadends []string
	target   string
}

// ans 是答案
// one 代表第一个答案
type ans752 struct {
	one int
}

func Benchmark_Problem752(b *testing.B) {

	qs := []question752{

		{
			para752{[]string{"0201", "0101", "0102", "1212", "2002"}, "0202"},
			ans752{6},
		},

		{
			para752{[]string{"8888"}, "0009"},
			ans752{1},
		},

		{
			para752{[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"},
			ans752{-1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans752, q.para752
		(openLock(p.deadends, p.target))
	}
}}}
