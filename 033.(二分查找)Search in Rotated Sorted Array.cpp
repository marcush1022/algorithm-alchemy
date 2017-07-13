/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

旋转数组分2种情况：
1. 最左大于中间
2. 最左小于中间
旋转数组与普通数组二分查找的区别是只需要进行一半的查找，主要难点在边界值处理
*/

class Solution {
public:
    int search(vector<int>&nums, int target) {
        int len=nums.size();
        if(len==0)
            return -1;
        int left=0, right=len, mid=0;
        while(left<right)
        {
            mid=(left+right)/2;
            if(nums[mid]==target)
                return mid;
            if(nums[mid]>=nums[left])
            {
                if(nums[left]<=target && nums[mid]>target)
                    right=mid;
                else
                    left=mid+1;
            }
            else if(nums[mid]<=nums[left])
            {
                if(nums[mid]<target && nums[right-1]>=target)
                    left=mid+1;
                else
                    right=mid;
            }
        }
        return -1;
    }
};
