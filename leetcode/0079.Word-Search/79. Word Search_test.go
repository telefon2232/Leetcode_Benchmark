package leetcode

import (
	"testing"
)

type question79 struct {
	para79
	ans79
}

// para 是参数
// one 代表第一个参数
type para79 struct {
	b    [][]byte
	word string
}

// ans 是答案
// one 代表第一个答案
type ans79 struct {
	one bool
}

func Benchmark_Problem79(b *testing.B) {

	qs := []question79{

		{
			para79{[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			}, "ABCCED"},
			ans79{true},
		},

		{
			para79{[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			}, "SEE"},
			ans79{true},
		},

		{
			para79{[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			}, "ABCB"},
			ans79{false},
		},

		{
			para79{[][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			}, "oath"},
			ans79{true},
		},

		{
			para79{[][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			}, "pea"},
			ans79{false},
		},

		{
			para79{[][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			}, "eat"},
			ans79{true},
		},

		{
			para79{[][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			}, "rain"},
			ans79{false},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans79, q.para79
				(exist(p.b, p.word))
			}
		}
	}
}
