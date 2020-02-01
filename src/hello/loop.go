package main

import "fmt"

func main() {

	i := 0
	for i < 10 {
		fmt.Print(i)
		i++
	}

	/*
		$ go run loop.go
		0123456789
	*/

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	/*
		$ go run loop.go
		0123456789
	*/

	// Odd Numbers using continue
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i)
	}

	/*
		$ go run loop.go
		13579
	*/

	breakCount := 5
	count := 0
	for {
		fmt.Println("Hello World 🌍")
		count++

		if count == breakCount {
			break
		}
	}

	/*
		$ go run loop.go
		Hello World 🌍
		Hello World 🌍
		Hello World 🌍
		Hello World 🌍
		Hello World 🌍
		...
		...
	*/

	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	for _, day := range days {
		fmt.Println(day)
	}

	/*
		$ go run loop.go
		sunday
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
	*/

	for {
		fmt.Println("Infinite loop")
	}

	/*
		$ go run loop.go
		Infinite loop
		Infinite loop
		Infinite loop
		Infinite loop
		Infinite loop
		...
		...
	*/

}
