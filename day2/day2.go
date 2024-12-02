package day2

import (
	"fmt"
	"github.com/mikezm/advent-of-code-2024/helpers"
	"strconv"
	"strings"
)

const INPUTS = "./day2/input.txt"

type Challenge struct{}

func (d Challenge) A() {
	reports, err := readInput()
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	fmt.Println("reports: ", reports[2])
}

func (d Challenge) B() {
	fmt.Println("function B() not yet implemented")
}

func isSafe(r report) bool {

	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(r); i++ {
		dist := r[i] - r[i-1]
		if r[i] > r[i-1] {
			isDecreasing = false
		} else if r[i] < r[i-1] {
			isIncreasing = false
			dist = dist * -1
		}

		if dist >= 3 || dist == 0 || (!isIncreasing && !isDecreasing) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type report []int

func readInput() ([]report, error) {
	ir, err := helpers.NewReader(INPUTS)
	if err != nil {
		return nil, err
	}

	var data []report
	for _, line := range ir.Lines() {
		nums := strings.Fields(line)
		r := report{}
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			r = append(r, n)
		}
		data = append(data, r)
	}

	return data, nil
}
