/*
Given a collection of numbers that might contain duplicates, return all possible unique
permutations.

For example,
[1,1,2] have the following unique permutations:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/

class Solution {
public:
    void dfs(vector<vector<int>> &ret, vector<int> nums, int index)
    {
        if(index==nums.size()-1)
        {
            ret.push_back(nums);
            return ;
        }
        sort(nums.begin()+index, nums.end());
        for(int i=index; i<nums.size(); i++)
        {
            if(i!=index && nums[i]==nums[i-1])
                continue;
            swap(nums[i], nums[index]);
            dfs(ret, nums, index+1);
            swap(nums[i], nums[index]);
        }
    }
    
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        vector<vector<int>> ret;
        dfs(ret, nums, 0);
        return ret;
    }
};
