class MyQueue {
public:
    	stack<int> stk1;
	stack<int> stk2;
	
	MyQueue();
	
	void reverseStk()
	{
		while(stk1.top())
		{
			stk2.push(stk1.top());
			stk1.pop();
		}
	}
    
    	// Push element x to the back of queue.
    	void push(int x) {
        	stk1.push(x);
    	}
    
    	// Removes the element from in front of queue.
    	void pop() {
        	reverseStk();
		stk2.pop();
    	}

    	// Get the front element.
    	int peek() {
        	reverseStk();
		return stk2.top();
    	}

    	// Return whether the queue is empty.
    	bool empty() {
        	reverseStk();
        	return (stk2.size()==0);
    	}
};
