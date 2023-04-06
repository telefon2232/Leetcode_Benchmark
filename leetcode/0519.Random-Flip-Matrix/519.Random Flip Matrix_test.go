package leetcode

import (
	"fmt"
	"testing"
)

type question519 struct {
	para519
	ans519
}

// para 是参数
type para519 struct {
	para []string
	val  [][]int
}

// ans 是答案
type ans519 struct {
	ans [][]int
}

func Benchmark_Problem519(b *testing.B) {

	qs := []question519{

		{
			para519{[]string{"Solution", "flip", "flip", "flip", "reset", "flip"}, [][]int{{3, 1}, {}, {}, {}, {}, {}}},
			ans519{[][]int{nil, {1, 0}, {2, 0}, {0, 0}, nil, {2, 0}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans519, q.para519
				sol := Constructor(0, 0)
				for _, v := range p.para {
					if v == "Solution" {
						sol = Constructor(q.val[0][0], q.val[0][1])
						fmt.Printf("【input】:%v    【output】:%v\n", v, nil)
					} else if v == "flip" {
						fmt.Printf("【input】:%v    【output】:%v\n", v, sol.Flip())
					} else {
						sol.Reset()
						fmt.Printf("【input】:%v    【output】:%v\n", v, nil)
					}
				}
			}
		}
	}
}
