package solutions

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SixthDay() error {
	fmt.Println("--- Day 6: Trash Compactor ---")
	// input, err := readFileLines("data/6_test.txt")
	input, err := readFileLines("data/6.txt")
	if err != nil {
		return err
	}

	// numbers := make([][]int, 0)

	actions := strings.Split(input[len(input)-1], " ")
	actions = slices.DeleteFunc(actions, func(a string) bool { return a == "" })

	results := make([]int64, len(actions))
	for i, a := range actions {
		switch a {
		case "*":
			results[i] = 1
		case "+":
			results[i] = 0
		}
	}

	for i := range len(input) - 1 {
		unparsed := strings.Split(input[i], " ")
		unparsed = slices.DeleteFunc(unparsed, func(u string) bool { return u == "" })
		for j, u := range unparsed {
			n, err := strconv.Atoi(u)
			if err != nil {
				return err
			}

			switch actions[j] {
			case "*":
				results[j] *= int64(n)
			case "+":
				results[j] += int64(n)
			}

		}
	}

	res := int64(0)
	for _, r := range results {
		res += r
	}
	fmt.Printf("Answer: %v\n", res)

	return nil
}
