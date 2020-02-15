package main

import (
	"fmt"
	"github.com/Chandan-Nayak/Go/lib"
	"sort"
	"time"
)

type input struct {
	numbers []int
	sum     int
}

func solutionOne(i input) []int {
	var result []int
	sum := i.sum
	numbers := i.numbers

	for k := 0; k < len(numbers); k++ {
		for l := k + 1; l < len(numbers); l++ {
			if numbers[k]+numbers[l] == sum {
				result = append(result, numbers[k], numbers[l])
			}
		}
	}
	return result
}

func solutionTwo(i input) []int {
	var result []int
	sum := i.sum
	numbers := i.numbers

	dict := make(map[int]int)

	for i := 0; i < len(numbers); i++ {
		num := numbers[i]
		diff := sum - num
		val, ok := dict[num]
		if ok {
			result = append(result, val, num)
			//return result
		} else {
			dict[diff] = num
		}
	}
	return result
}

func solutionThree(i input) []int {
	var result []int
	sum := i.sum
	numbers := i.numbers
	sort.Ints(numbers)

	lower := 0
	upper := len(numbers) - 1

	for upper > 0 {

		if numbers[lower]+numbers[upper] == sum {
			result = append(result, numbers[upper], numbers[lower])
			if upper-lower == 1 {
				break
			}
			upper = upper - 1
			lower = lower + 1
			//return result
			//break
		} else {
			if upper-lower == 1 {
				break
			}
			if numbers[lower]+numbers[upper] > sum {
				upper = upper - 1
			} else {
				lower = lower + 1
			}
		}

	}
	return result
}

func main() {

	currentInput := input{
		numbers: lib.Unique(lib.GenerateRandomSlice(20)),
		sum:     lib.GenerateRandomNumber(15),
	}

	fmt.Println("Input: ", currentInput.numbers)

	var solutions []func(i input) []int
	solutions = append(solutions, solutionOne, solutionTwo, solutionThree)

	var allElapsedTime = make(map[time.Duration]string)
	for _, fc := range solutions {
		start := time.Now()
		functionName := lib.GetFuncName(fc)
		fmt.Printf("%s > ArrayLength: %d, SUM: %d, PAIRS: %d\n",
			functionName, len(currentInput.numbers), currentInput.sum, fc(currentInput))
		elapsed := time.Since(start)
		allElapsedTime[elapsed] = functionName
	}
	lib.GetTimeComparision(allElapsedTime)

}
