package leetcode

import (
	"testing"
)

type question986 struct {
	para986
	ans986
}

// para 是参数
// one 代表第一个参数
type para986 struct {
	one []Interval
	two []Interval
}

// ans 是答案
// one 代表第一个答案
type ans986 struct {
	one []Interval
}

func Benchmark_Problem986(b *testing.B) {

	qs := []question986{

		{
			para986{[]Interval{{Start: 0, End: 2}, {Start: 5, End: 10}, {Start: 13, End: 23}, {Start: 24, End: 25}},
				[]Interval{{Start: 1, End: 5}, {Start: 8, End: 12}, {Start: 15, End: 24}, {Start: 25, End: 26}}},
			ans986{[]Interval{{Start: 1, End: 2}, {Start: 5, End: 5}, {Start: 8, End: 10},
				{Start: 15, End: 23}, {Start: 24, End: 24}, {Start: 25, End: 25}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans986, q.para986
				(intervalIntersection(p.one, p.two))
			}
		}
	}
}
