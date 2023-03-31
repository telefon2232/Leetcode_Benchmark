package leetcode

import (
	"fmt"
	"testing"
)

type question384 struct {
	para384
	ans384
}

// para 是参数
type para384 struct {
	ops   []string
	value [][]int
}

// ans 是答案
type ans384 struct {
	ans [][]int
}

func Benchmark_Problem384(b *testing.B) {

	qs := []question384{

		{
			para384{ops: []string{"Solution", "shuffle", "reset", "shuffle"}, value: [][]int{{1, 2, 3}, {}, {}, {}}},
			ans384{[][]int{nil, {3, 1, 2}, {1, 2, 3}, {1, 3, 2}}},
		},
	}


	for _, q := range qs {
		sol := Constructor(nil)
		_, p := q.ans384, q.para384
		for _, op := range p.ops {
			if op == "Solution" {
				sol = Constructor(q.value[0])
			} else if op == "reset" {
				fmt.Printf("【input】:%v    【output】:%v\n", op, sol.Reset())
			} else {
				fmt.Printf("【input】:%v    【output】:%v\n", op, sol.Shuffle())
			}
		}
	}
}
