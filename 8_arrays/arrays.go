package main

import "fmt"

func main() {
	var nums [5]int
	nums[0] = 1

	fmt.Println(nums[0])
	fmt.Println(len(nums))

	values := [5]int{1, 2, 3, 4, 5}
	fmt.Println(values)

	array := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(array)

}
