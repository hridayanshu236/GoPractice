// variadic functions can be called with any number of trailing arguments. Function that is called with the varying number of arguments. i.e user is allowed to pass zero or more arguments in the variadic functions eg fmt.Println
package main

import "fmt"

func sum(nums ...int) { //...-->ellipsis
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
