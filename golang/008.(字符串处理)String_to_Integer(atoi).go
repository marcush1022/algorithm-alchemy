// package golang
// name 008.(字符串处理)String_to_Integer(atoi)

package golang

import (
	"errors"
	"fmt"
)

/*
Implement atoi to convert a string to an integer.

Hint: Carefully consider all possible input cases. If you want a challenge,
please do not see below and ask yourself what are the possible input cases.

Notes: It is intended for this problem to be specified vaguely (ie, no given
input specs). You are responsible to gather all the input requirements up
front.
*/

const intSize = 32 << (^uint(0) >> 63)

var ErrSyntax = errors.New("invalid syntax")

func isSpace(ch byte) bool {
	if string(ch) == " " {
		return true
	}
	return false
}

func isDigit(ch byte) bool {
	ch -= '0'
	if ch > 9 {
		fmt.Println(ErrSyntax, "ch > 9")
		return false
	}
	return true
}

func MyAtoi(str string) int {
	sLen := len(str)
	index := 0
	// skip leading spaces
	for index < len(str) && isSpace(str[index]) {
		index++
	}
	sign := str[index]

	// fast path for small integers that fit int type
	if intSize == 32 && (0 < sLen && sLen < 10) || intSize == 64 && (0 < sLen && sLen < 19) {
		// deal +-
		if str[index] == '-' || str[index] == '+' {
			index++
		}

		str = str[index:]
		if len(str) < 1 {
			fmt.Println(ErrSyntax, "len(str) < 1")
			return 0
		}

		sum := 0
		for _, ch := range []byte(str) {
			ch -= '0'
			if ch > 9 {
				break
			}
			sum = sum*10 + int(ch)
			fmt.Println(">>>>>>>>>>>> sum", sum, int(ch), str)
		}

		if sign == '-' {
			sum = -sum
		}
		fmt.Println(">>>>>>>>>>>> results", sum)
		if sum > 2147483647 {
			return 2147483647
		}

		if sum < -2147483648 {
			return -2147483648
		}
		return sum
	}

	if sign == '-' {
		return -2147483648
	}
	return 2147483647

}
