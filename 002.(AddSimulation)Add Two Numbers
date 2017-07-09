/*************************************************************************************************/
You are given two non-empty linked lists representing two non-negative integers. 
The digits are stored in reverse order and each of their nodes contain a single digit. 
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
/*************************************************************************************************/

class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *ret=new ListNode(0);
		int sum=0;
		ListNode *cur=ret;
		while(1)
		{
			if(l1!=NULL)
			{
				sum+=l1->val;
				l1=l1->next;
			}
			if(l2!=NULL)
			{
				sum+=l2->val;
				l2=l2->next;
			}
			cur->val=sum%10;
			sum=sum/10;
			if(l1!=NULL || l2!=NULL || sum)
			{
				cur->next=new ListNode(0);
				cur=cur->next;
			}
			else
				break;
		}
		return ret;
    }
};

Update: 
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *head=new ListNode(0);
        ListNode *p=head;
        int sum=0, carry=0;
        while(1)
        {
            sum=0;
            if(l1)
            {
                sum+=l1->val;
                l1=l1->next;
            }
            if(l2)
            {
                sum+=l2->val;
                l2=l2->next;
            }
            p->val=(sum+carry)%10;
            carry=(sum+carry)/10;
            if(l1 || l2 || carry) //#########################
            {
                p->next=new ListNode(0);
                p=p->next;
            }
            else
                break;
        }
        return head;
    }
};

