package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

// don't touch below this line

func main() {
	test("Bonnie Garmus,", " You wrote an incredible book!")
	test("Bruce Springsteen,", " Your songs are the best.")
	test("Go", " is rather fun to learn.")
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}
