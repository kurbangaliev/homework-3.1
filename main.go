package main

import (
	"fmt"
	"math/rand"
)

func sortAsc(numbers []int) []int {
	for {
		ready := true
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				ready = false
			}
		}
		if ready {
			break
		}
	}
	return numbers
}

func main() {
	var lenArray = 10
	var numbers []int

	for i := 0; i < lenArray; i++ {
		numbers = append(numbers, rand.Intn(100))
	}
	fmt.Println("Generated numbers:", numbers)

	fmt.Println("copy slice array...")
	sliceCopy := make([]int, len(numbers))
	copy(sliceCopy, numbers[:])
	fmt.Println("sorted numbers:", sortAsc(sliceCopy))
	fmt.Println("original numbers:", numbers)
}
