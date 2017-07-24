/*
Given a string S and a string T, find the minimum window in S which will contain all the 
characters in T in complexity O(n).

For example,
S = "ADOBECODEBANC"
T = "ABC"
Minimum window is "BANC".

Note:
If there is no such window in S that covers all characters in T, return the empty string "".

If there are multiple such windows, you are guaranteed that there will always be only one 
unique minimum window in S.

这道题的要求是在给定的字符串S中找到最小的窗口使其完全包含字符串T中所有字符，如果不存在，则返回空串""。

和Longest Substring Without Repeating Characters及Substring with Concatenation of All Words
类似的思路，使用l和r两个指针维护子串，用Hash表记录T字串中出现的字符。S中，每次循环r右移1位，然后判断
r右移之后所指向的字符是否在Hash表中出现过：如果出现过，则表示在T中。此时通过计数器cnt判断T中字符是否
都出现过，如果是，则记录l和r之间子串长度，并与最短长度比较。然后逐步右移l并在Hash表中删除l指向的字符
直到计数器cnt小于T中字符数量。

注意一下，由于T中同一字符的数量可能减到负值，因此需要2重判断：先判断是否出现此字符，在判断此字符出现
的具体数量。

时间复杂度：O(l1+l2)（l1和l2分别为字符串S和T的长度）

空间复杂度：O(l2)
*/

class Solution {
public:
    string minWindow(string s, string t) {
        int lenS=s.length(), lenT=t.length();
        map<int, int> mp;
        int right=0, left=0;
        int count=0, minLen=INT_MAX, minLeft;
        for(int i=0; i<lenT; i++)
            mp[t[i]]++;
        for(int right=0; right<lenS; right++)
        {
            if(mp.find(s[right])!=mp.end())
            {
                if(--mp[s[right]]>=0)
                    count++;
            
                while(count==lenT)
                {
                    if(right-left+1<minLen)
                    {
                        minLen=right-left+1;
                        minLeft=left;
                    }
                    if(mp.find(s[left])!=mp.end())
                        if(++mp[s[left]]>0)
                            count--;
                    left++;
                }
            }
        }
        if(minLen>lenS)
            return "";
        else
            return s.substr(minLeft, minLen);
    }
};
