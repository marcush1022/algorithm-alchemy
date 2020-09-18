/*
Given an array of integers that is already sorted in ascending order, find two numbers 
such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to 
the target, where index1 must be less than index2. Please note that your returned 
answers (both index1 and index2) are not zero-based.

You may assume that each input would have exactly one solution and you may not use 
the same element twice.

Input: numbers={2, 7, 11, 15}, target=9
Output: index1=1, index2=2
*/

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        int len=nums.size();
        vector<int> ret;
        if(len==0)
            return ret;
        int left=0, right=len-1;
        while(left<right)
        {
            int tmp=nums[left]+nums[right];
            if(tmp==target)
            {
                ret.push_back(left+1);
                ret.push_back(right+1);
                return ret;
            }
            else if(tmp<target)
                ++left;
            else
                --right;
        }
        return ret;
    }
};
