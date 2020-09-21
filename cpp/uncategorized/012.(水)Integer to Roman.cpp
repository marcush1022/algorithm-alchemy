/*********************************************************************************************/
Given an integer, convert it to a roman numeral.
Input is guaranteed to be within the range from 1 to 3999.
/*********************************************************************************************/

class Solution {
public:
    string intToRoman(int num) {
        string str="";
        if(num<=0 || num>3999)
            return str;
	int x=num;

        char* digit[10]={"","I","II","III","IV","V","VI","VII","VIII","IX"};
	char* ten[10]={"","X","XX","XXX","XL","L","LX","LXX","LXXX","XC"};
	char* hundred[10]={"","C","CC","CCC","CD","D","DC","DCC","DCCC","CM"};
	char* thousand[4]={"","M","MM","MMM"};

	int di,te,hu,th;
	di=num%10;
	te=num%100/10;
	hu=num%1000/100;
	th=num/1000;


	str+=thousand[th];
	str+=hundred[hu];
	str+=ten[te];
	str+=digit[di];

	return str;
    }
};
