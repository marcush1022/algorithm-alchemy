/*
Given a non-negative integer represented as a non-empty array of digits, plus 
one to the integer.

You may assume the integer do not contain any leading zero, except the number
0 itself.

The digits are stored such that the most significant digit is at the head of 
the list.

额外申请o(n)空间存储结果
*/

class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {
        int len=digits.size();
        vector<int> tmp=digits;
        reverse(tmp.begin(), tmp.end());
        int carry=0, i=0;
        carry=(tmp[i]+1)/10;
        tmp[i]=(tmp[i]+1)%10;
        i++;
        while(carry && i<len)
        {
            carry=(tmp[i]+carry)/10;
            tmp[i]=(tmp[i]+1)%10;
            i++;
        }
        if(i==len && carry)
            tmp.push_back(1);
        reverse(tmp.begin(), tmp.end());
        return tmp;
    }
};
