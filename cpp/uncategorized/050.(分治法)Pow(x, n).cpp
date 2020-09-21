/*
Implement pow(x, n).

分治解法：
分：将n分成n/2  直到n=0时，返回1；
治：对n为偶数，返回两数相乘的结果，奇数再乘多一个x

注意当x=INT_MIN时，反过来-n会越界需要让x=INT_MAX再乘一个x
*/

class Solution
{
	public:
	double myPow(double x, int n)
	{
		double res=1;
        	if(x==1 || n==0)
        		return 1.0;
        	if(n<0)
        	{
            		if(n==INT_MIN)
            		{
                		n=INT_MAX;
                		res*=1.0/x;
                		x=1.0/x;
            		}
            		else
            		{
                		n*=-1;
                		x=1.0/x;
            		}
        	}
		while(n>0)
		{
			if(n%2)
                		res=res*x;
			n=n>>1;
			x*=x;
		}
		return res;
	}
};

递归:
class Solution {
public:
    double myPow(double x, int n) {
        if(n==0 || x==1)
            return 1.0;
        if(n<0)
        {
            if(n==INT_MIN)
                return 1.0/((myPow(x, INT_MAX))*x);
            else
                return 1.0/(myPow(x, -n));
        }
        double half=myPow(x, n>>1);
        if(n%2)
            return half*half*x;
        else
            return half*half;
    }
};
