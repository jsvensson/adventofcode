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
	fmt.Printf("Common: %s (%s)", solve(ch), time.Since(start))
}

func solve(input <-chan string) string {
	ids := getValidIDs(input)

	for _, box := range ids {
		for _, box2 := range ids {
			if box == box2 {
				continue
			}

			if countMismatch(box, box2) == 1 {
				return getCommon(box, box2)
			}
		}
	}

	return "not found"
}

// Returns count of non-matching characters in two strings
func countMismatch(a, b string) int {
	var mismatch int
	for i, _ := range a {
		if a[i] != b[i] {
			mismatch++
		}
	}

	return mismatch
}

// Returns common characters in two strings
func getCommon(a, b string) string {
	result := make([]byte, 0)

	for i, _ := range a {
		if a[i] == b[i] {
			result = append(result, a[i])
		}
	}

	return string(result)
}

func getValidIDs(ch <-chan string) []string {
	valid := make([]string, 0)
	for str := range ch {
		count := make(map[int32]int)
		for _, v := range str {
			count[v] += 1
		}

		for _, v := range count {
			switch v {
			case 2, 3:
				valid = append(valid, str)
			}
		}
	}

	return valid
}
