func decompressRLElist(nums []int) []int {
  i := 0
  j := 0
  toFill := 0
  timesToFill := 0
  totalCount := 0
  n := len(nums)
  
  for i = 0; i < n; i+=2 {
    totalCount += nums[i]
  }
  
  decompressed := make([]int, totalCount)
  filledTillNow := 0

  for i = 0; i < n; i+=2 {
    toFill = nums[i+1]
    timesToFill = nums[i]
    
    for j = 0; j < timesToFill; j++ {
      decompressed[filledTillNow + j] = toFill
    }
    filledTillNow += timesToFill
    
  }
  return decompressed;
}
