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
	"math/rand"
	"testing"
)

const setSize = 1000

func randomInts(n int) []int {
	i := make([]int, n)
	for ind := range i {
		i[ind] = rand.Int()
	}
	return i
}

func BenchmarkAdd(b *testing.B) {
	ints := randomInts(setSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := New[int]()
		for _, v := range ints {
			s.Add(v)
		}
	}
}

func BenchmarkAddAll(b *testing.B) {
	ints := randomInts(setSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := New[int]()
		s.AddAll(ints)
	}
}