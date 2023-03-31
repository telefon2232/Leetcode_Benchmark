package leetcode

import (
	"fmt"
	"testing"
)

type question839 struct {
	para839
	ans839
}

// para 是参数
// one 代表第一个参数
type para839 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans839 struct {
	one int
}

func Benchmark_Problem839(b *testing.B) {

	qs := []question839{

		{
			para839{[]string{"tars", "rats", "arts", "star"}},
			ans839{2},
		},
	}


	for _, q := range qs {
		_, p := q.ans839, q.para839
		fmt.Printf("【input】:%v       【output】:%v\n", p, numSimilarGroups(p.one))
	}
}
