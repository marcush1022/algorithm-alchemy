/*
Given an array of size n, find the majority element. The majority element is the element 
that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the 
array.

类似快排的分治方法，
如果数组中存在出现次数超过一半的数字，
则以这个数字进行划分的话他一定位于数组的中间。
即：寻找划分之后位置位于数组中间的数字。
*/

class Solution {
public:
	int divide(vector<int> &nums, int left, int right)
	{
		int val=nums[right];
		for(int i=left; i<right; i++)
        {
            if(nums[i]<val)
                swap(nums[i], nums[left++]);
        }
        swap(nums[left], nums[right]);
        return left;
	}

    int majorityElement(vector<int>& nums) {
        int len=nums.size();
	int pos=len/2, ans=0;
	int left=0, right=len-1;
	while((ans=divide(nums, left, right))!=pos)
        {
            if(ans>pos)
                right=ans-1;
            else
                left=ans+1;
        }
        return nums[pos];
    }
};
