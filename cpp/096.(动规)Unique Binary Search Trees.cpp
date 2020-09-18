/*
Given n, how many structurally unique BST's (binary search trees) that store values 1...n?

For example,
Given n = 3, there are a total of 5 unique BST's.

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

用递推的方法，dp[i]表示i个数能够构成的不同的BST数量，
初始化边界值，dp[0]=dp[1]=1，
当i>=3时，当以j为根节点，左边i-1节点构成子树数量为dp[j-1]，右边i-j节点构成子树数量为dp[i-j]时，
以j为根构成的总数量为dp[i-1]*dp[i-j]

j可以取1~i中的任意一个值，把这些所有计算出来的总数相加就是dp[i]的值
所以 for(int j = 1; j <= i; j++) {
dp[i] += dp[j-1] * dp[i-j];
}
最后返回的值是dp[n]的值，表示1~n能组成的BST的个数
*/

class Solution {
public:
    int numTrees(int n) {
        vector<int> dp(n+1);
        dp[0]=dp[1]=1;
        for(int i=2; i<=n; i++)
        {
            for(int j=1; j<=i; j++)
                dp[i]+=dp[j-1]*dp[i-j];
        }
        return dp[n];
    }
};
