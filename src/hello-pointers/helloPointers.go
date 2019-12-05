package main

import "fmt"

func main()  {

	var osType string = "ubuntu"
	var p = &osType
	fmt.Println("Value of p:", *p)
	fmt.Println("Value of osType:", osType)
	fmt.Println("Address of osType:", &osType)
	fmt.Println("Address of p:", &p)

	*p = "darwin"
	fmt.Println("-------AFTER CHANGE--------")
	fmt.Println("Value of p:", *p)
	fmt.Println("Value of osType:", osType)
	fmt.Println("Address of osType:", &osType)
	fmt.Println("Address of p:", &p)

}


/*
$ cd $GOPATH && cd src
$ go run hello-pointers

OUTPUT >
Value of p: ubuntu
Value of osType: ubuntu
Address of osType: 0xc0000741e0
Address of p: 0xc000094018
-------AFTER CHANGE--------
Value of p: darwin
Value of osType: darwin
Address of osType: 0xc0000741e0
Address of p: 0xc000094018

*/