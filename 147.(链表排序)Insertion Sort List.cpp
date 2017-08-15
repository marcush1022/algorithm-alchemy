//Sort a linked list using insertion sort.

class Solution {
public:
    ListNode* insertionSortList(ListNode* head) {
        if(head==NULL || head->next==NULL)
            return head;
        ListNode *pre, *p, *pNext;
        pre=head, p=head->next;
        pre->next=NULL;
        while(p)
        {
            pNext=p->next;
            pre=head;
            if(p->val<pre->val)
            {
                p->next=pre;
                head=p;
                p=pNext;
            }
            else
            {
                while(pre->next && pre->next->val<p->val)
                    pre=pre->next;
                p->next=pre->next;
                pre->next=p;
                p=pNext;
            }
        }
        return head;
    }
};
