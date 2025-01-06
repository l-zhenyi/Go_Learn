package main

import "fmt"

func main() {
	firstName, _ := getNames()
	fmt.Println("Welcome to browsing my Golang coursework,", firstName)
}

// don't edit below this line

func getNames() (string, string) {
	return "Ali", "Abu"
}
