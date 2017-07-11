/******************************************************************************************/
Merge two sorted linked lists and return it as a new list. The new list should be made by 
splicing together the nodes of the first two lists.
/******************************************************************************************/

class Solution {
public:
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        ListNode *head;
		ListNode *p=head;
        if(l1==NULL)
            return l2;
        if(l2==NULL)
            return l1;
		if(l1->val<l2->val)
		{
			head=p=l1;
			l1=l1->next;
		}
		else
		{
			head=p=l2;
			l2=l2->next;
		}
		while(l1!=NULL && l2!=NULL)
		{
			if(l1->val<l2->val)
			{
				p->next=l1;
				p=l1;
				l1=l1->next;
			}
			else
			{
				p->next=l2;
				p=l2;
				l2=l2->next;
			}
		}
		if(l1)
			p->next=l1;
		if(l2)
			p->next=l2;
		return head;
    }
};
