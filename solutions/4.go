package solutions

import "fmt"

func FourthDay() error {
	fmt.Println("--- Day 4: Printing Department ---")
	input, err := readFileLines("data/4.txt")
	// input, err := readFileLines("data/4_test.txt")
	if err != nil {
		return err
	}

	res := 0

	for i, l := range input {
		for j, p := range l {
			if p == '@' && countadjAcentRolls(i, j, input) < 4 {
				res += 1
			}
		}
	}

	fmt.Printf("Answer: %v\n", res)

	return nil
}

func countadjAcentRolls(y, x int, input []string) int {
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
