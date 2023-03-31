package leetcode

import (
	"fmt"
	"testing"
)

type question2182 struct {
	para2182
	ans2182
}

// para 是参数
// one 代表第一个参数
type para2182 struct {
	one   string
	limit int
}

// ans 是答案
// one 代表第一个答案
type ans2182 struct {
	one string
}

func Benchmark_Problem2182(b *testing.B) {

	qs := []question2182{

		{
			para2182{"cczazcc", 3},
			ans2182{"zzcccac"},
		},

		{
			para2182{"aababab", 2},
			ans2182{"bbabaa"},
		},
	}


	for _, q := range qs {
		_, p := q.ans2182, q.para2182
		fmt.Printf("【input】:%v       【output】:%v\n", p, repeatLimitedString(p.one, p.limit))
	}
}
