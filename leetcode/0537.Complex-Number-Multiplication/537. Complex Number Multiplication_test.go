package leetcode

import (
	"testing"
)

type question537 struct {
	para537
	ans537
}

// para 是参数
// one 代表第一个参数
type para537 struct {
	a string
	b string
}

// ans 是答案
// one 代表第一个答案
type ans537 struct {
	one string
}

func Benchmark_Problem537(b *testing.B) {

	qs := []question537{

		{
			para537{"1+1i", "1+1i"},
			ans537{"0+2i"},
		},

		{
			para537{"1+-1i", "1+-1i"},
			ans537{"0+-2i"},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans537, q.para537
		(complexNumberMultiply(p.a, p.b))
	}
}}}
