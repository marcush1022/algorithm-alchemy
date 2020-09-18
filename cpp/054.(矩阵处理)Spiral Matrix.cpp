/*
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

For example,
Given the following matrix:

[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
You should return [1,2,3,6,9,8,7,4,5].
*/

class Solution {
public:
    vector<int> spiralOrder(vector<vector<int>>& matrix) {
        
		vector<int> ret;
		if(matrix.empty())
			return ret;
		int start, col, row, index;
		start=0;
		row=matrix.size();
		col=matrix[0].size();
		
		while(start*2<col && start*2<row)
		{
			getSpiralOrder(matrix, col, row, start, ret);
			start++;
		}
		
		return ret;
    }
	
	void getSpiralOrder(vector<vector<int>>& matrix, int col, int row, int start, vector<int>& ret)
	{
		int endRow=row-start-1;
		int endCol=col-start-1;
		
		for(int i=start; i<= endCol; i++)
		{
			ret.push_back(matrix[start][i]);
		}
		
		if(start<endRow)
		{
			for(int i=start+1; i<= endRow; i++)
			{
				ret.push_back(matrix[i][endCol]);
			}
		}
		
		if(start<endRow && start<endCol)
		{
			for(int i=endCol-1; i>=start; i--)
			{
				ret.push_back(matrix[endRow][i]);
			}
		}
		
		if(start<endCol && start<endRow-1)
		{
			for(int i=endRow-1; i>=start+1; i--)
			{
				ret.push_back(matrix[i][start]);
			}
		}
		
	}
};
