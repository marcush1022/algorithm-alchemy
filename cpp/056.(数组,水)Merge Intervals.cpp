/*
Given a collection of intervals, merge all overlapping intervals.

For example,
Given [1,3],[2,6],[8,10],[15,18],
return [1,6],[8,10],[15,18].
*/

struct Interval {
     int start;
     int end;
     Interval() : start(0), end(0) {}
     Interval(int s, int e) : start(s), end(e) {}
 };
 
bool isBigger(Interval i1, Interval i2)
{
    return (i1.start<i2.start);
}
class Solution {
public:
    vector<Interval> merge(vector<Interval>& intervals) {
        vector<Interval> ret;
        int len=intervals.size();
        if(len==0)
            return ret;
        sort(intervals.begin(), intervals.end(), isBigger);
        ret.push_back(intervals[0]);
        for(int i=0; i<len; i++)
        {
            if(intervals[i].start<=ret[ret.size()-1].end)
                ret[ret.size()-1].end=max(ret[ret.size()-1].end, intervals[i].end);
            else
                ret.push_back(intervals[i]);
        }
        return ret;
    }
};

