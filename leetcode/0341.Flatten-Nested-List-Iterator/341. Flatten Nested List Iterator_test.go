package leetcode

import (
	"testing"
)

func Benchmark_Problem338(b *testing.B) {
	obj := Constructor([]*NestedInteger{})

	obj.Next()
	obj.HasNext()
}
