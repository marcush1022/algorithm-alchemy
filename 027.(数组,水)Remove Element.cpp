/*
Given an array and a value, remove all instances of that value in place and return the new length.

Do not allocate extra space for another array, you must do this in place with constant memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.

Example:
Given input array nums = [3,2,2,3], val = 3

Your function should return length = 2, with the first two elements of nums being 2.
//同上题，就是快慢指针
*/

class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        int len=nums.size();
        if(len==0)
            return 0;
        int ret=0;
        for(int i=0; i<len; i++)
        {
            if(nums[i]!=val)
                nums[ret++]=nums[i];
        }
        return ret;
    }
};
