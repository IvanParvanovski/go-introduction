package main

import (
	"bytes"
	"fmt"
)

// its zero value represents the empty set
type IntSet struct {
	words []uint64
}

// reports awhether the set contains the non-negative value x
func (s *IntSet) Has(x int) bool {
	word, bit :=  x/64, uint(x%64)
	return word < len(s.words) && s.words[word] & (1 << bit) != 0
}

// adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x / 64, uint(x % 64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// adds the non-negative value x to the set.
func (s *IntSet) AddAll(numbers ...int) {
	for _, x := range numbers {
		word, bit := x / 64, uint(x % 64)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}

// sets s to the union of s and t 
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1 << uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var y IntSet
	// x.Add(1)
	// x.Add(144)
	// x.Add(9)
	// fmt.Println(x.String())
	// fmt.Println(x.Len())

	y.Add(9)
	y.Add(42)
	y.AddAll(80, 30, 2)
	fmt.Println(y.String())

	// x.UnionWith(&y)
	// fmt.Println(x.String())
	// fmt.Println(x.Has(9), x.Has(123))
}