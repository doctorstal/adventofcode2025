package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func SecondDay() error {
	fmt.Println("--- Day 2: Gift Shop ---")
	lines, err := readFileLines("data/2_test.txt")
	if err != nil {
		return err
	}

	input := strings.Join(lines, "")
	ranges := strings.SplitSeq(input, ",")

	var sum uint64

	for r := range ranges {
		startEnd := strings.Split(r, "-")
		if len(startEnd) != 2 {
			return fmt.Errorf("invalid input, range has %d numbers", len(startEnd))
		}

		start, err := strconv.Atoi(startEnd[0])
		if err != nil {
			return err
		}

		end, err := strconv.Atoi(startEnd[1])
		if err != nil {
			return err
		}

		for i := start; i <= end; i += 1 {
			if repeatedTwice(i) {
				sum += uint64(i)
			}
		}

	}

	fmt.Printf("Answer: %d\n", sum)

	return nil
}

func repeatedTwice(n int) bool {
	str := strconv.FormatInt(int64(n), 10)
	l := len(str)
	if l%2 != 0 {
		return false
	}
	return str[0:l/2] == str[l/2:]
}
