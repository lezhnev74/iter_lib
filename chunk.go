package iter_lib

import "iter"

// Chunk splits the input sequence into chunks of the specified size.
// The last chunk may contain fewer elements if the input length is not divisible by size.
// Size must be greater than 0.
func Chunk[T any](input iter.Seq[T], size int) iter.Seq[[]T] {
	if size <= 0 {
		panic("chunk size must be positive")
	}

	return func(yield func([]T) bool) {
		buffer := make([]T, 0, size)

		for v := range input {
			buffer = append(buffer, v)

			if len(buffer) == size {
				if !yield(buffer) {
					return
				}
				buffer = make([]T, 0, size)
			}
		}

		if len(buffer) > 0 {
			yield(buffer)
		}
	}
}

// Chunk2 splits the input key-value sequence into chunks of the specified size.
// The last chunk may contain fewer elements if the input length is not divisible by size.
// Size must be greater than 0.
func Chunk2[K, V any](input iter.Seq2[K, V], size int) iter.Seq[[]KV[K, V]] {
	if size <= 0 {
		panic("chunk size must be positive")
	}

	return func(yield func([]KV[K, V]) bool) {
		buffer := make([]KV[K, V], 0, size)

		for k, v := range input {
			buffer = append(buffer, KV[K, V]{K: k, V: v})

			if len(buffer) == size {
				if !yield(buffer) {
					return
				}
				buffer = make([]KV[K, V], 0, size)
			}
		}

		if len(buffer) > 0 {
			yield(buffer)
		}
	}
}
