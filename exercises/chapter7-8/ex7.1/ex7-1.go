package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int 
type LineCounter int

type Stringer interface {
	String()string 
}


// implementation 1
func (c *WordCounter) CountWords(s string) (int, error) {
	words, err := c.extractWords(s)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	*c += WordCounter(len(words))

	return len(words), nil
}

func (c WordCounter) extractWords(s string) ([]string, error) {
	var words []string
	scanner := bufio.NewScanner(strings.NewReader(s))
    scanner.Split(bufio.ScanWords) 

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

// implementation 2
func (c *LineCounter) CountLines(s string) (int, error) {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err := scanner.Err()
	if err != nil {
		return 0, err
	}

	*c = LineCounter(len(lines))
	return len(lines), nil
}


// func (c *WordCounter) Write(s []string) (int, error) {

// }

// func (c *ByteCounter) Write(p []byte) (int, error) {
// 	fmt.Println(p)
	
// 	*c += ByteCounter(len(p))
	
// 	return len(p), nil
// }


func main() {
    input := "This is a random text"
	lineInput := "This is the first line \n This is the second line \n This is the third line"

	var c WordCounter
	c.CountWords(input)
	fmt.Println(c)

	var l LineCounter
	l.CountLines(lineInput)
	fmt.Println(l)

	// var c ByteCounter
	// c.Write([]byte("hållo"))

	// fmt.Println(c)
	// c = 0
	
	// var name = "Dolly"
	
	// fmt.Fprintf(&c, "hello %s", name)
	// fmt.Println(c)
}

// func (c *ByteCounter) Write(p []byte) {
// 	var c ByteCounter
// 	c.Write([]byte("hello"))
// 	fmt.Println(c)

// 	c = 0
// 	var name = "Dolly"
// 	fmt.Fprintf(&c, "hello, %s", name)
// 	fmt.Println(c)

// 	*c += ByteCounter(len(p))
	
// 	return len(p), nil
// }
