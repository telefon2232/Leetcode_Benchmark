package leetcode

import (
	"testing"
)

type question438 struct {
	para438
	ans438
}

// para 是参数
// one 代表第一个参数
type para438 struct {
	s string
	p string
}

// ans 是答案
// one 代表第一个答案
type ans438 struct {
	one []int
}

func Benchmark_Problem438(b *testing.B) {

	qs := []question438{

		{
			para438{"abab", "ab"},
			ans438{[]int{0, 1, 2}},
		},

		{
			para438{"cbaebabacd", "abc"},
			ans438{[]int{0, 6}},
		},

		{
			para438{"", "abc"},
			ans438{[]int{}},
		},

		{
			para438{"abacbabc", "abc"},
			ans438{[]int{1, 2, 3, 5}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans438, q.para438
				(findAnagrams(p.s, p.p))
			}
		}
	}
}
