package main

import "time"

func main() {
	i := 0

	switch i {
	case 0:
		println("i is 0")
	case 1:
		println("i is 1")
	case 2:
		println("i is 2")
	default:
		println("other")
	}

	//multiple conditions switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("It's the weekend")
	default:
		println("It's a weekday")
	}

	//type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			println("I am integer")
		case bool:
			println("I am boolean")
		default:
			println("I am of a different type")
		}
	}
	whoAmI(1)
	whoAmI(true)
	whoAmI("hello")
	whoAmI(3.14)

}
