// package twoPointers
// dir_path golang/twoPointers
// name 83.easy_Remove_Duplicates_from_Sorted_List

package twoPointers

/**
83. Remove Duplicates from Sorted List

Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

Example 1:

Input: head = [1,1,2]
Output: [1,2]

Example 2:

Input: head = [1,1,2,3,3]
Output: [1,2,3]

Constraints:

The number of nodes in the list is in the range [0, 300].
-100 <= Node.val <= 100
The list is guaranteed to be sorted in ascending order.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	for fast := slow.Next; fast != nil; fast = fast.Next {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
	}
	slow.Next = nil
	return head
}
