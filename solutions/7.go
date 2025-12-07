package solutions

import "fmt"

func SeventhDay() error {
	fmt.Println("--- Day 7: Laboratories ---")
	input, err := readFileLines("data/7_test.txt")
	// input, err := readFileLines("data/7.txt")
	if err != nil {
		return err
	}

	manifold := make([][]rune, len(input))
	for i, line := range input {
		manifold[i] = []rune(line)
	}

	splits := 0

	for i := 1; i < len(manifold); i++ {
		for j := range manifold[i] {
			if manifold[i-1][j] == 'S' || manifold[i-1][j] == '|' {
				if manifold[i][j] != '^' {
					manifold[i][j] = '|'
				} else {
					splits += 1
					manifold[i][j+1] = '|'
					manifold[i][j-1] = '|'
				}
			}
		}
	}

	fmt.Printf("Answer: %v\n", splits)

	return nil
}
