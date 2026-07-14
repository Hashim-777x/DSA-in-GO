func twoSum(nums []int, target int) []int {
 seen := make(map[int]int, len(nums)) // allocating fixed size 
 for index, current := range nums {
    x := target - current
    if xIndex , value := seen[x] ; value {
        return []int {xIndex, index}
    }
    seen[current]= index
 }
 return nil // so save memory space 
 }
