/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    let mp = {};
    let len = nums.length;
    for (let i = 0; i < len; i++) {
        if( mp[ target - nums[i] ] != undefined) {
            return [mp[ target - nums[i] ], i];
        } else {
            mp[nums[i]] = i;
        }
    }
};

