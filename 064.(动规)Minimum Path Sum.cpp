/*
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom 
right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

数字矩阵求sum为最小的路径
当i为0时，每个位置的和为左边的数加grid对应值，当j为0时和为上边的数加grid对应值，其他情况取最小值
并与grid当前值相加即可
*/

class Solution {
public:
    int minPathSum(vector<vector<int>>& grid) {
        int m=grid.size(), n=grid[0].size();
        if(m==0 || n==0)
            return 0;
        vector<int> arr(n);
        arr[0]=grid[0][0];
        for(int j=1; j<n; j++)
            arr[j]=arr[j-1]+grid[0][j];
        for(int i=1; i<m; i++)
        {
            for(int j=0; j<n; j++)
            {
                if(j==0)
                    arr[0]+=grid[i][0];
                else
                    arr[j]=min(arr[j-1], arr[j])+grid[i][j];
            }
        }
        return arr[n-1];
    }
};
