package solutions

import "fmt"

func ThirdDay() error {
	fmt.Println("--- Day 3: Lobby ---")

	banks, err := readFileLines("data/3.txt")
	// banks, err := readFileLines("data/3_test.txt")
	if err != nil {
		return err
	}

	totalJoltage := 0
	for _, bank := range banks {
		totalJoltage += findMaxJoltageInBank(bank)
	}

	fmt.Printf("Answer: %d\n", totalJoltage)

	return nil
}

func findMaxJoltageInBank(bank string) int {
	var m1 byte = 0
	m1Pos := 0
	for i := range len(bank) - 1 {
		if bank[i] > m1 {
			m1 = bank[i]
			m1Pos = i
		}
	}
	var m2 byte = 0
	for i := m1Pos + 1; i < len(bank); i += 1 {
		if bank[i] > m2 {
			m2 = bank[i]
		}
	}
	// fmt.Printf("m1: %v, m2: %v\n", m1-'0', m2-'0')

	return int((m1-'0')*10 + (m2 - '0'))
}
