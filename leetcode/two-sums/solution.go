func twoSum(nums []int, target int) []int {
    result := make([]int, 2)
    positions := make(map[int]int)
    for i, v := range nums {
        
        positionOfOther, present := positions[target - v]
        if present {
            result[0] = positionOfOther
            result[1] = i
            break
        } else {
            positions[v] = i
        }
    }
    return result
}

