/********************************************************************************************************/
Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? 
Find all unique triplets in the array which gives the sum of zero.
/********************************************************************************************************/

class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        	int len=nums.size();
		vector<vector<int>> ret;
		vector<int> tmp;
		if(len==0)
			return ret;
		int left, right;
		int target=0;
		sort(nums.begin(), nums.end());
		for(int i=0; i<len; i++)
		{
			left=i+1, right=len-1;
			while(left<right)
			{
				int sum=nums[left]+nums[right]+nums[i];
				if(sum==target)
				{
                    			tmp.clear();
					tmp.push_back(nums[i]);
					tmp.push_back(nums[left]);
					tmp.push_back(nums[right]);
					ret.push_back(tmp);
                    			++left;
                    			--right;
                   			while(left<right && nums[left]==nums[left-1])
					    ++left;
				   	while(left<right && nums[right]==nums[right+1])
					    --right;
				}
				else if(sum<target)
					++left;
				else
					--right;
			}
			while(i<len && nums[i]==nums[i+1])
				++i;
		}
		return ret;
    }
};
