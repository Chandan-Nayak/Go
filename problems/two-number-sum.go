package main

import (
	"fmt"
	"github.com/Chandan-Nayak/Go/lib"
	"sort"
)

/*
Find a pair of number from collection of numbers which is equal to a sum/number
    example: [1,2,3,4] > sum = 5 > True
    example: [1,2,6,5] > sum = 5 > False
*/

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
			return result
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

	for lower < upper {
		if numbers[lower]+numbers[upper] == sum {
			result = append(result, numbers[upper], numbers[lower])
			return result
			break
		} else {
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
	input1 := []int{3, 5, -4, 8, 11, 1, -1, 6}
	sum1 := 10

	var solutions []func(i input) []int
	solutions = append(solutions, solutionOne, solutionTwo, solutionThree)

	for _, fc := range solutions {
		fmt.Printf("%s > %d\n", lib.GetFuncName(fc), fc(input{numbers: input1, sum: sum1}))
	}

}
