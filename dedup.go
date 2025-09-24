package iter_lib

import "iter"

// Dedup returns an iterator that yields only non-duplicate consecutive values from the input sequence.
// The eq function is used to compare elements for equality.
func Dedup[T any](s iter.Seq[T], eq func(T, T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		var prev T
		first := true
		for v := range s {
			if first || !eq(prev, v) {
				if !yield(v) {
					return
				}
				prev = v
				first = false
			}
		}
	}
}

// Dedup2 returns an iterator that yields only key-value pairs where consecutive values are not duplicates.
// The eq function is used to compare values for equality.
func Dedup2[K, V any](s iter.Seq2[K, V], eq func(k1 K, v1 V, k2 K, v2 V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var prevVal V
		var prevKey K
		first := true
		for k, v := range s {
			if first || !eq(prevKey, prevVal, k, v) {
				if !yield(k, v) {
					return
				}
				prevVal = v
				prevKey = k
				first = false
			}
		}
	}
}
