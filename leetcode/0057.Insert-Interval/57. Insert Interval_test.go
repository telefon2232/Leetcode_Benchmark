package leetcode

import (
	"testing"
)

type question57 struct {
	para57
	ans57
}

// para 是参数
// one 代表第一个参数
type para57 struct {
	one []Interval
	two Interval
}

// ans 是答案
// one 代表第一个答案
type ans57 struct {
	one []Interval
}

func Benchmark_Problem57(b *testing.B) {

	qs := []question57{

		{
			para57{[]Interval{}, Interval{}},
			ans57{[]Interval{}},
		},

		{
			para57{[]Interval{{Start: 1, End: 3}, {Start: 6, End: 9}}, Interval{Start: 4, End: 8}},
			ans57{[]Interval{{Start: 1, End: 5}, {Start: 6, End: 9}}},
		},

		{
			para57{[]Interval{{Start: 1, End: 3}, {Start: 6, End: 9}}, Interval{Start: 2, End: 5}},
			ans57{[]Interval{{Start: 1, End: 5}, {Start: 6, End: 9}}},
		},

		{
			para57{[]Interval{{Start: 1, End: 2}, {Start: 3, End: 5}, {Start: 6, End: 7}, {Start: 8, End: 10}, {Start: 12, End: 16}}, Interval{Start: 4, End: 8}},
			ans57{[]Interval{{Start: 1, End: 2}, {Start: 3, End: 10}, {Start: 12, End: 16}}},
		},

		{
			para57{[]Interval{{Start: 1, End: 5}}, Interval{Start: 5, End: 7}},
			ans57{[]Interval{{Start: 1, End: 7}}},
		},

		{
			para57{[]Interval{{Start: 1, End: 2}, {Start: 3, End: 5}, {Start: 6, End: 7}, {Start: 8, End: 10}, {Start: 12, End: 16}}, Interval{Start: 9, End: 12}},
			ans57{[]Interval{{Start: 1, End: 2}, {Start: 3, End: 5}, {Start: 6, End: 7}, {Start: 8, End: 16}}},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans57, q.para57
				(insert(p.one, p.two))
			}
		}
	}
}
