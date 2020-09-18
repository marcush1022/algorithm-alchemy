/*
Write an algorithm to determine if a number is "happy".

A happy number is a number defined by the following process: Starting with any positive 
integer, replace the number by the sum of the squares of its digits, and repeat the process 
until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does 
not include 1. Those numbers for which this process ends in 1 are happy numbers.

Example: 19 is a happy number

12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
*/

class Solution {
public:
    map<int, int> mp;

    bool isHappy(int n)
    {
        int sum=0;
        if(n==1 || sum==1)
            return true;
        if(n==0)
            return false;
        mp[n]++;
        while(n)
        {
            int tmp=n%10;
            sum+=tmp*tmp;
            n/=10;
        }
        mp[sum]++;
        if(mp.find(sum)!=mp.end())
            return false;
        return isHappy(sum);
    }

};
