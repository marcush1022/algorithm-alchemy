/*
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.

click to show more practice.

More practice:
If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

递推公式: dp[i]=max(dp[i-1]+num[i], num[i])
*/

class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int sum=0, maxSum=INT_MIN;
        for(int i=0; i<nums.size(); i++)
        {
            if(sum<0)
                sum=nums[i];
            else
                sum+=nums[i];
            maxSum=max(sum, maxSum);
        }
        return maxSum;
    }
};
