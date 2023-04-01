package leetcode

import (
	"testing"
)

type question1268 struct {
	para1268
	ans1268
}

// para 是参数
// one 代表第一个参数
type para1268 struct {
	products   []string
	searchWord string
}

// ans 是答案
// one 代表第一个答案
type ans1268 struct {
	one [][]string
}

func Benchmark_Problem1268(b *testing.B) {

	qs := []question1268{

		{
			para1268{[]string{"bags", "baggage", "banner", "box", "cloths"}, "bags"},
			ans1268{[][]string{
				{"baggage", "bags", "banner"}, {"baggage", "bags", "banner"}, {"baggage", "bags"}, {"bags"},
			}},
		},

		{
			para1268{[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"},
			ans1268{[][]string{
				{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
			}},
		},

		{
			para1268{[]string{"havana"}, "havana"},
			ans1268{[][]string{
				{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"},
			}},
		},

		{
			para1268{[]string{"havana"}, "tatiana"},
			ans1268{[][]string{
				{}, {}, {}, {}, {}, {}, {},
			}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ { 
for _, q := range qs { {
		_, p := q.ans1268, q.para1268
		(suggestedProducts(p.products, p.searchWord))
	}
}}}
