package solutions

import (
	"fmt"
	"math"
)

func ThirdDay() error {
	fmt.Println("--- Day 3: Lobby ---")

	banks, err := readFileLines("data/3.txt")
	// banks, err := readFileLines("data/3_test.txt")
	if err != nil {
		return err
	}

	totalJoltage := 0
	totalJoltagePt2 := 0
	for _, bank := range banks {
		totalJoltage += findMaxJoltageInBank(bank)
		totalJoltagePt2 += findMaxJoltageInBankPt2(bank)
	}

	fmt.Printf("Answer: %d\n", totalJoltage)
	fmt.Printf("Answer Pt2: %d\n", totalJoltagePt2)

	return nil
}

func findMaxJoltageInBank(bank string) int {
	currPos := 0
	res := 0
	for i := 2; i >= 1; i -= 1 {
		pos := posMaxDigitInBank(bank, currPos, len(bank)-i+1)
		currPos = pos + 1
		res += int(bank[pos]-'0') * int(math.Pow10(i-1))
	}
	return res
}

func findMaxJoltageInBankPt2(bank string) int {
	currPos := 0
	res := 0
	for i := 12; i >= 1; i -= 1 {
		pos := posMaxDigitInBank(bank, currPos, len(bank)-i+1)
		currPos = pos + 1
		res += int(bank[pos]-'0') * int(math.Pow10(i-1))
	}
	return res
}

func posMaxDigitInBank(bank string, start, end int) int {
	var m1 byte = 0
	m1Pos := 0
	for i := start; i < end; i++ {
		if bank[i] > m1 {
			m1 = bank[i]
			m1Pos = i
		}
	}
	return m1Pos
}
