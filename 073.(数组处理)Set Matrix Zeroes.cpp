/*
Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in place.

用两个vector当map分别存储有0的行和列，然后在遍历，把map中vector中有的行和列置为0
*/

class Solution {
public:
    void setZeroes(vector<vector<int>>& matrix) {
        int rows=matrix.size(), cols=matrix[0].size();    
        if(rows==0 || cols==0)
            return ;
        vector<int> rowHas0;
        vector<int> colHas0;
        for(int i=0; i<rows; i++)
        {
            for(int j=0; j<cols; j++)
            {
                if(matrix[i][j]==0)
                {
                    rowHas0.push_back(i);
                    colHas0.push_back(j);
                }
            }
        } 
        for(int i=0; i<rows; i++)
        {
            for(int j=0; j<colHas0.size(); j++)
                matrix[i][colHas0[j]]=0;
        }
        for(int i=0; i<rowHas0.size(); i++)
        {
            for(int j=0; j<cols; j++)
                matrix[rowHas0[i]][j]=0;
        }
    }
};
