func containsDuplicate(nums []int) bool {
    hashmap := make(map[int]bool, len(nums)) 
    for i:=0 ; i<len(nums) ; i++   {
        if _,val := hashmap[nums[i]] ; val {
            return true 
        }
        hashmap[nums[i]]= true
    }
    return false
}