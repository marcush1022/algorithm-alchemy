/*
Given a collection of integers that might contain duplicates, nums, return all possible subsets.

Note: The solution set must not contain duplicate subsets.

For example,
If nums = [1,2,2], a solution is:

[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]

和78一个思路，只不过这个有重复
*/

class Solution {
public:
    void dfs(int index, vector<vector<int>> &ret, vector<int> nums, vector<int> tmp)
    {
        ret.push_back(tmp);
        for(int i=index; i<nums.size(); i++)
        {
            tmp.push_back(nums[i]);
            dfs(i+1, ret, nums, tmp);
            tmp.pop_back();
            while(nums[i]==nums[i+1])
                i++;
        }
    }
    
    vector<vector<int>> subsetsWithDup(vector<int>& nums) {
        vector<vector<int>> ret;
        vector<int> tmp;
        sort(nums.begin(), nums.end());
        dfs(0, ret, nums, tmp);
        return ret;
    }
};
