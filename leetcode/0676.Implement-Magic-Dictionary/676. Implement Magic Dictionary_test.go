package leetcode

import (
	"testing"
)

func Benchmark_Problem676(b *testing.B) {
	dict := []string{"hello", "leetcode"}
	obj := Constructor676()
	obj.BuildDict(dict)
	for bbe := 0; bbe < b.N; bbe++ {
		obj.Search("hello")
		obj.Search("apple")
		obj.Search("leetcode")
		obj.Search("leetcoded")
		obj.Search("hhllo")
		obj.Search("hell")
	}
}
