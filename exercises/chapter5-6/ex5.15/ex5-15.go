package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var vals []int
	fmt.Println(vals)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	numbersAsStr := strings.Fields(line)
	for _, num := range numbersAsStr {
		n, err := strconv.Atoi(num)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		vals = append(vals, n)
	}

	min, max := getMinMaxFromSlice(vals...)
	fmt.Println(min)
	fmt.Println(max)
}

func getMinMaxFromSlice(vals ...int) (min int, max int) {
	min = math.MaxInt64
	max = math.MinInt64

	for _, v := range vals {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return
}
