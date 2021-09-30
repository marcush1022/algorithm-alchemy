// package golang
// name 022.(dfs)Generate_Parentheses

package golang

import "fmt"

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed
parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

"在当前局面下，你有若干种选择。那么尝试每一种选择。如果已经发现某种选择肯定不行（因为违反了某些限定条
件），就返回；如果某种选择试到最后发现是正确解，就将其加入解集"

至于要输出所有括号的正确组合形式，可以采用递归。用两个变量l和r记录剩余左括号和右括号的数量，当且仅当左
右括号数量都为0时，正常结束。当然还有一点限制，就是剩余的右括号数量比左括号多时才能添加右括号。
*/

func stringPopBack(str string) string {
	return str[:len(str)-1]
}

func dfsParentheses(result *[]string, path string, leftRemainNum, rightRemainNum int) {
	fmt.Println(">>>>>> left:", leftRemainNum, ", right:", rightRemainNum)
	if leftRemainNum == 0 && rightRemainNum == 0 {
		// valid path
		*result = append(*result, path)
		fmt.Println(">>>>>", result)
	}

	// search one path until there is no left remaining number
	if leftRemainNum > 0 {
		dfsParentheses(result, path+"(", leftRemainNum-1, rightRemainNum)
	}
	// search the other path until rightRemainNum > leftRemainNum
	if rightRemainNum > 0 && rightRemainNum > leftRemainNum {
		dfsParentheses(result, path+")", leftRemainNum, rightRemainNum-1)
	}
}

func dfsParentheses2(result *[]string, path string, leftRemainNum, rightRemainNum int) {
	fmt.Println(">>>>>> left:", leftRemainNum, ", right:", rightRemainNum)
	if leftRemainNum == 0 && rightRemainNum == 0 {
		// valid path
		*result = append(*result, path)
		fmt.Println(">>>>>", result)
	}

	if leftRemainNum > 0 {
		path += "("
		dfsParentheses2(result, path, leftRemainNum-1, rightRemainNum)
		path = stringPopBack(path)
	}

	if rightRemainNum > 0 && rightRemainNum > leftRemainNum {
		path += ")"
		dfsParentheses2(result, path, leftRemainNum, rightRemainNum-1)
		path = stringPopBack(path)
	}
}

func GenerateParentheses(n int) []string {
	result := make([]string, 0)
	dfsParentheses(&result, "", n, n)
	return result
}

func GenerateParentheses2(n int) []string {
	result := make([]string, 0)
	dfsParentheses2(&result, "", n, n)
	return result
}
