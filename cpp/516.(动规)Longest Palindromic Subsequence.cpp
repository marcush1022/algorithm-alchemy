/*
Given a string s, find the longest palindromic subsequence's length in s. You may assume that the maximum length of s is 1000.

Example 1:
Input:

"bbbab"
Output:
4
One possible longest palindromic subsequence is "bbbb".
Example 2:
Input:

"cbbd"
Output:
2
One possible longest palindromic subsequence is "bb".

用动态规划，申请一个len*len的数组dp，将dp[i][i]初始化为1，dp[i][j]表示下标
i~j中包括的最长回文子序列长度。
从后向前遍历，只用数组的右上半部分。
最后返回0~len-1中最长子序列长度
*/

class Solution {
public:
    int longestPalindromeSubseq(string s) {
        int len=s.length();
        vector<vector<int>> dp(len, vector<int>(len));
        for(int i=len-1; i>=0; i--)
        {
            dp[i][i]=1;
            for(int j=i+1; j<len; j++)
            {
                if(s[i]==s[j])
                    dp[i][j]=dp[i+1][j-1]+2;
                else
                    dp[i][j]=max(dp[i][j-1], dp[i+1][j]);
            }
        }
        return dp[0][len-1];
    }
};
