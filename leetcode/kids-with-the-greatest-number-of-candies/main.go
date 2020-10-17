func kidsWithCandies(candies []int, extraCandies int) []bool {
  toRet := make([]bool, len(candies))
  var max int;
  for i:= 0; i < len(candies); i++ {
    if (max < candies[i]) {
      max = candies[i]
    }
  }
  
  for i:= 0; i < len(candies); i++ {
    if (max <= extraCandies + candies[i]) {
      toRet[i] = true
    }
  }
  return toRet
  
}

