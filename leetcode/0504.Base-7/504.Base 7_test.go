package leetcode

import (
	"fmt"
	"testing"
)

type question504 struct {
	para504
	ans504
}

// para 是参数
type para504 struct {
	num int
}

// ans 是答案
type ans504 struct {
	ans string
}

func Benchmark_Problem504(b *testing.B) {

	qs := []question504{

		{
			para504{100},
			ans504{"202"},
		},

		{
			para504{-7},
			ans504{"-10"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans504, q.para504
				fmt.Printf("【input】:%v      ", p.num)
				(convertToBase7(p.num))
			}
		}
	}
}
