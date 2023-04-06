package leetcode

import (
	"testing"
)

type question447 struct {
	para447
	ans447
}

// para 是参数
// one 代表第一个参数
type para447 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans447 struct {
	one int
}

func Benchmark_Problem447(b *testing.B) {

	qs := []question447{

		{
			para447{[][]int{{0, 0}, {1, 0}, {2, 0}}},
			ans447{2},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans447, q.para447
				(numberOfBoomerangs(p.one))
			}
		}
	}
}
