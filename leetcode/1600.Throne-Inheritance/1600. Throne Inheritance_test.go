package leetcode

import (
	"testing"
)

func Benchmark_Problem1600(b *testing.B) {
	obj := Constructor("king")

	obj.Birth("king", "andy")

	obj.Birth("king", "bob")

	obj.Birth("king", "catherine")

	obj.Birth("andy", "matthew")

	obj.Birth("bob", "alex")

	obj.Birth("bob", "asha")

	obj.Death("bob")

}
