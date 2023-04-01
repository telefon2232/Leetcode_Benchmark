package leetcode

import (
	"testing"
)

type question386 struct {
	para386
	ans386
}

// para 是参数
// one 代表第一个参数
type para386 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans386 struct {
	one []int
}

func Benchmark_Problem386(b *testing.B) {

	qs := []question386{

		{
			para386{13},
			ans386{[]int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans386, q.para386
		(lexicalOrder(p.n))
	}
}}}
