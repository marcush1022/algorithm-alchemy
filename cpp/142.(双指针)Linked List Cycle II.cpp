/*
Given a linked list, return the node where the cycle begins. If there is 
no cycle, return null.

Note: Do not modify the linked list.

Follow up:
Can you solve it without using extra space?

count记录链表环中的节点数量-1, p先移动count个位置，然后移动p和pre直到p->next=pre
*/

class Solution {
public:
    ListNode *meetNode(ListNode *head)
    {
        ListNode *fast, *slow;
        fast=head, slow=head;
        while(fast && fast->next && slow)
        {
            fast=fast->next->next;
            slow=slow->next;
            if(fast==slow)
                return fast;
        }
        return NULL;
    }
    
    ListNode *detectCycle(ListNode *head) {
        ListNode *p=meetNode(head);
        if(!p)
            return NULL;
        ListNode *q=p;
        int count=0;
        while(q->next!=p)
        {
            q=q->next;
            ++count;
        }
        ListNode *pre=head;
        p=head;
        while(count--)
        {
            p=p->next; 
        }
        while(p->next!=pre)
        {
            pre=pre->next;
            p=p->next;
        }
        return pre;
    }
};
