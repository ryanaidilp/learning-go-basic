package main

import "fmt"

func main() {
	const firstName string = "Ryan"
	const lastName = "Pratama"

	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		fullName = "Ryan Pratama"
		value    = 100000
	)

	fmt.Println(fullName)
	fmt.Println(value)
}
