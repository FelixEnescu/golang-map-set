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

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewFromSlice(t *testing.T) {
	cases := []struct {
		name     string
		slice    []string
		expected Set[string]
	}{
		{
			name:     "empty slice",
			slice:    []string{},
			expected: Set[string]{},
		},
		{
			name:     "non-empty slice",
			slice:    []string{"a", "b", "c"},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := NewFromSlice(c.slice)
			require.Equal(t, c.expected, actual, "expected %v, got %v", c.expected, actual)
		})
	}
}

func TestNewFromMapKeys(t *testing.T) {
	cases := []struct {
		name     string
		m        map[string]interface{}
		expected Set[string]
	}{
		{
			name:     "empty map",
			m:        map[string]interface{}{},
			expected: Set[string]{},
		},
		{
			name:     "non-empty map",
			m:        map[string]interface{}{"a": nil, "b": nil, "c": nil},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := NewFromMapKeys(c.m)
			require.Equal(t, c.expected, actual, "expected %v, got %v", c.expected, actual)
		})
	}
}

func TestSetToSlice(t *testing.T) {
	cases := []struct {
		name     string
		set      Set[string]
		expected []string
	}{
		{
			name:     "empty set",
			set:      Set[string]{},
			expected: []string{},
		},
		{
			name:     "non-empty set",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			expected: []string{"a", "b", "c"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.set.ToSlice()
			// sort actual and expected slices, as ToSlice() does not guarantee order
			slices.Sort(actual)
			slices.Sort(c.expected)
			require.Equal(t, c.expected, actual, "expected %v, got %v", c.expected, actual)
		})
	}
}

func TestSetEquals(t *testing.T) {
	cases := []struct {
		name     string
		set      Set[string]
		other    Set[string]
		expected bool
	}{
		{
			name:     "empty sets",
			set:      Set[string]{},
			other:    Set[string]{},
			expected: true,
		},
		{
			name:     "non-empty sets",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			expected: true,
		},
		{
			name:     "different sets",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "d": struct{}{}},
			expected: false,
		},
		{
			name:     "different sizes",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}},
			expected: false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.set.Equals(c.other)
			require.Equal(t, c.expected, actual, "expected %v, got %v", c.expected, actual)
		})
	}
}

func TestSetIsSubset(t *testing.T) {
	cases := []struct {
		name     string
		set      Set[string]
		other    Set[string]
		expected bool
	}{
		{
			name:     "empty sets",
			set:      Set[string]{},
			other:    Set[string]{},
			expected: true,
		},
		{
			name:     "non-empty sets",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			expected: true,
		},
		{
			name:     "different sets",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "d": struct{}{}},
			expected: false,
		},
		{
			name:     "different sizes",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}},
			expected: false,
		},
		{
			name:     "subset",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			expected: true,
		},
		{
			name:     "not a subset",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "d": struct{}{}},
			expected: false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.set.IsSubsetOf(c.other)
			require.Equal(t, c.expected, actual, "expected %v, got %v", c.expected, actual)
		})
	}
}

func TestSetIsProperSubset(t *testing.T) {
	cases := []struct {
		name     string
		set      Set[string]
		other    Set[string]
		expected bool
	}{
		{
			name:     "empty sets",
			set:      Set[string]{},
			other:    Set[string]{},
			expected: false,
		},
		{
			name:     "non-empty sets equals",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			expected: false,
		},
		{
			name:     "different sets",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "d": struct{}{}},
			expected: false,
		},
		{
			name:     "different sizes",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}},
			expected: false,
		},

		{
			name:     "subset",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			expected: true,
		},
		{
			name:     "not a subset",
			set:      Set[string]{"a": struct{}{}, "e": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "d": struct{}{}},
			expected: false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.set.IsProperSubsetOf(c.other)
			require.Equal(t, c.expected, actual, "expected %v, got %v", c.expected, actual)
		})
	}
}

func TestSetContains(t *testing.T) {
	cases := []struct {
		name     string
		set      Set[string]
		s        string
		expected bool
	}{
		{
			name:     "empty set",
			set:      Set[string]{},
			s:        "a",
			expected: false,
		},
		{
			name:     "non-empty set",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			s:        "a",
			expected: true,
		},
		{
			name:     "non-empty set, not found",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			s:        "d",
			expected: false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.set.Contains(c.s)
			require.Equal(t, c.expected, actual, "expected %v, got %v", c.expected, actual)
		})
	}
}

func TestSetAdd(t *testing.T) {
	cases := []struct {
		name     string
		set      Set[string]
		s        string
		expected Set[string]
	}{
		{
			name:     "empty set",
			set:      Set[string]{},
			s:        "a",
			expected: Set[string]{"a": struct{}{}},
		},
		{
			name:     "non-empty set",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			s:        "d",
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}, "d": struct{}{}},
		},
		{
			name:     "duplicate",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			s:        "a",
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.set.Add(c.s)
			require.Equal(t, c.expected, c.set, "expected %v, got %v", c.expected, c.set)
		})
	}
}

func TestSetRemove(t *testing.T) {
	cases := []struct {
		name     string
		set      Set[string]
		s        string
		expected Set[string]
	}{
		{
			name:     "empty set",
			set:      Set[string]{},
			s:        "a",
			expected: Set[string]{},
		},
		{
			name:     "non-empty set",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			s:        "d",
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
		},
		{
			name:     "duplicate",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			s:        "a",
			expected: Set[string]{"b": struct{}{}, "c": struct{}{}},
		},
		{
			name:     "resulting empty set",
			set:      Set[string]{"a": struct{}{}},
			s:        "a",
			expected: Set[string]{},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.set.Remove(c.s)
			require.Equal(t, c.expected, c.set, "expected %v, got %v", c.expected, c.set)
		})
	}
}

