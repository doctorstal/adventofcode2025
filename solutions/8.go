package solutions

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type point3d struct {
	x, y, z int
}

func EighthDay() error {
	fmt.Println("--- Day 8: Playground ---")
	// input, err := readFileLines("data/8_test.txt")
	// maxConn := 10
	input, err := readFileLines("data/8.txt")
	maxConn := 1000
	if err != nil {
		return err
	}

	jboxes := make([]point3d, len(input))

	for i, line := range input {
		unparsed := strings.Split(line, ",")
		if len(unparsed) != 3 {
			return fmt.Errorf("invalid input at line %d: [%s]", i, line)
		}
		var numbers [3]int
		for i := range 3 {
			numbers[i], err = strconv.Atoi(unparsed[i])
			if err != nil {
				return err
			}
		}
		jboxes[i] = point3d{
			x: numbers[0],
			y: numbers[1],
			z: numbers[2],
		}
	}

	type distance struct {
		dist   float64
		p1, p2 int
	}
	distances := make([]distance, 0)

	for i := range jboxes {
		for j := i + 1; j < len(jboxes); j++ {
			distances = append(distances, distance{
				dist: calcDistance3d(jboxes[i], jboxes[j]),
				p1:   i,
				p2:   j,
			})
		}
	}

	slices.SortFunc(distances, func(c1, c2 distance) int {
		return int(c1.dist - c2.dist)
	})

	pointToCircuit := make(map[int]int)
	for i := range len(jboxes) {
		pointToCircuit[i] = i
	}

	for i := range maxConn {
		d := distances[i]
		p1c := pointToCircuit[d.p1]
		p2c := pointToCircuit[d.p2]
		if p1c != p2c {
			for j := range pointToCircuit {
				if pointToCircuit[j] == p2c {
					pointToCircuit[j] = p1c
				}
			}
		}
	}

	circuitCount := make(map[int]int)
	for _, c := range pointToCircuit {
		circuitCount[c] += 1
	}

	curcuitsSizes := make([]int, 0, len(circuitCount))
	for _, n := range circuitCount {
		curcuitsSizes = append(curcuitsSizes, n)
	}

	slices.SortFunc(curcuitsSizes, func(a, b int) int { return b - a })

	res := 1
	for i := range 3 {
		res *= curcuitsSizes[i]
	}
	fmt.Printf("Answer: %v\n", res)

	return nil
}

func calcDistance3d(p1, p2 point3d) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}
