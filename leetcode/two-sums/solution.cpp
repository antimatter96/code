class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
      vector<int> ans;
      ans.resize(2);
      map<int, int> mp;
      int second;
      int len = nums.size();
      for(int i = 0; i < len; i++ ) {
        second = target - nums[i];
        if(mp.count(second) > 0) {
          ans[0] = mp[second];
          ans[1] = i;
          break;
        } else {
          mp[nums[i]] = i;
        }
      } 
      return ans;
    }
};

