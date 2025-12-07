package solutions

import "fmt"

func SeventhDay() error {
	fmt.Println("--- Day 7: Laboratories ---")
	// input, err := readFileLines("data/7_test.txt")
	input, err := readFileLines("data/7.txt")
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

	quantumManifold := make([][]int, len(input))
	for i, line := range input {
		quantumManifold[i] = make([]int, len(line))
		for j, r := range line {
			switch r {
			case '.':
				quantumManifold[i][j] = 0
			case '^':
				quantumManifold[i][j] = -1
			case 'S':
				quantumManifold[i][j] = 1
			}
		}
	}

	markTimelines(quantumManifold)
	// printWithTabs(quantumManifold)

	timelinesCount := totalTimelines(quantumManifold[len(quantumManifold)-1])
	fmt.Printf("Answer pt2: %v\n", timelinesCount)

	return nil
}

func totalTimelines(lastQuantumState []int) int {
	times := 0
	for _, r := range lastQuantumState {
		if r > 0 {
			times += int(r)
		}
	}
	return times
}

func markTimelines(manifold [][]int) {
	// technically we only need to keep track of current stat, i.e. last row for this
	// but it's nice to print whole result out
	for i := 1; i < len(manifold); i++ {
		for j := range manifold[i] {
			if manifold[i-1][j] > 0 {
				if manifold[i][j] != -1 {
					manifold[i][j] += manifold[i-1][j]
				} else {
					manifold[i][j+1] += manifold[i-1][j]
					manifold[i][j-1] += manifold[i-1][j]
				}
			}
		}
	}
}
