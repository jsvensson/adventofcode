package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jsvensson/adventofcode/2018/utils"
)

const data = "../data/01-a.txt"

func main() {
	start := time.Now()

	var sum int
	ch := make(chan string)

	go utils.ReadLines(data, ch)

	for str := range ch {
		sum += parseValue(str)
	}

	fmt.Printf("Sum: %d (%s)\n", sum, time.Since(start))
}

func parseValue(s string) int {
	val, err := strconv.Atoi(s[1:])
	if err != nil {
		panic(err)
	}

	switch {
	case strings.HasPrefix(s, "-"):
		return -val
	default:
		return val
	}
}
