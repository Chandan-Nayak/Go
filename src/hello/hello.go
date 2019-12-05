package main

import (
	"fmt" // Checkout > https://godoc.org/fmt
	"reflect"
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

// Main function - acts as the starting point
func main() {
	shortDynamicVars()
	packageLevelVars()
	mixedVar()

	fmt.Println("From another file in same package, MyLibVersion:", MyLibVersion)
}

// Short dynamic variables
// Go calls it -Type inference (automatic detection of type)
func shortDynamicVars() {
	firstName, lastName := "Alexa", "Amazon"
	company := "amazon"
	fmt.Printf("I am %v %v \nfirstName := %T\n", firstName, lastName, firstName)
	fmt.Println("lastName := ", reflect.TypeOf(lastName) )
	fmt.Println("I was built in:", company)
}

// Var keyword is mandatory for variables on package level either clubbed or separate
//	alexaAge is a package scope variable
func packageLevelVars() {
	fmt.Println("My age is:", alexaAge)
	//fmt.Println("My age is:", myl )
}

func mixedVar() {
	var a, b, c = 3, 4, "foo"
	fmt.Printf("a: %v, b: %v, c: %v\n", a, b, c)

}


/*
$ cd $GOPATH && cd src
$ go run hello

OUTPUT >
I am Alexa Amazon
firstName := string
lastName :=  string
I was built in: amazon
My age is: 3
a: 3, b: 4, c: foo
From another file in same package, MyLibVersion: 1.2
*/