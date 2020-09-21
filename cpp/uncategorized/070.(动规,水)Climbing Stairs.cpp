/*
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

注意: 直接递归会超时
*/

class Solution {
public:
    int climbStairs(int n) {
        if(n<=2)
            return n;
        int ppre=1, pre=2;
        int p;
        for(int i=3; i<=n; i++)
        {
            p=ppre+pre;
            ppre=pre;
            pre=p;
        }
        return p;
    }
};
