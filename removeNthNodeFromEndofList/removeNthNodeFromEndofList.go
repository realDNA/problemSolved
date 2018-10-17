/*
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.
*/

package main

import "fmt"


type ListNode struct {
     Val int
     Next *ListNode
}

type Feed struct {
	length int
	start *ListNode
}

func (f *Feed) append(newListNode *ListNode) {
	if f.length == 0{
		f.start = newListNode
	} else {
		move := f.start
		move.Next = newListNode
		f.start = newListNode
	}
	f.length++
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
		
	runner_index := &ListNode{Val:0}
	op_index := runner_index
	runner_index.Next = head

    for i := 0; i < n; i++{
    	runner_index = runner_index.Next
    }

    for runner_index.Next != nil {
    	runner_index = runner_index.Next
    	op_index = op_index.Next   	
    }

    if op_index.Next == head {
    	head = head.Next
    }

    op_index.Next = op_index.Next.Next

    return head
}

func main() {
	f := &Feed{}
    n1 := ListNode{
    	Val:1,
    }
    f.append(&n1)
    
    n2 := ListNode{
    	Val:2,
    }
    f.append(&n2)

    n3 := ListNode{
    	Val:3,
    }
    f.append(&n3)

    n4 := ListNode{
    	Val:4,
    }
    f.append(&n4)

    n5 := ListNode{
    	Val:5,
    }
    f.append(&n5)

    start_node := removeNthFromEnd(&n1, 1)

   
		for start_node.Next != nil {
		fmt.Printf(" whole linked list = %v \n", start_node)
		start_node = start_node.Next
	    }
	    fmt.Printf(" whole linked list = %v \n", start_node)
    
    

}