/*
Given a set of candidate numbers (C) (without duplicates) and a target number (T), find all unique 
combinations in C where the candidate numbers sums to T.

The same repeated number may be chosen from C unlimited number of times.

Note:
All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
For example, given candidate set [2, 3, 6, 7] and target 7, 
A solution set is: 
[
  [7],
  [2, 2, 3]
]

dfs模板级题目...
取数字和为给定值，可以重复，求有多少种组合

Note: 回溯算法递归算法框架：
int a[n];
 try(int i)
 {
    if(i>n)
        输出结果;
      else
     {
        for(j = 下界; j <= 上界; j=j+1)  // 枚举i所有可能的路径
        {
            if(fun(j))                 // 满足限界函数和约束条件
              {
                 a[i] = j;
               ...                         // 其他操作
                 try(i+1);
               回溯前的清理工作（如a[i]置空值等）;
               }
          }
      }
 }
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
        if(sum>target)   //不要漏掉sum>target情况,不然死循环
            return ;
        for(int i=index; i<nums.size(); i++)
        {
            sum+=nums[i];
            tmp.push_back(nums[i]);
            dfs(nums, ret, tmp, sum, i, target);
            sum-=nums[i];
            tmp.pop_back();
        }
    }
    
    vector<vector<int>> combinationSum(vector<int>& nums, int target) {
        vector<vector<int>> ret;
        vector<int> tmp;
        int sum=0;
        dfs(nums, ret, tmp, sum, 0, target);
        return ret;
    }
};
