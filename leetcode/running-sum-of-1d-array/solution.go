func runningSum(nums []int) []int {
  for i:=1; i < len(nums);i++ {
    nums[i] += nums[i-1]
  }
  return nums
}

func runningSum(nums []int) []int {
  for i:=0; i < len(nums) - 1;i++ {
    nums[i+1] += nums[i]
  }
  return nums
}
