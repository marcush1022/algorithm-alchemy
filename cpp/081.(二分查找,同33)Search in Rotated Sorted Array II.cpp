/*
Follow up for "Search in Rotated Sorted Array":
What if duplicates are allowed?

Would this affect the run-time complexity? How and why?
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

Write a function to determine if a given target is in the array.

The array may contain duplicates.

旋转数组分2种情况：
1. 最左大于等于中间
2. 最左小于等于中间
旋转数组与普通数组二分查找的区别是只需要进行一半的查找
注意11131和13111的情况，需要left++跳过重复数字
*/

class Solution {
public:
    bool search(vector<int>& nums, int target) {
        int len=nums.size();
        if(len==0)
            return false;
        int left=0, right=len-1;
        while(left<right)
        {
            int mid=(left+right)/2;
            if(nums[mid]==target)
                return true;
            if(nums[mid]<nums[left])
            {
                if(nums[mid]<target && nums[right]>=target)
                    left=mid+1;
                else
                    right=mid;
            }
            else if(nums[mid]>nums[left])
            {
                if(nums[mid]>target && nums[left]<=target)
                    right=mid;
                else
                    left=mid+1;
            }
            else
                left++;
        }
        return nums[left]==target? true : false;
    }
};
