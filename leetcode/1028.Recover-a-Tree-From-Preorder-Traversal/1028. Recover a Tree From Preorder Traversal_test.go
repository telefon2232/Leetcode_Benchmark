package leetcode

import (
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question1028 struct {
	para1028
	ans1028
}

// para 是参数
// one 代表第一个参数
type para1028 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans1028 struct {
	one []int
}

func Benchmark_Problem1028(b *testing.B) {

	qs := []question1028{
		{
			para1028{"1-2--3--4-5--6--7"},
			ans1028{[]int{1, 2, 5, 3, 4, 6, 7}},
		},

		{
			para1028{"1-2--3---4-5--6---7"},
			ans1028{[]int{1, 2, 5, 3, structures.NULL, 6, structures.NULL, 4, structures.NULL, 7}},
		},

		{
			para1028{"1-401--349---90--88"},
			ans1028{[]int{1, 401, structures.NULL, 349, 88, 90}},
		},
	}

	for _, q := range qs {
		_, p := q.ans1028, q.para1028
		(structures.Tree2ints(recoverFromPreorder(p.one)))
	}
}
