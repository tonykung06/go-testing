package pack

import "fmt"

func Add(nums ...int) int {
	var sum int
	if len(nums) < 1 {
		fmt.Println("No args provided.")
		return 0
	}
	for _, i := range nums {
		sum += i
	}
	return sum
}
