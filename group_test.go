package iter_lib_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	iter_lib "github.com/lezhnev74/iter"
)

func TestGroup(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		groupBy  func(int) int
		expected [][]int
	}{
		{
			name:     "empty sequence",
			input:    []int{},
			groupBy:  func(x int) int { return x },
			expected: [][]int(nil),
		},
		{
			name:     "single value",
			input:    []int{1},
			groupBy:  func(x int) int { return x },
			expected: [][]int{{1}},
		},
		{
			name:     "no groups",
			input:    []int{1, 2, 3},
			groupBy:  func(x int) int { return x },
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "consecutive groups",
			input:    []int{1, 1, 2, 2, 3, 3},
			groupBy:  func(x int) int { return x },
			expected: [][]int{{1, 1}, {2, 2}, {3, 3}},
		},
		{
			name:     "mixed groups",
			input:    []int{1, 1, 2, 3, 3, 2},
			groupBy:  func(x int) int { return x },
			expected: [][]int{{1, 1}, {2}, {3, 3}, {2}},
		},
		{
			name:     "group by modulo",
			input:    []int{1, 3, 5, 2, 4, 6},
			groupBy:  func(x int) int { return x % 2 },
			expected: [][]int{{1, 3, 5}, {2, 4, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				in := slices.Values(tt.input)
				out := iter_lib.Group(in, tt.groupBy)
				var results [][]int
				for group := range out {
					results = append(results, group)
				}
				require.Equal(t, tt.expected, results)
			},
		)
	}
}

func TestGroup2(t *testing.T) {
	tests := []struct {
		name     string
		input    []kv
		groupBy  func(string, int) string
		expected []kvGroup
	}{
		{
			name:     "empty sequence",
			input:    []kv{},
			groupBy:  func(k string, v int) string { return k },
			expected: []kvGroup(nil),
		},
		{
			name:     "single value",
			input:    []kv{{"a", 1}},
			groupBy:  func(k string, v int) string { return k },
			expected: []kvGroup{{"a", []kv{{"a", 1}}}},
		},
		{
			name:     "no groups",
			input:    []kv{{"a", 1}, {"b", 2}, {"c", 3}},
			groupBy:  func(k string, v int) string { return k },
			expected: []kvGroup{{"a", []kv{{"a", 1}}}, {"b", []kv{{"b", 2}}}, {"c", []kv{{"c", 3}}}},
		},
		{
			name:     "consecutive groups",
			input:    []kv{{"a", 1}, {"a", 2}, {"b", 1}, {"b", 2}},
			groupBy:  func(k string, v int) string { return k },
			expected: []kvGroup{{"a", []kv{{"a", 1}, {"a", 2}}}, {"b", []kv{{"b", 1}, {"b", 2}}}},
		},
		{
			name:    "mixed groups",
			input:   []kv{{"a", 1}, {"a", 2}, {"b", 1}, {"a", 3}},
			groupBy: func(k string, v int) string { return k },
			expected: []kvGroup{
				{"a", []kv{{"a", 1}, {"a", 2}}},
				{"b", []kv{{"b", 1}}},
				{"a", []kv{{"a", 3}}},
			},
		},
		{
			name:  "group by value parity",
			input: []kv{{"a", 1}, {"b", 3}, {"c", 2}, {"d", 4}},
			groupBy: func(k string, v int) string {
				if v%2 == 0 {
					return "even"
				}
				return "odd"
			},
			expected: []kvGroup{
				{"odd", []kv{{"a", 1}, {"b", 3}}},
				{"even", []kv{{"c", 2}, {"d", 4}}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				in := kvsToSeq(tt.input)
				out := iter_lib.Group2(in, tt.groupBy)
				var results []kvGroup
				for k, group := range out {
					var kvs []kv
					for _, item := range group {
						kvs = append(kvs, kv{item.K, item.V})
					}
					results = append(results, kvGroup{k, kvs})
				}
				require.Equal(t, tt.expected, results)
			},
		)
	}
}
