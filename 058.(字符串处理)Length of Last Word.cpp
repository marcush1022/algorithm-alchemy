/***************************************************************************************/
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', 
return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

For example, 
Given s = "Hello World",
return 5.
/***************************************************************************************/

class Solution {
public:
    int lengthOfLastWord(string s) {
        int len=s.length();
        bool flag=0;
        int cnt=0;
        for(int i=len-1; i>=0; i--)
        {
            if(s[i]==' ' && flag==0)
                continue;
            if(s[i]!=' ')
            {
                flag=1;
               cnt++;
            }
            else
                 break;      
        }
        return cnt;
    }
};
