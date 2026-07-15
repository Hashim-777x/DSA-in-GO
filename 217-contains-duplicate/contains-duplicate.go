func containsDuplicate(nums []int) bool {
    hashmap := make(map[int]bool, len(nums)) 
    for _, numb := range nums  {
        if _,val := hashmap[numb] ; val {
            return true 
        }
        hashmap[numb]= true
    }
    return false
}