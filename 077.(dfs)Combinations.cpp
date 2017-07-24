/*
Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

For example,
If n = 4 and k = 2, a solution is:

[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
典型回溯问题，需要序列中数字递增，在递归调用dfs时i+1即可
*/

class Solution {
public:
    void dfs(vector<vector<int>> &ret, vector<int> tmp, int k, int n, int index)
    {
        if(tmp.size()==k)
        {
            ret.push_back(tmp);
            return ;
        }    
        for(int i=index; i<=n; i++)
        {
            tmp.push_back(i);
            dfs(ret, tmp, k, n, i+1);
            tmp.pop_back();
        }
    }
    
    vector<vector<int>> combine(int n, int k) {
        vector<vector<int>> ret;
        vector<int> tmp;
        dfs(ret, tmp, k, n, 1);
        return ret;
    }
};
