package leetcode

import (
	"testing"
)

type question1189 struct {
	para1189
	ans1189
}

// para 是参数
// one 代表第一个参数
type para1189 struct {
	text string
}

// ans 是答案
// one 代表第一个答案
type ans1189 struct {
	one int
}

func Benchmark_Problem1189(b *testing.B) {

	qs := []question1189{

		{
			para1189{"nlaebolko"},
			ans1189{1},
		},

		{
			para1189{"loonbalxballpoon"},
			ans1189{2},
		},

		{
			para1189{"leetcode"},
			ans1189{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1189, q.para1189
		(maxNumberOfBalloons(p.text))
	}
}}}
