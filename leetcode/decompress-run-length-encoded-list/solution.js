/**
 * @param {number[]} nums
 * @return {number[]}
 */
var decompressRLElist = function(nums) {
  let decompressed = [];
  let count = 0;
  let number = 0;
  let j = 0;
  let i = 0;
  let n = nums.length;
  for(i = 0; i < n; i+=2) {
    
    count = nums[i];
    number = nums[i+1];
    
    for(j = 0; j < count; j++) {
      decompressed.push(number);
    }
    
  }
  return decompressed;
};
