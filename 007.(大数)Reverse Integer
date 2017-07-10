/*****************************************************************************************************************************/
Reverse digits of an integer.

Example1: x = 123, return 321
Example2: x = -123, return -321

click to show spoilers.

Note:
The input is assumed to be a 32-bit signed integer. Your function should return 0 when the reversed integer overflows.
注意越界
/*****************************************************************************************************************************/

class Solution {
public:
    int reverse(int x) {
        long long tmp = abs((long long)x);
        long long ret = 0;
        while (tmp) {
            ret = ret * 10 + tmp % 10;
            if (ret > INT_MAX)
                return 0;
            tmp /= 10;
        }
        if (x > 0)
            return (int)ret;
        else
            return (int)-ret;
    }
};
