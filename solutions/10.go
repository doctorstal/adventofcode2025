package solutions

import (
	"container/list"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type machine struct {
	lights     int64
	lightsReq  int64
	buttons    []int64
	joltageReq []int
}

// var machineRegexp *regexp.Regexp = regexp.MustCompile(`\[(?P<lightsReq>[.#]+)\] (?P<buttons>\([0-9,]+\) )+ (?P<joltageReq>\{[0-9,]+\})`)
var (
	machineRegexp *regexp.Regexp = regexp.MustCompile(`\[(?P<lightsReq>[.#]+)\] ((:?\([0-9,]+\) )+)`)
	buttonsRegexp *regexp.Regexp = regexp.MustCompile(`\(([0-9,]+)\) `)
)

func newMachine(line string) *machine {
	parsed := machineRegexp.FindStringSubmatch(line)
	lightsReq := int64(0)
	for i, r := range parsed[1] {
		if r == '#' {
			lightsReq += int64(math.Pow(2, float64(i)))
		}
	}
	// fmt.Printf("%s: %s\n", parsed[1], strconv.FormatInt(lightsReq, 2))
	parsedButtons := buttonsRegexp.FindAllStringSubmatch(parsed[2], -1)
	buttons := make([]int64, 0, len(parsedButtons))
	for _, btn := range parsedButtons {
		ranksS := strings.Split(btn[1], ",")
		b := int64(0)
		for _, rS := range ranksS {
			r, _ := strconv.Atoi(rS)
			b += int64(math.Pow(2, float64(r)))
		}
		// fmt.Printf("%s: %s\n", btn[1], strconv.FormatInt(b, 2))
		buttons = append(buttons, b)
	}

	return &machine{
		lightsReq: lightsReq,
		buttons:   buttons,
	}
}

func TenthDay() error {
	fmt.Println("--- Day 10: Factory ---")
	input, err := readFileLines("data/10.txt")
	// input, err := readFileLines("data/10_test.txt")
	if err != nil {
		return err
	}

	res := 0
	for _, l := range input {
		machine := newMachine(l)
		// fmt.Printf("machine.lightsReq: %v\n", strconv.FormatInt(machine.lightsReq, 2))
		// found := minPresses(machine.lightsReq, 0, machine.buttons, make(map[int64]bool))
		found := bfsPresses(machine.lightsReq, machine.buttons)
		// fmt.Printf("found: %v\n", found)
		res += found
	}
	fmt.Printf("Answer: %v\n", res)
	return nil
}

func bfsPresses(req int64, buttons []int64) int {
	type step struct {
		value   int64
		btn     int64
		presses int
		prev    *step
	}
	// printSteps := func(s *step) {
	// 	for s != nil {
	// 		fmt.Printf("step %d: %s\n", s.presses, strconv.FormatInt(s.btn, 2))
	// 		s = s.prev
	// 	}
	// }
	q := list.New()
	q.PushBack(&step{0, 0, 0, nil})

	visited := make(map[int64]bool)

	for q.Len() > 0 {
		curr := q.Front().Value.(*step)
		q.Remove(q.Front())

		if curr.value == req {
			// printSteps(curr)
			return curr.presses
		}

		if visited[curr.value] {
			continue
		}
		visited[curr.value] = true

		for _, b := range buttons {
			q.PushBack(&step{curr.value ^ b, b, curr.presses + 1, curr})
		}
	}
	return -1
}
