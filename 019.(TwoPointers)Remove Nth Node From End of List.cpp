/***********************************************************************************/
Given a linked list, remove the nth node from the end of list and return its head.

For example,

   Given linked list: 1->2->3->4->5, and n = 2.

   After removing the second node from the end, the linked list becomes 1->2->3->5.

特殊情况：删第一个和最后一个
left指向要删的结点的前一个节点
/***********************************************************************************/

class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        if(head==NULL)
		return head;
	ListNode *left, *right;
	ListNode *p;
	left=right=head;
	int i=0;
	while(i<n)
	{
		++i;
		right=right->next;
	}
	//right为空说明要删的节点为倒数n个即第一个
        if(right==NULL)
            	return head->next;
	while(right->next!=NULL)
	{
		right=right->next;
		left=left->next;
	}
	left->next=left->next->next;
	return head;
    }
};












