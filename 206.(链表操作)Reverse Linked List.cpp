//Reverse a singly linked list.

class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        if(head==NULL || head->next==NULL)
            return head;
        ListNode *pre=head, *p=head->next, *pNext;
        pre->next=NULL;
        while(p)
        {
            pNext=p->next;
            p->next=pre;
            pre=p;
            p=pNext;
        }
        return pre;
    }
};

递归:

class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        if(head==NULL || head->next==NULL)
            return head;
        ListNode *p=head->next;
        ListNode *nhead=reverseList(p);
        head->next=NULL;
        p->next=head;
        return nhead;
    }
};
