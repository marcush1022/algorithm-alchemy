/*
Find the kth largest element in an unsorted array. 
Note that it is the kth largest element in the sorted order, not the kth distinct element.

For example,
Given [3,2,1,5,6,4] and k = 2, return 5.

Note: 
You may assume k is always valid, 1 ≤ k ≤ array's length.

用快排的分治partiton，题意为求第k大的元素，即partition之后该元素位置为len-k
*/

class Solution {
public:
    int partition(vector<int> &nums, int left, int right)
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
    
    int findKthLargest(vector<int>& nums, int k) {
        int len=nums.size();
        int left=0, right=len-1;
        int pos=len-k, ans=0;
        while((ans=partition(nums, left, right))!=pos)
        {
            if(ans>pos)
                right=ans-1;
            else
                left=ans+1;
        }
        return nums[pos];
    }
};
