package leetcode

import (
	"testing"
)

type question1455 struct {
	para1455
	ans1455
}

// para 是参数
// one 代表第一个参数
type para1455 struct {
	sentence   string
	searchWord string
}

// ans 是答案
// one 代表第一个答案
type ans1455 struct {
	one int
}

func Benchmark_Problem1455(b *testing.B) {

	qs := []question1455{

		{
			para1455{"i love eating burger", "burg"},
			ans1455{4},
		},

		{
			para1455{"this problem is an easy problem", "pro"},
			ans1455{2},
		},

		{
			para1455{"i am tired", "you"},
			ans1455{-1},
		},

		{
			para1455{"i use triple pillow", "pill"},
			ans1455{4},
		},

		{
			para1455{"hello from the other side", "they"},
			ans1455{-1},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1455, q.para1455

				(isPrefixOfWord(p.sentence, p.searchWord))
			}
		}
	}
}
