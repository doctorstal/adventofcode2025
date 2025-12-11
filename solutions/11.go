package solutions

import (
	"container/list"
	"fmt"
	"strings"
)

type device struct {
	name    string
	outputs []string
}

func EleventhDay() error {
	fmt.Println("--- Day 11: Reactor ---")
	// input, err := readFileLines("data/11_test2.txt")
	input, err := readFileLines("data/11.txt")
	if err != nil {
		return err
	}

	devices := make(map[string]device)

	for _, line := range input {
		split := strings.Split(line, ": ")
		if len(split) != 2 {
			return fmt.Errorf("malformed input: %s", line)
		}
		devices[split[0]] = device{
			name:    split[0],
			outputs: strings.Split(split[1], " "),
		}
	}

	res := numPaths("you", "out", devices)

	fmt.Printf("Answer: %v\n", res)

	return nil
}

func numPaths(in, out string, graph map[string]device) int {
	res := 0
	q := list.New()
	q.PushBack(in)

	for q.Len() > 0 {
		curr := q.Front().Value.(string)
		q.Remove(q.Front())

		if curr == out {
			res += 1
		} else {
			for _, v := range graph[curr].outputs {
				q.PushBack(v)
			}
		}
	}
	return res
}
