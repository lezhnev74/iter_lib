# Lib For Standard Iterators

Go's standard iterators `iter.Seq[v]` and `iter.Seq2[K,V]` have no function to combine
them into more useful pipelines. This lib provides that.

## Unary Iterators

- `Filter` and `Filter2`: Filter returns an iterator that yields only values that match a predicate function, while
  `Filter2`   does the same for key-value pairs
- `Map` and `Map2`: Transform values from an input sequence using a mapping function, with `Map2` handling key-value
  pairs
- `Dedup` and `Dedup2`: Remove consecutive duplicate values from a sequence using a custom equality function, with
  `Dedup2` handling key-value pairs
- `Group` and `Group2`: Group consecutive elements from a sequence based on a grouping function, with `Group2` handling
  key-value pairs and returning groups of elements with the same key
- `Chunk` and `Chunk2`: Split the input sequence into fixed-size chunks, with `Chunk2` handling key-value pairs and
  returning slices of elements of specified size. The last chunk may be smaller if the input length is not divisible by
  chunk size

## Binary Iterators

- `MergeOrdered` and `MergeOrdered2`: Merge two ordered sequences into a single sequence maintaining sort order, with
  `MergeOrdered2` handling key-value pairs 




 