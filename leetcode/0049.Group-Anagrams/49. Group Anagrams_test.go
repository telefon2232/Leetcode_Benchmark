package leetcode

import (
	"testing"
)

type question49 struct {
	para49
	ans49
}

// para 是参数
// one 代表第一个参数
type para49 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans49 struct {
	one [][]string
}

func Benchmark_Problem49(b *testing.B) {

	qs := []question49{

		{
			para49{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			ans49{[][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans49, q.para49
				(groupAnagrams(p.one))
			}
		}
	}
}
