/*
Given a collection of candidate numbers (C) and a target number (T), find all unique combinations in C where the 
candidate numbers sums to T.

Each number in C may only be used once in the combination.

Note:
All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
For example, given candidate set [10, 1, 2, 7, 6, 1, 5] and target 8, 
A solution set is: 
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]

Note: 不同点: 
1. 原数组有重复
2. 只能用一次
3. 结果不能有重复
*/

class Solution {
public:
    void dfs(vector<int> nums, vector<vector<int>> &ret, vector<int> &tmp, int &sum, int index, int target)
    {
        if(sum==target)
        {
            ret.push_back(tmp);
            return ;
        }
        if(sum>target)  
            return ;
        for(int i=index; i<nums.size(); i++)
        {
            if(i!=index && nums[i]==nums[i-1])  //跳过与前一个数字相同的数
                continue;
            sum+=nums[i];
            tmp.push_back(nums[i]);
            dfs(nums, ret, tmp, sum, i+1, target);
            sum-=nums[i];
            tmp.pop_back();
            //while(i<nums.size()-1 && nums[i]==nums[i+1])  //如果下一个数与当前数相同，则跳过
			    //i++;
        }
    }
    
    vector<vector<int>> combinationSum2(vector<int>& nums, int target) {
        vector<vector<int>> ret;
        vector<int> tmp;
        int sum=0;
        sort(nums.begin(), nums.end());
        dfs(nums, ret, tmp, sum, 0, target);
        return ret;
    }
};
