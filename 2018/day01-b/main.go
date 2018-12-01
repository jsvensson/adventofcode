package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jsvensson/adventofcode/2018/utils"
)

const data = "../data/01-a.txt"

func main() {
	known := make(map[int]struct{})
	var sum int

	for {
		ch := make(chan string)
		go utils.ReadLines(data, ch)
		for str := range ch {
			sum += parseValue(str)
			if _, ok := known[sum]; ok {
				fmt.Println("First repeat:", sum)
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
