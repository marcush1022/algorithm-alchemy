/*
Implement strStr().

Returns the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
*/

/*string find(): 
Return :
The position of the first character of the first match.
If no matches were found, the function returns string::npos.*/

I. 
class Solution {
public:
    int strStr(string str, string substr) {
        int found=str.find(substr);
        return (found==string::npos)? -1 : found;
    }
};


II. 
class Solution {
public:
    int strStr(string str, string sub) {
        int len1=str.length(), len2=sub.length();
        if(len1==len2 && len1==0)
            return 0;
        int index, j;
        for(int i=0; i<len1; i++)
        {
            index=i, j=0;
            while(j<len2 && str[index]==sub[j])
            {
                ++index;
                ++j;
            }
            if(j==len2)
                return i;
        }
        return -1;
    }
};
