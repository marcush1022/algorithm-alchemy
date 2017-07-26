/*
Given two sorted integer arrays nums1 and nums2, 
merge nums2 into nums1 as one sorted array.

Note:
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold 
additional elements from nums2.
The number of elements initialized in nums1 and nums2 are m and n respectively.

从两数组和的尾部开始添加元素, 直到其中一个数组遍历完，这时index停在的位置等于较长的数组中剩余元素的数量
*/

class Solution {
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
        int i=m-1, j=n-1;
        int index=m+n-1;
        while(i>=0 && j>=0)
        {
            if(nums1[i]<nums2[j])
                nums1[index--]=nums2[j--];
            else
                nums1[index--]=nums1[i--];
        }
        while(j>=0)
        {
            nums1[index--]=nums2[j--];
        }
    }
};
