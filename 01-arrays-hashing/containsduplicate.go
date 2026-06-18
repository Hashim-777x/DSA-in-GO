package main 

import "fmt"

func containsDuplicate(nums []int) bool	{
 if len(nums) <= 1 {
 	return false
 }	
 seen := make(map[int]bool)
 for _, num := range nums {
 	if seen[num] {
 		return true
 	}
 	seen[num] = true
 }
 return false
}

func main() {
	arr := []int{1, 2, 3, 1}
	result := containsDuplicate(arr)

	fmt.Println("nums:", arr)
	fmt.Println("result:", result)
}
