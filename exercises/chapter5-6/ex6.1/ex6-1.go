package main

import (
	"bytes"
	"fmt"
	"math/bits"
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
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
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

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += bits.OnesCount64(word)
	}
	return count
}

// func (s *IntSet) Remove(x int) {
// 	var result IntSet
// 	var lastIndex int 

// 	// Add left part and ignore the element to be removed
// 	for i := 0; i < s.Len(); i++ {
// 		currentElement := s.words[i]
		
// 		if int(currentElement) == x {
// 			lastIndex = i
// 			break
// 		}
		
// 		result.Add(int(currentElement))
// 	}

// 	// Add right part
// 	for i := lastIndex + 1; i < s.Len(); i++ {
// 		currentElement := s.words[i]
// 		result.Add(int(currentElement))
// 	}

// 	s.words = result.words
// } 

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit // &^ is bit clear (AND NOT)
	}
}


func main() {
	var y IntSet

	y.Add(9)
	y.Add(42)
	y.Add(22)

	fmt.Println(y.Len())
	fmt.Println(y.words)
	y.Remove(42)
	fmt.Println(y.words)
}
