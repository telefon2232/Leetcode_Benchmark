package leetcode

import (
	"testing"
)

type question529 struct {
	para529
	ans529
}

// para 是参数
// one 代表第一个参数
type para529 struct {
	b     [][]byte
	click []int
}

// ans 是答案
// one 代表第一个答案
type ans529 struct {
	one [][]byte
}

func Benchmark_Problem529(b *testing.B) {

	qs := []question529{

		{
			para529{[][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'M', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			}, []int{3, 0}},
			ans529{[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			}},
		},

		{
			para529{[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			}, []int{1, 2}},
			ans529{[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'X', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans529, q.para529
				(updateBoard(p.b, p.click))
			}
		}
	}
}
