/*
Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of 
characters. No two characters may map to the same character but a character may map to itself.

For example,
Given "egg", "add", return true.

Given "foo", "bar", return false.

Given "paper", "title", return true.

Note:
You may assume both s and t have the same length.

有以下两种情况的不是映射
一对多:
f--b
o--a
o--r

多对一:
b--f
a--o
r--o
*/

class Solution {
public:
    bool isIsomorphic(string s, string t) {
        int len1=s.length(), len2=t.length();
        if(len1!=len2)
            return false;
        vector<char> map(300);
        vector<char> visited(300);
        for(int i=0; i<len1; i++)
        {
            if(map[s[i]] && map[s[i]]!=t[i]) //一对多
                return false;
            if(!map[s[i]] && visited[t[i]]) //多对一
                return false;
            map[s[i]]=t[i];
            visited[t[i]]=1;
        }
        return true;
    }
};
