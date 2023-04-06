package leetcode

import (
	"testing"
)

type question1629 struct {
	para1629
	ans1629
}

// para 是参数
// one 代表第一个参数
type para1629 struct {
	releaseTimes []int
	keysPressed  string
}

// ans 是答案
// one 代表第一个答案
type ans1629 struct {
	one byte
}

func Benchmark_Problem1629(b *testing.B) {

	qs := []question1629{

		{
			para1629{[]int{9, 29, 49, 50}, "cbcd"},
			ans1629{'c'},
		},

		{
			para1629{[]int{12, 23, 36, 46, 62}, "spuda"},
			ans1629{'a'},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans1629, q.para1629
				(slowestKey(p.releaseTimes, p.keysPressed))
			}
		}
	}
}
