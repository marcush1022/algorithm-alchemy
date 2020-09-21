/********************************************************************************************************/
Given an array S of n integers, are there elements a, b, c, and d in S such that a + b + c + d = target? 
Find all unique quadruplets in the array which gives the sum of target.

Note: The solution set must not contain duplicate quadruplets.

For example, given array S = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
/********************************************************************************************************/

class Solution {
public:
    vector<vector<int>> fourSum(vector<int>& nums, int target) {
        	int len=nums.size();
		vector<vector<int>> ret;
		if(len==0 || len<4)
			return ret;
		
		std::sort(nums.begin(), nums.end());
		
		for(int i=0; i<len-3; i++)
		{
			for(int j=i+1; j<len-2; j++)
			{
				int left=j+1;
				int right=len-1;
				while(left<right)
				{
					int sum=nums[j]+nums[left]+nums[right]+nums[i];
					if(sum<target)
						left++;
					else if(sum>target)
						right--;
					else
					{
						ret.push_back({nums[i], nums[j], nums[left], nums[right]});
						left++;
						right--;
						while(nums[left]==nums[left-1] && left<right)
							left++;
						while(nums[right]==nums[right+1] && left<right)
							right++;
					}
				}
				while(nums[j]==nums[j+1] && j<len-2)
					j++;
			}
			while(nums[i]==nums[i+1] && i<len-3)
					i++;
		}
		return ret;
    }
};
