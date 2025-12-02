package solutions

import (
	"fmt"
	"strings"
)

func SecondDay() error {
	fmt.Println("--- Day 2: Gift Shop ---")
	lines, err := readFileLines("data/2.txt")
	if err != nil {
		return err
	}

	input := strings.Join(lines, "")
	fmt.Printf("input: %v\n", input)

	return nil
}
