package leetcode

import (
	"testing"
)

type question14 struct {
	para14
	ans14
}

// para 是参数
type para14 struct {
	strs []string
}

// ans 是答案
type ans14 struct {
	ans string
}

func Benchmark_Problem14(b *testing.B) {

	qs := []question14{

		{
			para14{[]string{"flower", "flow", "flight"}},
			ans14{"fl"},
		},

		{
			para14{[]string{"dog", "racecar", "car"}},
			ans14{""},
		},

		{
			para14{[]string{"ab", "a"}},
			ans14{"a"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans14, q.para14
				(longestCommonPrefix(p.strs))
			}
		}
	}
}
