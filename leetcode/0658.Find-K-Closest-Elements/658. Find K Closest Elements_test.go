package leetcode

import (
	"fmt"
	"testing"
)

type question658 struct {
	para658
	ans658
}

// para 是参数
// one 代表第一个参数
type para658 struct {
	arr []int
	k   int
	x   int
}

// ans 是答案
// one 代表第一个答案
type ans658 struct {
	one []int
}

func Benchmark_Problem658(b *testing.B) {

	qs := []question658{

		{
			para658{[]int{1, 2, 3, 4, 5}, 4, 3},
			ans658{[]int{1, 2, 3, 4}},
		},

		{
			para658{[]int{1, 2, 3, 4, 5}, 4, -1},
			ans658{[]int{1, 2, 3, 4}},
		},
	}


	for _, q := range qs {
		_, p := q.ans658, q.para658
		fmt.Printf("【input】:%v       【output】:%v\n", p, findClosestElements(p.arr, p.k, p.x))
	}
}
