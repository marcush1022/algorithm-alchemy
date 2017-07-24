/*
Given a set of distinct integers, nums, return all possible subsets.

Note: The solution set must not contain duplicate subsets.

For example,
If nums = [1,2,3], a solution is:

[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

使用一个变量来记录路径，每遍历到一个元素即表示找到一条路径，将其加入子集中。 
例如数组[1,2,3] 
从1开始递归查询2,3，对于2，继续向下搜索，搜索完后将2删除。 
*/

class Solution {
public:
    void dfs(vector<vector<int>> &ret, vector<int> tmp, vector<int> nums, int index)
    {
        if(index==nums.size())
            return ;   
        for(int i=index; i<nums.size(); i++)
        {
            tmp.push_back(nums[i]);
            ret.push_back(tmp);
            dfs(ret, tmp, nums, i+1);
            tmp.pop_back();
        }
    }
    
    vector<vector<int>> subsets(vector<int>& nums) {
        int len=nums.size();
        vector<vector<int>> ret;
        vector<int> tmp;
        ret.push_back(tmp);
        sort(nums.begin(), nums.end());
        dfs(ret, tmp, nums, 0);
        return ret;
    }
};
