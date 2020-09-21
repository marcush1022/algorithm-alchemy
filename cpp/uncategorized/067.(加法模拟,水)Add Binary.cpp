/*
Given two binary strings, return their sum (also a binary string).

For example,
a = "11"
b = "1"
Return "100".
*/

解法1：只有2种情况会改变进位的值：全为1或全为0
class Solution {
public:
    string addBinary(string a, string b) {
        int len1=a.length(), len2=b.length();
        reverse(a.begin(), a.end());
        reverse(b.begin(), b.end());
        string res="";
        int carry=0;
        for(int i=0; i<len1 || i<len2; i++)
        {
            if(i>len1)
                a+='0';
            if(i>len2)
                b+='0';
            if(a[i]=='1' && b[i]=='1')
            {
                res+=carry+'0';
                carry=1;
            }
            else if(a[i]=='1' || b[i]=='1')
            {
                res+=!carry+'0';
            }
            else
            {
                res+=carry+'0';
                carry=0;
            }
        }
        if(carry)
            res+='1';
        reverse(res.begin(), res.end());
        return res;
    }
};

解法2: 同样适用于十进制
class Solution {
public:
    string addBinary(string a, string b) {
        int len1=a.length(), len2=b.length();
        reverse(a.begin(), a.end());
        reverse(b.begin(), b.end());
        string res="";
        int i=0, tmp=0, carry=0;
        for(i=0; i<len1 && i<len2; i++)
        {
            tmp=a[i]-'0'+b[i]-'0'+carry;
            carry=tmp/2;
            res+=tmp%2+'0';
        }
        while(i<len1)
        {
            tmp=a[i]-'0'+carry;
            carry=tmp/2;
            res+=tmp%2+'0';
            ++i;
        }
        while(i<len2)
        {
            tmp=b[i]-'0'+carry;
            carry=tmp/2;
            res+=tmp%2+'0';
            ++i;
        }
        if(carry)
            res+='1';
        reverse(res.begin(), res.end());
        return res;
    }
};
