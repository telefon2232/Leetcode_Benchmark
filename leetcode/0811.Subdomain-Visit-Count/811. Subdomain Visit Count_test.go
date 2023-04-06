package leetcode

import (
	"testing"
)

type question811 struct {
	para811
	ans811
}

// para 是参数
// one 代表第一个参数
type para811 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans811 struct {
	one []string
}

func Benchmark_Problem811(b *testing.B) {

	qs := []question811{

		{
			para811{[]string{"9001 discuss.leetcode.com"}},
			ans811{[]string{"mqe", "mqE", "mQe", "mQE", "Mqe", "MqE", "MQe", "MQE"}},
		},

		{
			para811{[]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}},
			ans811{[]string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans811, q.para811
				(subdomainVisits(p.one))
			}
		}
	}
}
