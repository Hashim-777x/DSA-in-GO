package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, ok := indexMap[complement]; ok {
			return []int{j, i}
		}
		indexMap[num] = i
	}

	return nil
}

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(arr, target)
	fmt.Println("nums:", arr)
	fmt.Println("target:", target)
	fmt.Println("result:", result)
}

func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}	
