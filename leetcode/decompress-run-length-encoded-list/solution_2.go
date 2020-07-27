func decompressRLElist(nums []int) []int {
  i := 0
  j := 0
  toFill := 0
  timesToFill := 0
  n := len(nums)
  
  var decompressed []int

  for i = 0; i < n; i+=2 {
    toFill = nums[i+1]
    timesToFill = nums[i]
    
    for j = 0; j < timesToFill; j++ {
      decompressed = append(decompressed, toFill)
    }
    
  }
  return decompressed;
}
