package leetcode

import (
	"fmt"
	"testing"
)

type question1705 struct {
	para1705
	ans1705
}

// para 是参数
type para1705 struct {
	apples []int
	days   []int
}

// ans 是答案
type ans1705 struct {
	ans int
}

func Benchmark_Problem1705(b *testing.B) {

	qs := []question1705{

		{
			para1705{[]int{1, 2, 3, 5, 2}, []int{3, 2, 1, 4, 2}},
			ans1705{7},
		},

		{
			para1705{[]int{3, 0, 0, 0, 0, 2}, []int{3, 0, 0, 0, 0, 2}},
			ans1705{5},
		},
	}


	for _, q := range qs {
		_, p := q.ans1705, q.para1705
		fmt.Printf("【input】:%v    【output】:%v\n", p, eatenApples(p.apples, p.days))
	}
}