func TestSetAddAll(t *testing.T) {
	cases := []struct {
		name     string
		set      Set[string]
		slice    []string
		expected Set[string]
	}{
		{
			name:     "empty set",
			set:      Set[string]{},
			slice:    []string{"a", "b", "c"},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
		},
		{
			name:     "non-empty set",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			slice:    []string{"d", "e", "f"},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}, "d": struct{}{}, "e": struct{}{}, "f": struct{}{}},
		},
		{
			name:     "duplicate",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			slice:    []string{"a", "b", "c"},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
		},
		{
			name:     "overlapping",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}, "d": struct{}{}, "e": struct{}{}, "f": struct{}{}},
		},
		{
			name:     "adding empty slice",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			slice:    []string{},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.set.AddAll(c.slice)
			require.Equal(t, c.expected, c.set, "expected %v, got %v", c.expected, c.set)
		})
	}
}

func TestSetRemoveAll(t *testing.T) {
	cases := []struct {
		name     string
		set      Set[string]
		slice    []string
		expected Set[string]
	}{
		{
			name:     "empty set",
			set:      Set[string]{},
			slice:    []string{"a", "b", "c"},
			expected: Set[string]{},
		},
		{
			name:     "non-empty set",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			slice:    []string{"d", "e", "f"},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
		},
		{
			name:     "duplicate",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			slice:    []string{"a", "b", "c"},
			expected: Set[string]{},
		},
		{
			name:     "overlapping",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			expected: Set[string]{},
		},
		{
			name:     "removing empty slice",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			slice:    []string{},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
		},
		{
			name:     "partial overlap",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			slice:    []string{"a", "b", "d", "e", "f"},
			expected: Set[string]{"c": struct{}{}},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.set.RemoveAll(c.slice)
			require.Equal(t, c.expected, c.set, "expected %v, got %v", c.expected, c.set)
		})
	}
}

func TestSetUnion(t *testing.T) {
	cases := []struct {
		name     string
		set      Set[string]
		other    Set[string]
		expected Set[string]
	}{
		{
			name:     "empty sets",
			set:      Set[string]{},
			other:    Set[string]{},
			expected: Set[string]{},
		},
		{
			name:     "non-empty sets",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"d": struct{}{}, "e": struct{}{}, "f": struct{}{}},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}, "d": struct{}{}, "e": struct{}{}, "f": struct{}{}},
		},
		{
			name:     "duplicate",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
		},
		{
			name:     "overlapping",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}, "d": struct{}{}, "e": struct{}{}, "f": struct{}{}},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}, "d": struct{}{}, "e": struct{}{}, "f": struct{}{}},
		},
		{
			name:     "adding empty set",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
		},
		{
			name:     "partial overlap",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "d": struct{}{}, "e": struct{}{}, "f": struct{}{}},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}, "d": struct{}{}, "e": struct{}{}, "f": struct{}{}},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.set.Union(c.other)
			require.Equal(t, c.expected, actual, "expected %v, got %v", c.expected, actual)
		})
	}
}

func TestSetIntersection(t *testing.T) {
	cases := []struct {
		name     string
		set      Set[string]
		other    Set[string]
		expected Set[string]
	}{
		{
			name:     "empty sets",
			set:      Set[string]{},
			other:    Set[string]{},
			expected: Set[string]{},
		},
		{
			name:     "non-empty sets",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"d": struct{}{}, "e": struct{}{}, "f": struct{}{}},
			expected: Set[string]{},
		},
		{
			name:     "duplicate",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
		},
		{
			name:     "overlapping",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}, "d": struct{}{}, "e": struct{}{}, "f": struct{}{}},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
		},
		{
			name:     "partial overlap",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "d": struct{}{}, "e": struct{}{}, "f": struct{}{}},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}},
		},
		{
			name:     "adding empty set",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{},
			expected: Set[string]{},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.set.Intersection(c.other)
			require.Equal(t, c.expected, actual, "expected %v, got %v", c.expected, actual)
		})
	}
}

func TestSetDifference(t *testing.T) {
	cases := []struct {
		name     string
		set      Set[string]
		other    Set[string]
		expected Set[string]
	}{
		{
			name:     "empty sets",
			set:      Set[string]{},
			other:    Set[string]{},
			expected: Set[string]{},
		},
		{
			name:     "non-empty sets",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"d": struct{}{}, "e": struct{}{}, "f": struct{}{}},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
		},
		{
			name:     "duplicate",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			expected: Set[string]{},
		},
		{
			name:     "overlapping",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}, "d": struct{}{}, "e": struct{}{}, "f": struct{}{}},
			expected: Set[string]{},
		},
		{
			name:     "partial overlap",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{"a": struct{}{}, "b": struct{}{}, "d": struct{}{}, "e": struct{}{}, "f": struct{}{}},
			expected: Set[string]{"c": struct{}{}},
		},
		{
			name:     "adding empty set",
			set:      Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
			other:    Set[string]{},
			expected: Set[string]{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.set.Difference(c.other)
			require.Equal(t, c.expected, actual, "expected %v, got %v", c.expected, actual)
		})
	}
}
