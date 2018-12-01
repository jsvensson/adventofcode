package utils

import (
	"bufio"
	"fmt"
	"os"
)

// ReadLines reads one line at a time from an input file,
// and outputs the lines on the provided channel.
func ReadLines(fileName string, outCh chan<- string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		outCh <- scanner.Text()
	}

	close(outCh)
}
