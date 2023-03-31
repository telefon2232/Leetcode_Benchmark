package leetcode

import (
	"testing"
)

type question744 struct {
	para744
	ans744
}

// para 是参数
// one 代表第一个参数
type para744 struct {
	letters []byte
	target  byte
}

// ans 是答案
// one 代表第一个答案
type ans744 struct {
	one byte
}

func Benchmark_Problem744(b *testing.B) {

	qs := []question744{

		{
			para744{[]byte{'c', 'f', 'j'}, 'a'},
			ans744{'c'},
		},

		{
			para744{[]byte{'c', 'f', 'j'}, 'c'},
			ans744{'f'},
		},

		{
			para744{[]byte{'c', 'f', 'j'}, 'd'},
			ans744{'f'},
		},

		{
			para744{[]byte{'c', 'f', 'j'}, 'g'},
			ans744{'j'},
		},

		{
			para744{[]byte{'c', 'f', 'j'}, 'j'},
			ans744{'c'},
		},

		{
			para744{[]byte{'c', 'f', 'j'}, 'k'},
			ans744{'c'},
		},
	}

	for _, q := range qs {
		_, p := q.ans744, q.para744
		(nextGreatestLetter(p.letters, p.target))
	}
}
