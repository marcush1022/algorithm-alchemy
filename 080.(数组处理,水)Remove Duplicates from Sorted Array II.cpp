/*
Follow up for "Remove Duplicates":
What if duplicates are allowed at most twice?

For example,
Given sorted array nums = [1,1,1,2,2,3],

Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 
and 3. It doesn't matter what you leave beyond the new length.
*/

class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int len=nums.size();
        if(len<=2)
            return len;
        int ret=0;
        for(int i=0; i<len-2; i++)
        {
            if(nums[i]!=nums[i+1] || nums[i]!=nums[i+2])
                nums[ret++]=nums[i];
        }
        nums[ret++]=nums[len-2];
        nums[ret++]=nums[len-1];
        return ret;
    }
};

class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int len=nums.size();
        map<int, int> mp;
        int ret=0;
        for(int i=0; i<len; i++)
        {
            if(++mp[nums[i]]<=2)
                nums[ret++]=nums[i];
        }
        return ret;
    }
};
