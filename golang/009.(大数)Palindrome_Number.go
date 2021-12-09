// package golang
// name 009.(大数)Palindrome_Number

package golang

import "math"

func IsPalindrome(n int) bool {
	if n < 0 {
		return false
	}

	if n > math.MaxInt32 {
		return false
	}

	length := 1 // n 一共有多少位
	for n/length >= 10 {
		length *= 10
	}

	for n > 0 {
		// 最左等于最右
		if n/length != n%10 {
			return false
		}
		n = n % length        // 去掉最左位
		n = n / 10            // 去掉最右位
		length = length / 100 // 减两位
	}
	return true
}
