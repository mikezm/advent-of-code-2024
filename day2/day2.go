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

	safeCount := 0
	for _, r := range reports {
		if isSafe(r) {
			safeCount++
		}
	}

	fmt.Println("Safe Reports: ", safeCount)
}

func (d Challenge) B() {
	reports, err := readInput()
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	safeCount := 0
	for _, r := range reports {
		if isSafeWhenDampened(r) {
			safeCount++
		}
	}

	fmt.Println("Safe Reports: ", safeCount)
}

func isSafe(r report) bool {
	isIncreasing := r[1]-r[0] >= 0

	for i := 1; i < len(r); i++ {
		dist := r[i] - r[i-1]
		if isIncreasing && dist < 0 {
			return false
		}
		if !isIncreasing && dist > 0 {
			return false
		}

		if dist < 0 {
			dist = dist * -1
		}

		if dist > 3 || dist == 0 {
			return false
		}
	}
	return true
}

func isSafeWhenDampened(r report) bool {

	if isSafe(r) {
		return true
	}

	for index := range r {
		newReport := removeElement(r, index)
		if isSafe(newReport) {
			return true
		}
	}

	return false
}

func removeElement(r report, index int) report {
	var newReport report
	for i := 0; i < index; i++ {
		newReport = append(newReport, r[i])
	}
	for j := index + 1; j < len(r); j++ {
		newReport = append(newReport, r[j])
	}

	return newReport
}

type report []int

func readInput() ([]report, error) {
	ir := helpers.NewInputReader(INPUTS)

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
