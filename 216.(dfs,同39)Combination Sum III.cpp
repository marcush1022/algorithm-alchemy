/*
Find all possible combinations of k numbers that add up to a number n, given that only numbers from 1 to 9 can be
used and each combination should be a unique set of numbers.


Example 1:

Input: k = 3, n = 7

Output:

[[1,2,4]]

Example 2:

Input: k = 3, n = 9

Output:

[[1,2,6], [1,3,5], [2,3,4]]

典型dfs
*/

class Solution {
    vector<vector<int>> ret;
public:
    void dfs(vector<int> &tmp, int k, int target, int sum, int index)
    {
        if(tmp.size()>k || sum>target)
            return ;
        if(tmp.size()==k && sum==target)
            ret.push_back(tmp);
        for(int i=index; i<=9; i++)
        {
            tmp.push_back(i);
            sum+=i;
            dfs(tmp, k, target, sum, i+1);
            sum-=i;
            tmp.pop_back();
        }
    }
    
    vector<vector<int>> combinationSum3(int k, int n) {
        vector<int> tmp;
        dfs(tmp, k, n, 0, 1);
        return ret;
    }
};
