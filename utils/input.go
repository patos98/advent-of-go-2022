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

		file, err := os.Open(input_path)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}()

	return c
}
