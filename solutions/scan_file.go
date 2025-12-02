package solutions

import (
	"bufio"
	"os"
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
