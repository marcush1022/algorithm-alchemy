/*
Given a linked list, determine if it has a cycle in it.

Follow up:
Can you solve it without using extra space?
*/

class Solution {
public:
    bool hasCycle(ListNode *head) {
        ListNode *fast, *slow;
        fast=head, slow=head;
        while(fast && fast->next && slow)
        {
            fast=fast->next->next;
            slow=slow->next;
            if(fast==slow)
                return true;
        }
        return false;
    }
};
