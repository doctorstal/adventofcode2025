package solutions

import (
	"fmt"
)

func FourthDay() error {
	fmt.Println("--- Day 4: Printing Department ---")
	// input, err := readFileLines("data/4.txt")
	input, err := readFileLines("data/4_test.txt")
	if err != nil {
		return err
	}

	rollsMap := make([][]rune, len(input))
	for i, l := range input {
		rollsMap[i] = []rune(l)
	}

	res := 0

	for i, l := range rollsMap {
		for j, p := range l {
			if p == '@' && countadjAcentRolls(i, j, rollsMap) < 4 {
				res += 1
			}
		}
	}

	fmt.Printf("Answer: %v\n", res)

	res2 := 0

	for {
		removeMore := false

		for i, l := range rollsMap {
			for j, p := range l {
				if p == '@' && countadjAcentRolls(i, j, rollsMap) < 4 {
					res2 += 1
					l[j] = 'x'
					removeMore = true
				}
			}
		}

		if !removeMore {
			break
		}
	}
	fmt.Printf("Answer: %v\n", res2)

	return nil
}

func countadjAcentRolls(y, x int, input [][]rune) int {
	count := 0
	for i := max(y-1, 0); i <= min(y+1, len(input)-1); i += 1 {
		for j := max(x-1, 0); j <= min(x+1, len(input[0])-1); j += 1 {
			if i == y && j == x {
				continue
			}
			if input[i][j] == '@' {
				count += 1
			}
		}
	}
	return count
}
