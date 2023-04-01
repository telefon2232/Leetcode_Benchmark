package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_NestedInteger(b *testing.B) {
	ast := assert.New(b)

	n := NestedInteger{}

	ast.True(n.IsInteger())

	n.SetInteger(1)
	ast.Equal(1, n.GetInteger())

	elem := NestedInteger{Num: 1}

	expected := NestedInteger{
		Num: 1,
		Ns:  []*NestedInteger{&elem},
	}
	n.Add(elem)

	ast.Equal(expected, n)

	ast.Equal(expected.Ns, n.GetList())
}
