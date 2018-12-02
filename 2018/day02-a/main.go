package main

import (
	"fmt"
	"time"

	"github.com/jsvensson/adventofcode/2018/utils"
)

const data = "../data/02.txt"

func main() {
	start := time.Now()
	ch := make(chan string)
	go utils.ReadLines(data, ch)
	fmt.Printf("Checksum: %d (%s)", solve(ch), time.Since(start))
}

func solve(input <-chan string) int {
	var twice, thrice int

	for str := range input {
		count := make(map[int32]int)
		for _, v := range str {
			count[v] += 1
		}

		var seenTwo, seenThree bool

		for _, v := range count {
			switch v {
			case 2:
				seenTwo = true
			case 3:
				seenThree = true
			}
		}

		if seenTwo {
			twice++
		}
		if seenThree {
			thrice++
		}

	}

	return twice * thrice
}
