package leetcode

import (
	"testing"
)

type question130 struct {
	para130
	ans130
}

// para 是参数
// one 代表第一个参数
type para130 struct {
	one [][]byte
}

// ans 是答案
// one 代表第一个答案
type ans130 struct {
	one [][]byte
}

func Benchmark_Problem130(b *testing.B) {

	qs := []question130{

		{
			para130{[][]byte{}},
			ans130{[][]byte{}},
		},

		{
			para130{[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}},
			ans130{[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans130, q.para130

				solve1(p.one)

			}
		}
	}
}
