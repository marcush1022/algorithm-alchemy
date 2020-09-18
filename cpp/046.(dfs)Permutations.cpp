/*
Given a collection of distinct numbers, return all possible permutations.

For example,
[1,2,3] have the following permutations:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

用index记录已确定的位置，每次确定一个元素的最终位置，当所有元素位置都确定时(index=nums.size()-1)输出
第一趟所有元素都和自己交换，所以输出原序列

注意：第一个字符与后面的某个位置的字符发生交换后，需要再次发生交换，不然顺序就会被打乱。
举个例子，在字符串abc中，在把第一个字符看成是a，后面的字符b、c看成一个整体的时候，
abc这个相对顺序不能改变，所以当b与c发生交换变成了acb之后，需要再次交换两个字符，重新回到abc。
*/

class Solution {
public:
    void dfs(vector<int> nums, vector<vector<int>> &ret, int index)
    {
        if(index==nums.size()-1)
        {
            ret.push_back(nums);
            return ;
        }
        for(int i=index; i<nums.size(); i++)
        {
            swap(nums[i], nums[index]);
            dfs(nums, ret, index+1);
            swap(nums[i], nums[index]);
        }
    }
    
    vector<vector<int>> permute(vector<int>& nums) {
        vector<vector<int>> ret;
        int len=nums.size();
        if(len==0)
            return ret;
        dfs(nums, ret, 0);
        return ret;
    }
};
