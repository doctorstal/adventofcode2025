package solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func TwelfthDay() error {
	fmt.Println("--- Day 12: Christmas Tree Farm ---")
	fmt.Println("The shapes are a lie!")
	input, err := readFileLines("data/12.txt")
	if err != nil {
		return err
	}

	fmt.Printf("input: %v\n", input)

	res := 0
	m := regexp.MustCompile(`^(\d+x\d+): (.*)$`)
	for _, l := range input {
		split := m.FindStringSubmatch(l)
		if split == nil {
			continue
		}

		dim := strings.Split(split[1], "x")
		w, _ := strconv.Atoi(dim[0])
		h, _ := strconv.Atoi(dim[1])

		pcounts := strings.Split(split[2], " ")
		sum := 0
		for _, c := range pcounts {
			n, _ := strconv.Atoi(c)
			sum += n
		}

		if w*h >= sum*9 {
			res += 1
		}
	}

	fmt.Printf("Answer: %v\n", res)

	return nil
}
