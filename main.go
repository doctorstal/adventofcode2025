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
			if pos == 0 && p != 0 {
				res -= 1
			}
			pos -= p
			for pos < 0 {
				fmt.Println(inst)
				pos += 100
				res += 1
			}
			if pos == 0 {
				fmt.Println(inst)
				res += 1
			}
		case 'R':
			pos += p
			for pos >= 100 {
				fmt.Println(inst)
				pos -= 100
				res += 1
			}
		}
		// if pos == 0 {
		// 	fmt.Println(inst)
		// 	res += 1
		// }
		// fmt.Println(pos)
	}
	fmt.Printf("Answer: %d\n", res)
	return nil
}
