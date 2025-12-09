package solutions

import (
	"fmt"
)

func NinthDay() error {
	fmt.Println("--- Day 9: Movie Theater ---")
	// input, err := readFileNumbers("data/9_test.txt", ",")
	input, err := readFileNumbers("data/9.txt", ",")
	if err != nil {
		return err
	}

	maxSquares := make([]int, len(input))

	for i, p1 := range input {
		for j := i + 1; j < len(input); j++ {
			p2 := input[j]
			area := (abs(p1[0]-p2[0]) + 1) * (abs(p1[1]-p2[1]) + 1)
			if area > maxSquares[i] {
				maxSquares[i] = area
			}
		}
	}

	max := 0
	for _, s := range maxSquares {
		if s > max {
			max = s
		}
	}

	fmt.Printf("Answer: %v\n", max)

	return nil
}
