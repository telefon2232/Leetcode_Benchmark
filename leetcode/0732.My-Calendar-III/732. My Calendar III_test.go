package leetcode

import (
	"testing"
)

func Benchmark_Problem732(b *testing.B) {
	obj := Constructor732()
	for bbe := 0; bbe < b.N; bbe++ {
		obj.Book(10, 20) // returns 1
		obj.Book(50, 60) // returns 1
		obj.Book(10, 40) // returns 2
		obj.Book(5, 15)  // returns 3
		obj.Book(5, 10)  // returns 3
		obj.Book(25, 55) // returns 3
	}
}
