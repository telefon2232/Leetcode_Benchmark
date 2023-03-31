package leetcode

import (
	"fmt"
	"testing"
)

type question344 struct {
	para344
	ans344
}

// para 是参数
// one 代表第一个参数
type para344 struct {
	one []byte
}

// ans 是答案
// one 代表第一个答案
type ans344 struct {
	one []byte
}

func Benchmark_Problem344(b *testing.B) {

	qs := []question344{

		{
			para344{[]byte{'h', 'e', 'l', 'l', 'o'}},
			ans344{[]byte{'o', 'l', 'l', 'e', 'h'}},
		},

		{
			para344{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}},
			ans344{[]byte{'h', 'a', 'n', 'n', 'a', 'H'}},
		},
	}


	for _, q := range qs {
		_, p := q.ans344, q.para344
		fmt.Printf("【input】:%v      ", p.one)
		reverseString(p.one)
		fmt.Printf("【output】:%v\n", p.one)
	}
}
