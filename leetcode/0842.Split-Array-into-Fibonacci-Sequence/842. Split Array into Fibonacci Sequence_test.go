package leetcode

import (
	"testing"
)

type question842 struct {
	para842
	ans842
}

// para 是参数
// one 代表第一个参数
type para842 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans842 struct {
	one []int
}

func Benchmark_Problem842(b *testing.B) {

	qs := []question842{

		{
			para842{"11235813"},
			ans842{[]int{1, 1, 2, 3, 5, 8, 13}},
		},

		{
			para842{"123456579"},
			ans842{[]int{123, 456, 579}},
		},

		{
			para842{"112358130"},
			ans842{[]int{}},
		},

		{
			para842{"0123"},
			ans842{[]int{}},
		},

		{
			para842{"1101111"},
			ans842{[]int{110, 1, 111}},
		},

		{
			para842{"539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511"},
			ans842{[]int{}},
		},

		{
			para842{"3611537383985343591834441270352104793375145479938855071433500231900737525076071514982402115895535257195564161509167334647108949738176284385285234123461518508746752631120827113919550237703163294909"},
			ans842{[]int{}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans842, q.para842
				(splitIntoFibonacci(p.one))
			}
		}
	}
}
