func numIdenticalPairs(nums []int) int {
  count := 0
  mp := make(map[int]int)
  n := len(nums)
  
  for i := 0 ; i < n ; i++ {
    mp[nums[i]]++
  }
  
  var x int
  for i := 0 ; i < n ; i++ {
    x = mp[nums[i]]
    if x > 1 {
      count += (x-1)
      mp[nums[i]]--
    }
    
    
    
  }
  
  return count
}
