//Count the number of prime numbers less than a non-negative number, n.

class Solution {
public:
    int countPrimes(int n) {
        if(n<2)
            return 0;
        int limit=sqrt(n), count=0;
        vector<bool> prime(n, true);
        for(int i=2; i<n; i++)
        {
            if(prime[i])
                ++count;
            if(i>limit)
                continue;
            for(int j=i*i; j<n; j+=i)
                prime[j]=false;
        }
        return count;
    } 
};
