// package twoPointers
// dir_path golang/twoPointers
// name 083.easy_Remove_Duplicates_from_Sorted_List_test

package twoPointers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	head := createListFromSlice83([]int{1, 1, 2, 3, 3})
	deleteDuplicates(head)
	str := convertListToString(head)
	assert.Equal(t, fmt.Sprint([]int{1, 2, 3}), str)

	head = createListFromSlice83([]int{1, 1, 2})
	deleteDuplicates(head)
	str = convertListToString(head)
	assert.Equal(t, fmt.Sprint([]int{1, 2}), str)
}

func createListFromSlice83(slice []int) (head *ListNode) {
	head = &ListNode{
		Val:  slice[0],
		Next: nil,
	}
	curNode := head
	for i := 1; i < len(slice); i++ {
		nextNode := &ListNode{
			Val:  slice[i],
			Next: nil,
		}
		curNode.Next = nextNode
		curNode = nextNode
	}
	return head
}

func convertListToString(head *ListNode) string {
	arr := make([]int, 0)
	for ; head != nil; head = head.Next {
		arr = append(arr, head.Val)
	}
	return fmt.Sprint(arr)
}
