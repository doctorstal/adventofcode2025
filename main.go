package main

import (
	"adventofcode2025/solutions"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Advent Of Code time!")
	var solution func() error
	switch os.Args[1] {
	case "1":
		solution = solutions.FirstDay
	case "2":
		solution = solutions.SecondDay
	case "3":
		solution = solutions.ThirdDay
	case "4":
		solution = solutions.FourthDay
	case "5":
		solution = solutions.FifthDay
	case "6":
		solution = solutions.SixthDay
	case "7":
		solution = solutions.SeventhDay
	case "8":
		solution = solutions.EighthDay
	default:
		log.Fatal("Which day to solve? \nProvide number in run arguments")
	}
	started := time.Now()
	if err := solution(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Elapsed: %s\n", time.Since(started))
}
