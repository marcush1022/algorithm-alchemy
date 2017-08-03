/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters 
and ignoring cases.

For example,
"A man, a plan, a canal: Panama" is a palindrome.
"race a car" is not a palindrome.

Note:
Have you consider that the string might be empty? This is a good question to ask during an 
interview.

For the purpose of this problem, we define empty string as valid palindrome.
*/

class Solution {
public:
    bool isPalindrome(string s) {
        int len=s.length();
        if(len==0)
            return true;
        int left=0, right=len-1;
        while(left<right)
        {
            if(!isalpha(s[left]) && !isdigit(s[left]))
                left++;
            else if(!isalpha(s[right]) && !isdigit(s[right]))  //#######
                right--;
            else if(tolower(s[left])==tolower(s[right]))
            {
                ++left;
                --right;
            }
            else 
                return false;
        }
        return true;
    }
};
