package main

import (
	"fmt" // Checkout > https://godoc.org/fmt
	"time"
)

const (
	author string = "Chandan Nayak"
	place  string = "Pune"
)

var (
	year             int = time.Now().Year() //variable with initialization
	alexaReleaseYear int = 2016
	alexaAge         int = year - alexaReleaseYear
)

func main() {
	shortDynamicVars()
	packageLevelVars()
	mixedVar()
}

func shortDynamicVars() {
	firstName, lastName := "Alexa", "Amazon"
	fmt.Printf("I am %v %v\n", firstName, lastName)
}

func packageLevelVars() {
	fmt.Println("My age is:", alexaAge)
}

func mixedVar() {
	var a, b, c = 3, 4, "foo"
	fmt.Printf("a: %v, b: %v, c: %v", a, b, c)
}
