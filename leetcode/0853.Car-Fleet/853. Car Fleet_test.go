package leetcode

import (
	"fmt"
	"testing"
)

type question853 struct {
	para853
	ans853
}

// para 是参数
// one 代表第一个参数
type para853 struct {
	target   int
	position []int
	speed    []int
}

// ans 是答案
// one 代表第一个答案
type ans853 struct {
	one int
}

func Benchmark_Problem853(b *testing.B) {

	qs := []question853{

		{
			para853{12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}},
			ans853{3},
		},
	}


	for _, q := range qs {
		_, p := q.ans853, q.para853
		fmt.Printf("【input】:%v       【output】:%v\n", p, carFleet(p.target, p.position, p.speed))
	}
}
