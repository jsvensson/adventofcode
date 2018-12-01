package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jsvensson/adventofcode/2018/utils"
)

const data = "../data/01-a.txt"

func main() {
	ch := make(chan string)

	go utils.ReadLines(data, ch)

	var sum int
	for str := range ch {
		sum += parseValue(str)
	}

	fmt.Println("Sum:", sum)
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
