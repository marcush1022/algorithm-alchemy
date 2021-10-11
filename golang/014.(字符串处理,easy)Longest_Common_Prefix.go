// package golang
// name 014.(字符串处理,easy)Longest_Common_Prefix

package golang

/*

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

	Input: strs = ["flower","flow","flight"]
	Output: "fl"

Example 2:

	Input: strs = ["dog","racecar","car"]
	Output: ""
	Explanation: There is no common prefix among the input strings.


Constraints:

	1 <= strs.length <= 200
	0 <= strs[i].length <= 200
	strs[i] consists of only lower-case English letters.

*/

func LongestCommonPrefix(strings []string) string {
	if len(strings) == 0 {
		return ""
	}
	prefix := strings[0]
	// indexOfPre is index of common prefix of each word
	indexOfPre := 0
	for indexOfWord := 0; indexOfWord < len(strings); indexOfWord++ {
		for indexOfPre = 0; indexOfPre < len(prefix) && indexOfPre < len(strings[indexOfWord]); indexOfPre++ {
			if strings[indexOfWord][indexOfPre] != prefix[indexOfPre] {
				break
			}
		}
		prefix = strings[indexOfWord][:indexOfPre]
	}
	return prefix
}
