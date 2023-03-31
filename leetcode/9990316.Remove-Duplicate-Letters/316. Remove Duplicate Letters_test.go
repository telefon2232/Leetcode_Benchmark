package leetcode

import (
	"fmt"
	"testing"
)

type question316 struct {
	para316
	ans316
}

// para 是参数
// one 代表第一个参数
type para316 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans316 struct {
	one string
}

func Benchmark_Problem316(b *testing.B) {

	qs := []question316{

		{
			para316{"bcabc"},
			ans316{"abc"},
		},
		{
			para316{"cbacdcbc"},
			ans316{"acdb"},
		},
	}


	for _, q := range qs {
		_, p := q.ans316, q.para316
		fmt.Printf("【input】:%v       【output】:%v\n", p, removeDuplicateLetters(p.one))
	}
}
