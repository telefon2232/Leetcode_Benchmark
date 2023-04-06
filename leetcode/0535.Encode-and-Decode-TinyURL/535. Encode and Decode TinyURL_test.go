package leetcode

import (
	"testing"
)

func Benchmark_Problem535(b *testing.B) {
	obj := Constructor()

	e := obj.encode("https://leetcode.com/problems/design-tinyurl")

	obj.decode(e)

}
