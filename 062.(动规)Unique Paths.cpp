/*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in 
the diagram below).

The robot can only move either down or right at any point in time. The robot 
is trying to reach the bottom-right corner of the grid (marked 'Finish' in 
the diagram below).

How many possible unique paths are there?

dp[i][j]=dp[i-1][j]+dp[i][j-1]
*/

class Solution {
public:
    int uniquePaths(int m, int n) {
        if(m==0 || n==0)
            return 0;
        int grid[m][n];
        for(int i=0; i<m; i++)
        {
            for(int j=0; j<n; j++)
            {
                if(i==0 && j==0)
                    grid[i][j]=1;
                else if(i==0 || j==0)
                    grid[i][j]=1;
                else
                    grid[i][j]=grid[i-1][j]+grid[i][j-1];
            }
        }
        return grid[m-1][n-1];
    }
};

o(n)
class Solution {
public:
    int uniquePaths(int m, int n) {
        if(m==0 || n==0)
            return 0;
        vector<int> arr(n, 1);
        for(int i=1; i<m; i++)
        {
            for(int j=1; j<n; j++)
                arr[j]+=arr[j-1];
        }
        return arr[n-1];
    }
};
