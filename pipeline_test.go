package iter_lib_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lezhnev74/iter_lib"
)

func TestPipeline1(t *testing.T) {
	var s1 iter.Seq[int] = slices.Values([]int{1, 2})
	var s2 iter.Seq[int] = slices.Values([]int{2, 3})

	merge := iter_lib.MergeOrdered(s1, s2, false)
	pipeline := iter_lib.Dedup(merge, func(i1, i2 int) bool { return i1 == i2 })

	result := slices.Collect(pipeline)
	require.Equal(t, []int{1, 2, 3}, result)
}
