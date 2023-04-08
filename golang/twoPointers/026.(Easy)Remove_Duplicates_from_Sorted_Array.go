// package twoPointers
// name 26.Remove_Duplicates_from_Sorted_Array

package twoPointers

import "fmt"

/**
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once.
The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums.
More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result.
It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
If all assertions pass, then your solution will be accepted.



Example 1:

Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
Example 2:

Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).


Constraints:

1 <= nums.length <= 3 * 104
-100 <= nums[i] <= 100
nums is sorted in non-decreasing order.

由于数组已经排序，所以重复的元素一定连在一起，找出它们并不难。但如果毎找到一个重复元素就立即原地删除它，
由于数组中删除元素涉及数据搬移，整个时间复杂度是会达到 O(N^2)。

高效解决这道题就要用到快慢指针技巧：

我们让慢指针 slow 走在后面，快指针 fast 走在前面探路，找到一个不重复的元素就赋值给 slow 并让 slow 前进一步。

这样，就保证了 nums[0]..nums[slow-1] 都是无重复的元素，当 fast 指针遍历完整个数组 nums 后，nums[0]..nums[slow-1] 就是整个数组去重之后的结果。
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slowIndex := 1
	for fastIndex := 1; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != nums[fastIndex-1] {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	fmt.Println(">>>>>> slowIndex", slowIndex)
	return slowIndex
}
