/*
Find the contiguous subarray within an array (containing at least one number) which has the largest product.

For example, given the array [2,3,-2,4],
the contiguous subarray [2,3] has the largest product = 6.

递推公式: max[i]=max[i-1]*num[i]
		OR min[i-1]*num[i]
		OR num[i];
*/

class Solution {
public:
    int maxProduct(vector<int>& nums) {
        int maxNum=nums[0], minNum=nums[0], globeMax=nums[0];
        if(nums.size()==1)
            return nums[0];
        if(nums.size()==0)
            return 0;
        for(int i=1; i<nums.size(); i++)
        {
            int a=nums[i]*maxNum;
            int b=nums[i]*minNum;
            
            maxNum=max(nums[i], max(a, b));
            minNum=min(nums[i], min(a, b));
            
            globeMax=max(globeMax, maxNum);
        }
        return globeMax;
    }
};
