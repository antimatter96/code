class Solution {
public:
    vector<int> runningSum(vector<int>& nums) {
        vector<int> ans;
        int len = nums.size();
        ans.resize(len);
        int sum = 0;
        for(int i = 0; i < len; i++) {
          sum += nums[i];
          ans[i] = sum;
        }
        return ans;
    }
};

