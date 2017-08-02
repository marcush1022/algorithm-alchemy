/*
Given an index k, return the kth row of the Pascal's triangle.

For example, given k = 3,
Return [1,3,3,1].

注：ret[i]=ret[i-1]*(rowsNumber-i+1)/i
*/

class Solution {
public:
    vector<int> getRow(int n) {
        vector<int> ret(n+1);
        ret[0]=1;

        for(int i=1; i<=n; i++)
            ret[i]=(long long int)ret[i-1]*(long long int)(n-i+1)/i;
        
        return ret;
    }
};
