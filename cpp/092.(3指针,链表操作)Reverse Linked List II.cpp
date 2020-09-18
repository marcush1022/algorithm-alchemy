/*
Reverse a linked list from position m to n. Do it in-place and in one-pass.

For example:
Given 1->2->3->4->5->NULL, m = 2 and n = 4,

return 1->4->3->2->5->NULL.

Note:
Given m, n satisfy the following condition:
1 ≤ m ≤ n ≤ length of list.

新建头结点nhead，当m=1时，返回nhead->next，因为从1开始的节点已经移除
*/

class Solution {
public:
    ListNode* reverseBetween(ListNode* head, int m, int n) {
        if(head==NULL)
            return head;
        ListNode *pre, *p, *q;
        ListNode *nhead=new ListNode(0);
        nhead->next=head;
        pre=nhead;
        for(int i=0; i<m-1; i++)
            pre=pre->next;
        p=pre->next;
        q=p->next;
        for(int i=m; i<n; i++)
        {
            p->next=q->next;
            q->next=pre->next;
            pre->next=q;
            q=p->next;
        }
        if(m==1)
            return nhead->next;
        else
            return head;
    }
};






