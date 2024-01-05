/*
Open Source Initiative OSI - The MIT License (MIT):Licensing

The MIT License (MIT)
Copyright (c) 2024 Felix Enescu

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package set

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(map[T]struct{})
}

// NewFromSlice creates a new Set from a slice of comparable.
func NewFromSlice[T comparable](slice []T) Set[T] {
	set := New[T]()
	for _, s := range slice {
		set[s] = struct{}{}
	}
	return set
}

// NewFromMapKeys creates a new Set from a map's keys.
func NewFromMapKeys[T comparable, V any](m map[T]V) Set[T] {
	set := New[T]()
	for k := range m {
		set[k] = struct{}{}
	}
	return set
}

// ToSlice returns an unordered slice of elements from a Set.
func (set Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(set))
	for s := range set {
		slice = append(slice, s)
	}
	return slice
}

// Equals returns true if two Sets are equal.
func (set Set[T]) Equals(other Set[T]) bool {
	if len(set) != len(other) {
		return false
	}
	for s := range set {
		if _, ok := other[s]; !ok {
			return false
		}
	}
	return true
}

// Contains returns true if a Set contains an element.
func (set Set[T]) Contains(s T) bool {
	_, ok := set[s]
	return ok
}

// IsSubsetOf returns true if a Set is a subset of another Set (they can be equal).
func (set Set[T]) IsSubsetOf(other Set[T]) bool {
	if len(set) > len(other) {
		return false
	}

	for item := range set {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// IsProperSubsetOf returns true if a Set is a proper subset of another Set.
func (set Set[T]) IsProperSubsetOf(other Set[T]) bool {
	return len(set) < len(other) && set.IsSubsetOf(other)
}

// Add adds an element to a Set.
func (set Set[T]) Add(s T) {
	set[s] = struct{}{}
}

// Remove removes an element from a Set.
func (set Set[T]) Remove(s T) {
	delete(set, s)
}

// AddAll adds a slice of elements to a Set.
func (set Set[T]) AddAll(slice []T) {
	for _, s := range slice {
		set[s] = struct{}{}
	}
}

// RemoveAll removes a slice of elements from a Set.
func (set Set[T]) RemoveAll(slice []T) {
	for _, s := range slice {
		delete(set, s)
	}
}

// Union returns the union of two Sets as new Set.
func (set Set[T]) Union(other Set[T]) Set[T] {
	result := make(Set[T])
	for s := range set {
		result[s] = struct{}{}
	}
	for s := range other {
		result[s] = struct{}{}
	}
	return result
}

// Intersection returns the intersection of two Sets as new Set.
func (set Set[T]) Intersection(other Set[T]) Set[T] {
	result := make(Set[T])
	for s := range set {
		if _, ok := other[s]; ok {
			result[s] = struct{}{}
		}
	}
	return result
}

// Difference returns the difference of two Sets as new Set.
func (set Set[T]) Difference(other Set[T]) Set[T] {
	result := make(Set[T])
	for s := range set {
		if _, ok := other[s]; !ok {
			result[s] = struct{}{}
		}
	}
	return result
}
