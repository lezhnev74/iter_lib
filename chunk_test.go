package iter_lib_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lezhnev74/iter_lib"
)

func TestChunk(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		size     int
		expected [][]int
	}{
		{
			name:     "empty sequence",
			input:    []int{},
			size:     2,
			expected: [][]int(nil),
		},
		{
			name:     "single value",
			input:    []int{1},
			size:     2,
			expected: [][]int{{1}},
		},
		{
			name:     "exact chunks",
			input:    []int{1, 2, 3, 4},
			size:     2,
			expected: [][]int{{1, 2}, {3, 4}},
		},
		{
			name:     "partial last chunk",
			input:    []int{1, 2, 3, 4, 5},
			size:     2,
			expected: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name:     "chunk size equals length",
			input:    []int{1, 2, 3},
			size:     3,
			expected: [][]int{{1, 2, 3}},
		},
		{
			name:     "chunk size larger than input",
			input:    []int{1, 2},
			size:     3,
			expected: [][]int{{1, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				in := slices.Values(tt.input)
				out := iter_lib.Chunk(in, tt.size)
				var results [][]int
				for chunk := range out {
					results = append(results, chunk)
				}
				require.Equal(t, tt.expected, results)
			},
		)
	}
}

func TestChunk2(t *testing.T) {
	tests := []struct {
		name     string
		input    []kv
		size     int
		expected [][]kv
	}{
		{
			name:     "empty sequence",
			input:    []kv{},
			size:     2,
			expected: [][]kv(nil),
		},
		{
			name:     "single value",
			input:    []kv{{"a", 1}},
			size:     2,
			expected: [][]kv{{{"a", 1}}},
		},
		{
			name:     "exact chunks",
			input:    []kv{{"a", 1}, {"b", 2}, {"c", 3}, {"d", 4}},
			size:     2,
			expected: [][]kv{{{"a", 1}, {"b", 2}}, {{"c", 3}, {"d", 4}}},
		},
		{
			name:     "partial last chunk",
			input:    []kv{{"a", 1}, {"b", 2}, {"c", 3}, {"d", 4}, {"e", 5}},
			size:     2,
			expected: [][]kv{{{"a", 1}, {"b", 2}}, {{"c", 3}, {"d", 4}}, {{"e", 5}}},
		},
		{
			name:     "chunk size equals length",
			input:    []kv{{"a", 1}, {"b", 2}, {"c", 3}},
			size:     3,
			expected: [][]kv{{{"a", 1}, {"b", 2}, {"c", 3}}},
		},
		{
			name:     "chunk size larger than input",
			input:    []kv{{"a", 1}, {"b", 2}},
			size:     3,
			expected: [][]kv{{{"a", 1}, {"b", 2}}},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				in := kvsToSeq(tt.input)
				out := iter_lib.Chunk2(in, tt.size)
				var results [][]kv
				for chunk := range out {
					var kvChunk []kv
					for _, _kv := range chunk {
						kvChunk = append(kvChunk, kv{_kv.K, _kv.V})
					}
					results = append(results, kvChunk)
				}
				require.Equal(t, tt.expected, results)
			},
		)
	}
}
