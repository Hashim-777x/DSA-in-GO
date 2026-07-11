func twoSum(nums []int, target int) []int {
      var  index = []int{0,0}
    for i:=0 ; i<len(nums) ; i++{
        for j:= i+1; j <len(nums);j++{
          sum := nums[i]+nums[j]

        if target == sum {
         index[0]= i
         index[1]= j
            return index
          }
    }
    }
        return index
}