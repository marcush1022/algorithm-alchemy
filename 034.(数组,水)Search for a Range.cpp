/*
Given an array of integers sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

For example,
Given [5, 7, 7, 8, 8, 10] and target value 8,
return [3, 4].
*/

class Solution {
public:
    vector<int> searchRange(vector<int>& nums, int target) {
        int len=nums.size();
        vector<int> ret;
        int left=-1, right=-1;
        for(int i=0; i<len; i++)
        {
            if(nums[i]==target)
            {
                left=i;
                while(nums[i]==target)
                    i++;
                right=i-1;
            }    
        }
        ret.push_back(left);
        ret.push_back(right);
        return ret;
    }
};
