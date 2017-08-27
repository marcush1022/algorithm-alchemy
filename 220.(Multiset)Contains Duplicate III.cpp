/*
Given an array of integers, find out whether there are two distinct indices i and j in the array 
such that the absolute difference between nums[i] and nums[j] is at most t and the absolute 
difference between i and j is at most k.

Note: 思路是维护一个大小不超过 k 的窗口，每加一个数前判断这个窗口中有没有差大于 t 的数。（nlogn)
主要考虑的是数据结构。

set和multiset会根据特定的排序准则，自动将元素进行排序。不同的是后者允许元素重复而前者不允许。

假设当前遍历到nums[i]，另一个数大小为x，我们要在二叉排序树中找到是否有满足:
nums[i] - t <= x <= nums[i] + t　的位置。在C++的STL中提供了二叉排序树的数据结构set，并且
其提供了一个函数lower_bound，可以查找树中第一个大于等于某值的位置，利用这个函数可以找到第一个大于等
于nums[i]-t的指针，然后再判断其值是否满足条件即可。
*/

class Solution {
public:
    bool containsNearbyAlmostDuplicate(vector<int>& nums, int k, int t) {
        if(nums.size()==0)
			return false;
		multiset<long long> window;
		for(int i=0; i<nums.size(); i++)
		{
			if(i>k)
				window.erase(window.find(nums[i-k-1]));
			auto it=window.lower_bound((long long)nums[i]-t);
			if(it!=window.end() && *it<=(long long)nums[i]+t)
				return true;
			window.insert(nums[i]);
		}
		return false;
    }
};
