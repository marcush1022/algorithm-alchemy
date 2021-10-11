// package golang
// name 007.(大数)Reverse_Integer

package golang

import "math"

/*
Reverse digits of an integer.

Example1: x = 123, return 321
Example2: x = -123, return -321

click to show spoilers.

Note:
The input is assumed to be a 32-bit signed integer. Your function should return 0 when the reversed integer overflows.
注意越界
*/

func ReverseInteger(number int) (res int) {
	tmp := int(math.Abs(float64(number)))
	for tmp > 0 {
		res = res*10 + tmp%10
		if res > math.MaxInt32 {
			return 0
		}
		tmp = tmp / 10
	}
	if number > 0 {
		return res
	}
	return -res
}
