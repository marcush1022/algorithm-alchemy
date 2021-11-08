// package golang
// name 005.(中心扩展法)Longest_Palindromic_Substring

package golang

import "fmt"

/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum
length of s is 1000.

Example:

Input: "babad"

Output: "bab"

Note: "aba" is also a valid answer.
Example:

Input: "cbbd"

Output: "bb"

Note: 采用中心扩展法，每个字符都有可能是回文字符串的中点。依次求得各字符作为中心的回文字符串，取其中最长
的那个即可。

注意到回文字符串分为奇、偶两种类型对应“aba”和“aa”，所以每个字符都要求两次回文字符。
*/

func getLongest(str string, left, right int) string {
	for left >= 0 && right < len(str) && str[left] == str[right] {
		left--
		right++
	}
	return str[left+1 : right]
}

func LongestPalindrome(str string) string {
	longest := str[0:1]
	for i := 0; i < len(str); i++ {
		// aba 情形
		tmp := getLongest(str, i, i)
		if len(tmp) > len(longest) {
			longest = tmp
		}

		// aa 情形
		tmp = getLongest(str, i, i+1)
		if len(tmp) > len(longest) {
			longest = tmp
		}
	}
	fmt.Println(">>>>>> longest", longest)
	return longest
}
