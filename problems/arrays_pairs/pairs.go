package main

import "fmt"

func main() {
	sum := 11
	arr := []int{1, 2, 3, 6, 5, 9, 10}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum {
				fmt.Printf("FOUND, %d + %d = %d\n", arr[i], arr[j], sum)
			}
		}
	}
}
