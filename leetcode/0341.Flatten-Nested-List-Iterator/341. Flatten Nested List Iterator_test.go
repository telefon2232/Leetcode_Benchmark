package leetcode

import (
	"testing"
)

func Benchmark_Problem338(b *testing.B) {
	obj := Constructor([]*NestedInteger{})
	for bbe := 0; bbe < b.N; bbe++ {
		obj.Next()
		obj.HasNext()
	}
}
