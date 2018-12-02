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
	ids := make([]string, 0)

	for str := range input {
		count := make(map[int32]int)
		for _, v := range str {
			count[v] += 1
		}

		for _, v := range count {
			switch v {
			case 2, 3:
				ids = append(ids, str)
			}
		}
	}

	for _, box := range ids {
		for _, box2 := range ids {
			if box == box2 {
				continue
			}

			var mismatch int
			for i, v := range box {
				if box2[i] != uint8(v) {
					mismatch++
				}
			}

			if mismatch == 1 {
				return getCommon(box, box2)
			}
		}
	}

	return "not found"
}

func getCommon(a, b string) string {
	result := make([]byte, 0)

	for i, _ := range a {
		if a[i] == b[i] {
			result = append(result, a[i])
		}
	}

	return string(result)
}
