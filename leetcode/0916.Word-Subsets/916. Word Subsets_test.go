package leetcode

import (
	"testing"
)

type question916 struct {
	para916
	ans916
}

// para 是参数
// one 代表第一个参数
type para916 struct {
	A []string
	B []string
}

// ans 是答案
// one 代表第一个答案
type ans916 struct {
	one []string
}

func Benchmark_Problem916(b *testing.B) {

	qs := []question916{

		{
			para916{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}},
			ans916{[]string{"facebook", "google", "leetcode"}},
		},

		{
			para916{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"}},
			ans916{[]string{"apple", "google", "leetcode"}},
		},

		{
			para916{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "oo"}},
			ans916{[]string{"facebook", "google"}},
		},

		{
			para916{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"lo", "eo"}},
			ans916{[]string{"google", "leetcode"}},
		},

		{
			para916{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"ec", "oc", "ceo"}},
			ans916{[]string{"facebook", "leetcode"}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans916, q.para916
		(wordSubsets(p.A, p.B))
	}
}}}
