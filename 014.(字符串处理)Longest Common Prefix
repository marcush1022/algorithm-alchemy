/******************************************************************************************************/
Write a function to find the longest common prefix string amongst an array of strings.

二重循环，一个从上往下，一个从左往右
/******************************************************************************************************/

I. 
class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
	int rows=strs.size();
        if(rows==0)
            return "";
        string tmp=strs[0];
	for(int i=1; i<rows; i++)
	{
		int j=0;
		for(j=0; tmp[j]!='\0' && strs[i][j]!='\0'; j++)
			if(strs[i][j]!=tmp[j])
				break;
		tmp=strs[i].substr(0, j);
	}
	return tmp;
    }
};

II. 
class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        string ret="";
        int len=strs.size();
	if(strs.size()==0)
		return ret;
	if(len==1)
		return strs[0];
	int index=0;
	while(1)
	{
		int i=1;
		while(i<len && strs[i][index]==strs[i-1][index])
		{
			if(strs[i][index]=='\0' || strs[i-1][index]=='\0')
				break;
			++i;
		}
		if(i==len)
		{
			ret+=strs[i-1][index];
			++index;
		}
		else
			break;
	}
	return ret;
    }
};
