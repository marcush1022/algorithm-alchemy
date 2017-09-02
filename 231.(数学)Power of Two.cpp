//Given an integer, write a function to determine if it is a power of two.

//如果是2的次方的话则n至多有一个1
class Solution {
public:
    bool isPowerOfTwo(int n) {
        return n>0 && !(n&(n-1));
    }
};

class Solution {
public:
    bool isPowerOfTwo(int n) {
        if(n==1)
            return true;
        if(n%2 || n<=0)
            return false;
        while(n)
        {
            n>>=1;
            if(n%2==1 && n>1)
                return false;
        }
        return true;
    }
};
