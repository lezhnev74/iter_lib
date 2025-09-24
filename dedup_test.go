package iter_lib_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	iter_lib "github.com/lezhnev74/iter"
)

func TestDedup(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		eq       func(int, int) bool
		expected []int
	}{
		{
			name:     "empty sequence",
			input:    []int{},
			eq:       func(a, b int) bool { return a == b },
			expected: []int(nil),
		},
		{
			name:     "single value",
			input:    []int{1},
			eq:       func(a, b int) bool { return a == b },
			expected: []int{1},
		},
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3},
			eq:       func(a, b int) bool { return a == b },
			expected: []int{1, 2, 3},
		},
		{
			name:     "consecutive duplicates",
			input:    []int{1, 1, 2, 2, 3, 3},
			eq:       func(a, b int) bool { return a == b },
			expected: []int{1, 2, 3},
		},
		{
			name:     "mixed duplicates",
			input:    []int{1, 1, 2, 3, 3, 2},
			eq:       func(a, b int) bool { return a == b },
			expected: []int{1, 2, 3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				in := slices.Values(tt.input)
				out := iter_lib.Dedup(in, tt.eq)
				results := slices.Collect(out)
				require.Equal(
					t, tt.expected, results,
					"Dedup(%v, %v) = %v, want %v",
					tt.input, tt.eq, results, tt.expected,
				)
			},
		)
	}
}

func TestDedup2(t *testing.T) {
	tests := []struct {
		name     string
		input    []kv
		eq       func(string, int, string, int) bool
		expected []kv
	}{
		{
			name:     "empty sequence",
			input:    []kv{},
			eq:       func(k1 string, v1 int, k2 string, v2 int) bool { return k1 == k2 && v1 == v2 },
			expected: []kv(nil),
		},
		{
			name:     "single pair",
			input:    []kv{{k: "a", v: 1}},
			eq:       func(k1 string, v1 int, k2 string, v2 int) bool { return k1 == k2 && v1 == v2 },
			expected: []kv{{k: "a", v: 1}},
		},
		{
			name:     "no duplicates",
			input:    []kv{{k: "a", v: 1}, {k: "b", v: 2}, {k: "c", v: 3}},
			eq:       func(k1 string, v1 int, k2 string, v2 int) bool { return k1 == k2 && v1 == v2 },
			expected: []kv{{k: "a", v: 1}, {k: "b", v: 2}, {k: "c", v: 3}},
		},
		{
			name:     "consecutive duplicates",
			input:    []kv{{k: "a", v: 1}, {k: "a", v: 1}, {k: "b", v: 2}, {k: "b", v: 2}, {k: "c", v: 3}, {k: "c", v: 3}},
			eq:       func(k1 string, v1 int, k2 string, v2 int) bool { return k1 == k2 && v1 == v2 },
			expected: []kv{{k: "a", v: 1}, {k: "b", v: 2}, {k: "c", v: 3}},
		},
		{
			name:     "mixed duplicates",
			input:    []kv{{k: "a", v: 1}, {k: "a", v: 1}, {k: "b", v: 2}, {k: "c", v: 3}, {k: "c", v: 3}, {k: "b", v: 2}},
			eq:       func(k1 string, v1 int, k2 string, v2 int) bool { return k1 == k2 && v1 == v2 },
			expected: []kv{{k: "a", v: 1}, {k: "b", v: 2}, {k: "c", v: 3}, {k: "b", v: 2}},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				in := kvsToSeq(tt.input)
				out := iter_lib.Dedup2(in, tt.eq)
				results := seq2ToKvs(out)
				require.Equal(t, tt.expected, results)
			},
		)
	}
}
