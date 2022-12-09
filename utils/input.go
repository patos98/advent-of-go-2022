package utils

import (
	"bufio"
	"os"
)

func GetFirstInputLine(input_path string) string {
	for line := range GetInputLines(input_path) {
		return line
	}
	return ""
}

func GetInputLines(input_path string) chan string {
	c := make(chan string)

	go func() {
		defer close(c)
		ProcessInputLines(input_path, func(line string) {
			c <- line
		})
	}()

	return c
}

type InputLineProcessor func(string)

func ProcessInputLines(input_path string, lp InputLineProcessor) {
	file, err := os.Open(input_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lp(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
