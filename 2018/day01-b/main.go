package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jsvensson/adventofcode/2018/utils"
)

const data = "../data/01-a.txt"

func main() {
	start := time.Now()

	var sum int
	var values []string
	known := make(map[int]struct{})
	ch := make(chan string)

	go utils.ReadLines(data, ch)

	for str := range ch {
		values = append(values, str)
	}

	for {
		for _, str := range values {
			sum += parseValue(str)
			if _, ok := known[sum]; ok {
				fmt.Printf("First repeat: %d (%s)\n", sum, time.Since(start))
				os.Exit(0)
			}

			known[sum] = struct{}{}
		}
	}
}

func parseValue(s string) int {
	val, err := strconv.Atoi(s[1:])
	if err != nil {
		panic(err)
	}

	switch {
	case strings.HasPrefix(s, "+"):
		return val
	default:
		return -val
	}
}
