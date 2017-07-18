/*
Follow up for "Unique Paths":

Now consider if some obstacles are added to the grids. How many unique paths 
would there be?

An obstacle and empty space is marked as 1 and 0 respectively in the grid.

For example,
There is one obstacle in the middle of a 3x3 grid as illustrated below.

[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
The total number of unique paths is 2.

Note: Beware of [0,1,0]

跟上题的区别是注意矩阵中的障碍，其他类似
*/

class Solution {
public:
    int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {
        int m=obstacleGrid.size(), n=obstacleGrid[0].size();
		if(m==0 || n==0)
			return 0;
		vector<int> arr(n);
		for(int j=0; j<n && !obstacleGrid[0][j]; j++)
			arr[j]=1;
		for(int i=1; i<m; i++)
		{
			if(obstacleGrid[i][0])
				arr[0]=0;
			for(int j=1; j<n; j++)
			{
				if(obstacleGrid[i][j])
					arr[j]=0;
				else
					arr[j]+=arr[j-1];
			}
		}
		return arr[n-1];
    }
};





