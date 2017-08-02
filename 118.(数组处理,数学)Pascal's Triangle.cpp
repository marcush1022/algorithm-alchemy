/*
Given numRows, generate the first numRows of Pascal's triangle.

For example, given numRows = 5,
Return

[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

被题目中的输出结果误导了，输出结果应该是这样：
0---[1],
1---[1,1],
2---[1,2,1],
3---[1,3,3,1],
...

每行先加一个1，然后通过上一行计算结果即可
当i>=1时需要在尾部再加一个1
*/

class Solution {
public:
    vector<vector<int>> generate(int n) {
        vector<vector<int>> ret(n);
        for(int i=0; i<n; i++)
        {
            ret[i].push_back(1);
            for(int j=1; j<i; j++)
                ret[i].push_back(ret[i-1][j-1]+ret[i-1][j]);
            if(i)
                ret[i].push_back(1);
        }
        return ret;
    }
};
