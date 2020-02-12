package main

import "fmt"

func main() {

	for n := 0; n <= 10; n++ {
		if n == 0 {
			continue
		} else if n%2 == 0 {
			fmt.Printf("%d: even number\n", n)
		} else {
			fmt.Printf("%d: odd number\n", n)
		}
	}

	/*
		$ go run conditionals.go
		1: odd number
		2: even number
		3: odd number
		4: even number
		5: odd number
		6: even number
		7: odd number
		8: even number
		9: odd number
		10: even number
	*/

	// case
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("not valid")
	}

	/*
		$ go run conditionals.go
		Write 2 as two
	*/

	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

	/*
		$ go run conditionals.go
		vowel
	*/


	food := "s"
	switch food {
	case "chicken", "mutton", "egg", "fish", "crabs":
		fmt.Println("Non vegetarian")
	case "potato", "eggplant", "tomato", "onion", "carrot", "cheese":
		fmt.Println("Vegetarian")
	default:
		fmt.Println("We do not know about the food")
	}

	/*
		$ go run conditionals.go
		We do not know about the food
	*/

	num := 75
	switch {
	case num >= 0 && num <= 50:
		fmt.Println("number is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("number is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("number is greater than 100")
	}

	/*
		$ go run conditionals.go
		number is greater than 51 and less than 100
	*/

	number := 70
	switch num := number; { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}

	/*
		$ go run conditionals.go
		70 is lesser than 100
		70 is lesser than 200

	*/

}