/********************************************************************************************************************/
Determine whether an integer is a palindrome. Do this without extra space.
/********************************************************************************************************************/

class Solution {
public:
    bool isPalindrome(int x) {
        if(x>=2147483647 || x<=-2147483648)
            return false;
        if(x<0)
            return false;
        int len=1;
	while(x/len>=10) //###########################
		len*=10;
	//len/=10;
	while(x)
	{
		if(x/len!=x%10)
                	return false;
		x=x%len;
		x=x/10;
		len/=100;
	}
	return true;
    }
};
