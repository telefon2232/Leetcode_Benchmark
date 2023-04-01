package leetcode

import (
	"fmt"
	"testing"
)

type question1446 struct {
	para1446
	ans1446
}

// para 是参数
type para1446 struct {
	s string
}

// ans 是答案
type ans1446 struct {
	ans int
}

func Benchmark_Problem1446(b *testing.B) {

	qs := []question1446{

		{
			para1446{"leetcode"},
			ans1446{2},
		},

		{
			para1446{"abbcccddddeeeeedcba"},
			ans1446{5},
		},

		{
			para1446{"triplepillooooow"},
			ans1446{5},
		},

		{
			para1446{"hooraaaaaaaaaaay"},
			ans1446{11},
		},

		{
			para1446{"tourist"},
			ans1446{1},
		},
	}


	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1446, q.para1446
		fmt.Printf("【input】:%v    【output】:%v\n", p.s, maxPower(p.s))
	}
}}}
