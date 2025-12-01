package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Advent Of Code time!")
	if err := firstDay(); err != nil {
		log.Fatal(err)
	}
}

func firstDay() error {
	// inputFile, err := os.Open("data/1_test.txt")
	inputFile, err := os.Open("data/1.txt")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(inputFile)
	pos := 50
	res := 0
	for scanner.Scan() {
		inst := scanner.Text()
		p, err := strconv.Atoi(inst[1:])
		if err != nil {
			return err
		}

		switch inst[0] {
		case 'L':
			pos -= p
		case 'R':
			pos += p
		}
		pos = (pos + 100) % 100
		if pos == 0 {
			res += 1
		}
	}
	fmt.Printf("Answer: %d\n", res)
	return nil
}
