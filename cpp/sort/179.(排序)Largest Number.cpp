/*
Given a list of non negative integers, arrange them such that they form the largest number.

For example, given [3, 30, 34, 5, 9], the largest formed number is 9534330.

Note: The result may be very large, so you need to return a string instead of an integer.
*/

class Solution {
public:
    bool isBigger(int n1, int n2)
    {
        string str1=to_string(n1);
        string str2=to_string(n2);
        return (str1+str2)>(str2+str1);
    }
    
    int partition(vector<int> &nums, int left, int right)
    {
        int val=nums[right];
        for(int i=left; i<right; i++)
        {
            if(isBigger(nums[i], val))
                swap(nums[i], nums[left++]);
        }
        swap(nums[left], nums[right]);
        return left;
    }
    
    void quickSort(vector<int> &nums, int left, int right)
    {
        if(left<right)
        {
            int pos=partition(nums, left, right);
            quickSort(nums, left, pos-1);
            quickSort(nums, pos+1, right);
        }
    }
    
    string largestNumber(vector<int>& nums) {
        int len=nums.size();
        int left=0, right=len-1;
        string ret="";
        quickSort(nums, left, right);
        for(int i=0; i<len; i++)
            ret+=to_string(nums[i]);
        if(ret[0]=='0')
            return "0";
        else
            return ret;
    }
};
