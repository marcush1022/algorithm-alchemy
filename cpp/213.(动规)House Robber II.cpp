/*
Note: This is an extension of House Robber.

After robbing those houses on that street, the thief has found himself a new place for 
his thievery so that he will not get too much attention. 
This time, all houses at this place are arranged in a circle. 
That means the first house is the neighbor of the last one. Meanwhile, 
the security system for these houses remain the same as for those in the previous street.

Given a list of non-negative integers representing the amount of money of each house, 
determine the maximum amount of money you can rob tonight without alerting the police.

在一个序列中取数，不能连续取，求最大和（序列首位相连）。

两种情况，第一种取第一个不取最后一个，另一种取最后一个不取第一个
*/

class Solution {
public:
    int rob(vector<int>& nums) {
        int len=nums.size();
        if(len==0)
            return 0;
        if(len==1)
            return nums[0];
        int maxAmt=0;
        int ppre=0, pre=0, cur1=0, cur2=0;
        for(int i=0; i<len-1; i++)
        {
            cur1=max(ppre+nums[i], pre);
            ppre=pre;
            pre=cur1;
        }
        ppre=0, pre=0;
        for(int i=1; i<len; i++)
        {
            cur2=max(ppre+nums[i], pre);
            ppre=pre;
            pre=cur2;
        }
        return max(cur1, cur2);
    }
};
