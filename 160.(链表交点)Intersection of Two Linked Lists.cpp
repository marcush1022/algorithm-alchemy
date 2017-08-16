/*
Write a program to find the node at which the intersection of two singly linked lists begins.


For example, the following two linked lists:

A:          a1 → a2
                   ↘
                     c1 → c2 → c3
                   ↗            
B:     b1 → b2 → b3
begin to intersect at node c1.


Notes:

If the two linked lists have no intersection at all, return null.
The linked lists must retain their original structure after the function returns.
You may assume there are no cycles anywhere in the entire linked structure.
Your code should preferably run in O(n) time and use only O(1) memory.

两个链表先比较长短并求两者的差diff，然后长的先走diff步，然后两者同时走直到结点相遇或走到尾端
*/

class Solution {
public:
    int getLen(ListNode *head)
    {
        if(head==NULL)
            return 0;
        int len=0;
        ListNode *p=head;
        while(p)
        {
            p=p->next;
            len++;
        }
        return len;
    }
    
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        if(headA==NULL || headB==NULL)
            return NULL;
        int len1, len2, diff;
        len1=getLen(headA);
        len2=getLen(headB);
        diff=len1-len2;
        ListNode *longList=headA;
        ListNode *shortList=headB;
        if(len1<len2)
        {
            longList=headB;
            shortList=headA;
            diff=len2-len1;
        }
        for(int i=0; i<diff; i++)
            longList=longList->next;
        while(longList && shortList)
        {
            if(longList==shortList)
                return longList;
            longList=longList->next;
            shortList=shortList->next;
        }
        return NULL;
    }
};
