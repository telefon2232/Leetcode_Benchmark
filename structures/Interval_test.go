package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_Interval2Ints(b *testing.B) {
	ast := assert.New(t)

	actual := Interval2Ints(Interval{Start: 1, End: 2})
	expected := []int{1, 2}
	ast.Equal(expected, actual)
}

func Benchmark_IntervalSlice2Intss(b *testing.B) {
	ast := assert.New(t)

	actual := IntervalSlice2Intss(
		[]Interval{
			{
				Start: 1,
				End:   2,
			},
			{
				Start: 3,
				End:   4,
			},
		},
	)
	expected := [][]int{
		{1, 2},
		{3, 4},
	}

	ast.Equal(expected, actual)
}
func Benchmark_Intss2IntervalSlice(b *testing.B) {
	ast := assert.New(t)

	expected := []Interval{
		{
			Start: 1,
			End:   2,
		},
		{
			Start: 3,
			End:   4,
		},
	}
	actual := Intss2IntervalSlice([][]int{
		{1, 2},
		{3, 4},
	},
	)

	ast.Equal(expected, actual)
}
