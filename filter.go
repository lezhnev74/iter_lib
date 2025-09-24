package iter_lib

import "iter"

// Filter returns an iterator that yields the values from s for which f returns true.
func Filter[V any](s iter.Seq[V], f func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range s {
			if f(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Filter2 returns an iterator that yields the values from s for which f returns true.
// The values are paired with their corresponding keys.
func Filter2[K, V any](s iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range s {
			if f(k, v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
