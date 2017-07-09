/***********************************************************************************************************/
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may
want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:

string convert(string text, int nRows);
convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".
/***********************************************************************************************************/

class Solution {
public:
string convert(string s, int nRows) {
  string result;
  
  if(nRows == 1) {
    return s;
  }
  
  for(int i = 0; i < nRows; i++) {
    int j = i;
    bool flag = true;
    while(j < s.size()) {
      result.push_back(s[j]);
      
      if(i == 0 || i == nRows - 1) {
        j += 2 * (nRows - 1);
      }
      else {
        if(flag) {
          j += 2 * (nRows - 1 - i);
          flag = false;
        }
        else {
          j += 2 * i;
          flag = true;
        }
      }
    }
  }
  return result;
}
};
