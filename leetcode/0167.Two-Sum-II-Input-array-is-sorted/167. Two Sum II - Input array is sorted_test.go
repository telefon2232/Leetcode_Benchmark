package leetcode

import (
	"testing"
)

type question167 struct {
	para167
	ans167
}

// para 是参数
// one 代表第一个参数
type para167 struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans167 struct {
	one []int
}

func Benchmark_Problem167(b *testing.B) {

	qs := []question167{

		{
			para167{[]int{2, 7, 11, 15}, 9},
			ans167{[]int{1, 2}},
		},
	}

	for _, q := range qs {
		_, p := q.ans167, q.para167
		(twoSum167(p.one, p.two))
	}
}
