package solutions

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type frange struct {
	start int
	end   int
}

func (f frange) len() int {
	return f.end - f.start + 1
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

	ranges = mergeRanges(ranges)
	fmt.Printf("ranges: %v\n", ranges)

	freshListLen := 0
	for _, r := range ranges {
		freshListLen += r.len()
	}
	fmt.Printf("Answer pt2: %v\n", freshListLen)

	freshCount := 0
	for _, p := range products {
		if isProductFresh(p, ranges) {
			freshCount += 1
		}
	}
	fmt.Printf("Answer: %v\n", freshCount)

	return nil
}

func mergeRanges(input []frange) []frange {
	slices.SortFunc(input, func(a, b frange) int {
		return a.start - b.start
	})

	output := make([]frange, 1)
	output[0] = input[0]
	li := 0

	for _, r := range input {
		if r.start > output[li].end {
			output = append(output, r)
			li = len(output) - 1
		}
		if r.end > output[li].end {
			output[li].end = r.end
		}
	}

	return output
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
