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
*/

class Solution {
public:
    void reverseWords(string &s) {
        int len=s.length();
        if(len==0)
            return ;
        int left=0, right=0, i=0;
        while(i<len)
        {
            while(i<len && s[i]==' ')
                i++;
            if(i==len)
                break;
            if(left!=right)
                s[right++]=' ';
            left=right;
            while(i<len && s[i]!=' ')
                s[right++]=s[i++];
            reverse(s.begin()+left, s.begin()+right);
        }
        s.resize(right);
        reverse(s.begin(), s.end());
    }
};
