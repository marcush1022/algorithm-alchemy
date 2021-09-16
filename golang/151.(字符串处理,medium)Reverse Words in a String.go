// package golang
// name 151.(字符串处理,medium)Reverse Words in a String

/*
Given an input string, reverse the string word by word.

For example,
Given s = "the sky is blue",
return "blue is sky the".

难点在处理多余的空格，left为每一个单词开始的位置，right为单词结束的位置
开始时left=right
right寻找每一个单词的末尾
找到后逆置
并加一个空格
结束时right为字符串最后一个单词的最后一个字母的位置

0. the sky is blue
1. eht yks si eulb
2. blue is sky the
*/

package golang

func reverseBytes(s []byte, left, right int) []byte {
	for i, j := left, right; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func ReverseWords(words string) string {
	original := []byte(words)
	length := len(original)
	if length == 0 {
		return words
	}

	left, right, index := 0, 0, 0
	for index < length {
		// skip the blanks util end
		for index < length && string(original[index]) == " " {
			index++
		}
		// if the string is full of blank
		if index == length {
			break
		}

		if left != right {
			right++ // skip blank
		}

		left = right
		// right looking for word's rear
		for index < length && string(original[index]) != " " {
			right++ // right is word's rear next blank
			index++
		}
		// reverse each word
		reverseBytes(original, left, right-1)
	}
	return string(reverseBytes(original, 0, length-1))
}
