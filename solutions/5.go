package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type frange struct {
	start int
	end   int
}

func FifthDay() error {
	fmt.Println("--- Day 5: Cafeteria ---")
	// input, err := readFileLines("data/5_test.txt")
	input, err := readFileLines("data/5.txt")
	if err != nil {
		return err
	}

	ranges, err := parseRanges(input)
	if err != nil {
		return err
	}

	products, err := parseProducts(input)
	if err != nil {
		return err
	}

	freshCount := 0
	for _, p := range products {
		if isProductFresh(p, ranges) {
			freshCount += 1
		}
	}
	fmt.Printf("Answer: %v\n", freshCount)

	return nil
}

func isProductFresh(p int, ranges []frange) bool {
	for _, r := range ranges {
		if p >= r.start && p <= r.end {
			return true
		}
	}
	return false
}

func parseRanges(input []string) ([]frange, error) {
	ranges := make([]frange, 0)

	for _, l := range input {
		if l == "" {
			break
		}
		r := strings.Split(l, "-")
		if len(r) != 2 {
			return nil, fmt.Errorf("malformed input at line [%s]", l)
		}
		s, err := strconv.Atoi(r[0])
		if err != nil {
			return nil, err
		}
		e, err := strconv.Atoi(r[1])
		if err != nil {
			return nil, err
		}

		ranges = append(ranges, frange{start: s, end: e})
	}
	return ranges, nil
}

func parseProducts(input []string) ([]int, error) {
	products := make([]int, 0)
	read := false
	for _, l := range input {
		if !read {
			if l == "" {
				read = true
			}
			continue
		}
		p, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
