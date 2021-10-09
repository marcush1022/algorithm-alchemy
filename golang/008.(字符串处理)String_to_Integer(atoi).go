// package golang
// name 008.(字符串处理)String_to_Integer(atoi)

package golang

import (
	"errors"
	"fmt"
	"strconv"
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
	// fast path for small integers that fit int type
	if intSize == 32 && (0 < sLen && sLen < 10) || intSize == 64 && (0 < sLen && sLen < 19) {
		index := 0
		for index < len(str) && isSpace(str[index]) {
			index++
		}

		sign := str[index]
		if str[index] == '-' || str[index] == '+' {
			index++
		}

		str = str[index:]
		if len(str) < 1 {
			fmt.Println(ErrSyntax, "len(str) < 1")
			return 0
		}

		fmt.Println(">>>>>>>>>>>>>> 1", str)

		for i := 0; i < len(str); i++ {
			// skip leading spaces
			if isSpace(str[i]) {
				continue
			}

			if !isDigit(str[i]) {
				str = str[:i]
				break
			}
		}

		if len(str) == 0 {
			return 0
		}

		fmt.Println(">>>>>>>>>>>>>> 2", string(str))

		sum := 0
		for _, ch := range str {
			sum = sum*10 + int(ch)
		}

		if sign == '-' {
			sum = -sum
		}
		return sum
	}

	// slow path for invalid, big integers
	i64, err := strconv.ParseInt(str, 10, 10)
	if err != nil {
		fmt.Println(ErrSyntax, err)
		return 0
	}
	return int(i64)
}
