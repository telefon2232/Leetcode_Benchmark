package leetcode

import (
	"testing"
)

type question864 struct {
	para864
	ans864
}

// para 是参数
// one 代表第一个参数
type para864 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans864 struct {
	one int
}

func Benchmark_Problem864(b *testing.B) {

	qs := []question864{

		{
			para864{[]string{".##..##..B", "##...#...#", "..##..#...", ".#..#b...#", "#.##.a.###", ".#....#...", ".##..#.#..", ".....###@.", "..........", ".........A"}},
			ans864{11},
		},

		{
			para864{[]string{"Dd#b@", ".fE.e", "##.B.", "#.cA.", "aF.#C"}},
			ans864{14},
		},

		{
			para864{[]string{"@...a", ".###A", "b.BCc"}},
			ans864{10},
		},

		{
			para864{[]string{"@Aa"}},
			ans864{-1},
		},

		{
			para864{[]string{"b", "A", "a", "@", "B"}},
			ans864{3},
		},

		{
			para864{[]string{"@.a.#", "#####", "b.A.B"}},
			ans864{-1},
		},

		{
			para864{[]string{"@.a.#", "###.#", "b.A.B"}},
			ans864{8},
		},

		{
			para864{[]string{"@..aA", "..B#.", "....b"}},
			ans864{6},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans864, q.para864
				(shortestPathAllKeys(p.one))
			}
		}
	}
}
