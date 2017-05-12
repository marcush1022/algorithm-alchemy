/***********************************************************************************************************/
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
/***********************************************************************************************************/

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        int len=nums.size();
		vector<int> ret;
		if(len==0)
			return ret;
		map<int, int> map;
		for(int i=0; i<len; i++)
		{
			//先判断余下部分是否存在，不然会重复
			//例如: [3,2,4], 6
			if(map.find(target-nums[i])!=map.end())
			{	
				ret.push_back(i);
				ret.push_back(map[target-nums[i]]);
				return ret;
			}
			map[nums[i]]=i;
		}
		return ret;
    }
};
