package main

import "fmt"

const GLOBAL = "global"

func main() {
	const NAME = "ABC"

	const (
		ADD_1 = "Add1"
		ADD_2 = "Add2"
	)

	fmt.Println(GLOBAL, NAME, ADD_1, ADD_2)
}
