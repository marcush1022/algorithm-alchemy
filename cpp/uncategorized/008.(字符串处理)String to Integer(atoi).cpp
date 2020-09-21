/*************************************************************************/
Implement atoi to convert a string to an integer.

Hint: Carefully consider all possible input cases. If you want a challenge, 
please do not see below and ask yourself what are the possible input cases.

Notes: It is intended for this problem to be specified vaguely (ie, no given
input specs). You are responsible to gather all the input requirements up 
front.
/*************************************************************************/

class Solution {
public:
    int myAtoi(string str) {
        	int len=str.length();
		if(len==0)
			return 0;
		int i=0, sign=1;
		long long int res;
		while(isspace(str[i]))
			i++;
		if(str[i]=='+' || str[i]=='-')
		{	
			if(str[i]=='-')
				sign=0;
			i++;
		}
		str=str.substr(i);
		for(int j=0; j<str.length(); j++)
			if(!isdigit(str[j]))
			{
				str=str.substr(0, j);
				break;
			}
        	if(str.length()==0)  //########################################
            		return 0;
		if(str.length()>10)
		{
			if(sign==0)
				return -2147483648;
			else
				return 2147483647;
		}
		res=stoll(str);
		if(sign==0)
			res=0-res;
		if(res>2147483647)
			return 2147483647;
		if(res<-2147483648)
			return -2147483648;
		else
			return res;
    }
};
