/**
 * @param {number[]} candies
 * @param {number} extraCandies
 * @return {boolean[]}
 */
var kidsWithCandies = function(candies, extraCandies) {
    var arr = new Array(candies.length);

    let max = 0;
    for (let i= 0; i < candies.length; i++) {
      if (max < candies[i]) {
        max = candies[i]
      }
    }

    for (let i= 0; i < candies.length; i++) {
      if (max <= candies[i] + extraCandies) {
        arr[i] = true
      } else {
        arr[i] = false
      }
    }
    return arr
};
