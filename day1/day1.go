package day1

import (
	"fmt"
)

type Challenge struct{}

func (d Challenge) A() {
	quickSort(leftList, 0, len(leftList)-1)
	quickSort(rightList, 0, len(rightList)-1)

	var td int
	for i := range leftList {
		dist := leftList[i] - rightList[i]
		if dist > 0 {
			td += dist
		} else {
			td += dist * -1
		}
	}

	fmt.Println("Total Distance: ", td)

}

func (d Challenge) B() {
	var score int
	for _, value := range leftList {
		var matches int
		for _, rightValue := range rightList {
			if value == rightValue {
				matches++
			}
		}
		score += value * matches
	}

	fmt.Println("Score: ", score)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
