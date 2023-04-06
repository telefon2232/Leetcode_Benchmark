package leetcode

import (
	"testing"
)

func Benchmark_Problem981(b *testing.B) {
	obj := Constructor981()
	obj.Set("foo", "bar", 1)
	obj.Get("foo", 1)
	obj.Get("foo", 3)
	obj.Set("foo", "bar2", 4)
	obj.Get("foo", 4)
	obj.Get("foo", 5)
}
