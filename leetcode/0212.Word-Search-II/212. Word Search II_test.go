package leetcode

import (
	"testing"
)

type question212 struct {
	para212
	ans212
}

// para 是参数
// one 代表第一个参数
type para212 struct {
	b    [][]byte
	word []string
}

// ans 是答案
// one 代表第一个答案
type ans212 struct {
	one []string
}

func Benchmark_Problem212(b *testing.B) {

	qs := []question212{

		{
			para212{[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			}, []string{"ABCCED", "SEE", "ABCB"}},
			ans212{[]string{"ABCCED", "SEE"}},
		},

		{
			para212{[][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			}, []string{"oath", "pea", "eat", "rain"}},
			ans212{[]string{"oath", "eat"}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans212, q.para212
				findWords(p.b, p.word)
			}
		}
	}
}
