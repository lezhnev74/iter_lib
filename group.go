package iter_lib

import "iter"

// Group returns an iterator that yields slices of consecutive elements from the input sequence
// that have the same group value as determined by the groupBy function.
func Group[T any, H comparable](s iter.Seq[T], groupBy func(T) H) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		var currentGroup []T
		var currentGroupKey H
		first := true

		for v := range s {
			key := groupBy(v)
			if first {
				currentGroupKey = key
				currentGroup = append(currentGroup, v)
				first = false
				continue
			}

			if key == currentGroupKey {
				currentGroup = append(currentGroup, v)
			} else {
				if !yield(currentGroup) {
					return
				}
				currentGroup = []T{v}
				currentGroupKey = key
			}
		}

		if len(currentGroup) > 0 {
			yield(currentGroup)
		}
	}
}

// Group2 returns an iterator that yields slices of consecutive key-value pairs from the input sequence
// that have the same group value as determined by the groupBy function.
func Group2[K, V any, H comparable](s iter.Seq2[K, V], groupBy func(K, V) H) iter.Seq2[H, []KV[K, V]] {
	return func(
		yield func(
			H, []KV[K, V],
		) bool,
	) {
		var currentGroup []KV[K, V]
		var currentGroupKey H
		first := true

		for k, v := range s {
			key := groupBy(k, v)
			if first {
				currentGroupKey = key
				currentGroup = append(currentGroup, KV[K, V]{k, v})
				first = false
				continue
			}

			if key == currentGroupKey {
				currentGroup = append(currentGroup, KV[K, V]{k, v})
			} else {
				if !yield(currentGroupKey, currentGroup) {
					return
				}
				currentGroup = []KV[K, V]{{k, v}}
				currentGroupKey = key
			}
		}

		if len(currentGroup) > 0 {
			yield(currentGroupKey, currentGroup)
		}
	}
}
