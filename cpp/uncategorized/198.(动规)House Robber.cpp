/*
You are a professional robber planning to rob houses along a street. 
Each house has a certain amount of money stashed, the only constraint stopping you
from robbing each of them is that adjacent houses have security system connected 
and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, 
determine the maximum amount of money you can rob tonight without alerting the police.

整数数列，不能连续取，求可取得的最大值。
递推公式：
dp[i]=max(ppre+nums[i], pre)
*/

class Solution {
public:
    int rob(vector<int>& nums) {
        int ppre=0, pre=0, cur=0;
        int len=nums.size();
        for(int i=0; i<len; i++)
        {
            cur=max(ppre+nums[i], pre);
            ppre=pre;
            pre=cur;
        }
        return cur;
    }
};
