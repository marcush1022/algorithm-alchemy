/*
Given an array of n positive integers and a positive integer s, find the minimal 
length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

For example, given the array [2,3,1,2,4,3] and s = 7,
the subarray [4,3] has the minimal length under the problem constraint.

采用滑动窗口的思想：
滑动窗口 [left, right] 初始大小为0，滑动的规则如下：

1. 若元素之和 < s，窗口右边沿向右延伸，直到 元素之和≥s 为止。
2. 若元素之和 ≥ s，更新最小长度。然后窗口左边沿右移一位（即移除窗口中第一个元素），重复第1步。

从第一个元素开始累加和，
当超过给定的s时，此时从最开始的ori点处
开始，用sum减去，直到此时的sum再次小于
s。注意，期间一旦sum >= s时，求最小的
元素数。
*/

class Solution {
public:
    int minSubArrayLen(int s, vector<int>& nums) {
        int len=nums.size();
        if(len==0)
            return 0;
        int left=0, right=0;
        int minLen=len+1, sum=0;
        
        while(left<len && right<len)
        {
            while(sum<s && right<len)
                sum+=nums[right++];
            while(sum>=s && left<=right)
            {
                minLen=min(minLen, right-left);
                sum-=nums[left++];
            }
        }
        if(minLen==len+1)
            return 0;
        else
            return minLen;
    }
};
