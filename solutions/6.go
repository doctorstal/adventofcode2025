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

	input2 := make([][]rune, len(input))
	for i, l := range input {
		input2[i] = []rune(l)
	}

	numbers := make([]int, 0)
	res2 := int64(0)

	for j := len(input2[0]) - 1; j >= 0; j-- {
		numbers = append(numbers, 0)
		for _, l := range input2 {
			switch l[j] {
			case ' ':
				continue
			case '+':
				res2 += sum(numbers)
				numbers = numbers[:0]
				j--
			case '*':
				res2 += mult(numbers)
				numbers = numbers[:0]
				j--
			default:
				numbers[len(numbers)-1] = numbers[len(numbers)-1]*10 + int(l[j]-'0')
			}
		}
	}
	fmt.Printf("Answer pt2: %v\n", res2)
	return nil
}

func sum(numbers []int) (res int64) {
	for _, n := range numbers {
		res += int64(n)
	}
	return
}

func mult(numbers []int) (res int64) {
	res = 1
	for _, n := range numbers {
		res *= int64(n)
	}
	return
}
