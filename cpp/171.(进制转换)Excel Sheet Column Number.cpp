/*
Related to question Excel Sheet Column Title

Given a column title as appear in an Excel sheet, return its corresponding column
number.

For example:

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28 
 
 即26进制转10进制。
 
 同理2进制转10进制:
 for(int i=0; i<str.length(); i++)
        ret=ret*2+(str[i]-'0');
        
 注意: 以'A'而不是0开头，因此要“+1”
*/

class Solution {
public:
    int titleToNumber(string s) {
        int ret=0;
        for(int i=0; i<s.length(); i++)
            ret=ret*26+s[i]-'A'+1;
        return ret;
    }
};
