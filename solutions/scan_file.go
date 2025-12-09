package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFileLines(path string) ([]string, error) {
	inputFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func readFileNumbers(path, separator string) ([][]int, error) {
	input, err := readFileLines(path)
	if err != nil {
		return nil, err
	}

	table := make([][]int, 0)
	for _, line := range input {
		unparsed := strings.Split(line, separator)
		row := make([]int, 0)
		for _, u := range unparsed {
			p, err := strconv.Atoi(u)
			if err != nil {
				return nil, fmt.Errorf("invalid input at [%s]:%w", line, err)
			}

			row = append(row, p)
		}
		table = append(table, row)

	}
	return table, nil
}
