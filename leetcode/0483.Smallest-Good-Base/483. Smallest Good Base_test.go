package leetcode

import (
	"testing"
)

type question483 struct {
	para483
	ans483
}

// para 是参数
// one 代表第一个参数
type para483 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans483 struct {
	one string
}

func Benchmark_Problem483(b *testing.B) {

	qs := []question483{

		{
			para483{"13"},
			ans483{"3"},
		},

		{
			para483{"4681"},
			ans483{"8"},
		},

		{
			para483{"1000000000000000000"},
			ans483{"999999999999999999"},
		},

		{
			para483{"727004545306745403"},
			ans483{"727004545306745402"},
		},
	}

	for _, q := range qs {
		_, p := q.ans483, q.para483
		(smallestGoodBase(p.one))
	}
}
