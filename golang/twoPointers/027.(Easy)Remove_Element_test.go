// package twoPointers
// dir_path golang/twoPointers
// name 027.(Easy)Remove_Element_test.go

package twoPointers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeElement(t *testing.T) {
	assert.Equal(t, "", checker27(3, []int{3, 2, 2, 3}, []int{2, 2, 0, 0}, 2))
	assert.Equal(t, "", checker27(2, []int{0, 1, 2, 2, 3, 0, 4, 2}, []int{0, 1, 3, 0, 4, 0, 0, 0}, 5))
}

func checker27(target int, original, expRes []int, expNum int) string {
	actNum, actRes := removeElement(original, target)
	if actNum != expNum {
		return fmt.Sprintf("actNum != expNum, actNum: %v, expNum: %v", actNum, expNum)
	}
	fmt.Println(">>>>>>> actNum, actRes", actNum, actRes)

	for i := 0; i < actNum; i++ {
		if actRes[i] != expRes[i] {
			return fmt.Sprintf("actRes[i] != expRes[i], i: %v, actRes[i]: %v, expRes[i]: %v", i, actRes[i], expRes[i])
		}
	}
	return ""
}
