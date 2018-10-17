/*
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
*/

package main

import(
"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

type Feed struct {
	length int
	start *ListNode
}

func (f *Feed) append(newListNode *ListNode){
	if(f.length == 0){
		f.start = newListNode
	} else {
		f.start.Next = newListNode
		f.start = newListNode
	}
	f.length++
}

func printNode(head *ListNode){
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
	fmt.Println("")
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	L1 := lists[0]
	start := L1

   for _, L := range lists[1:] {
   		L1 = start
   		for L != nil {
   			if(L1 == nil){
   				L1 = L
   				start = L1
   				break
   			} else if (L.Val >= L1.Val) && ((L1.Next == nil) ||(L.Val < L1.Next.Val)) {
				P := L
				L = L.Next
				P.Next = L1.Next
				L1.Next = P
				L1 = P
   			} else if L.Val < L1.Val {
   				P := L1
   				L1 = L
   				L = L.Next
   				L1.Next = P
   				start = L1
   			} else {
   				L1 = L1.Next
   			}
   		}
   }
   return start
}

func main() {
	
	f1 := &Feed{}
	n11 := ListNode{Val:1}
	f1.append(&n11)
	n12 := ListNode{Val:4}
	f1.append(&n12)
	n13 := ListNode{Val:5}
	f1.append(&n13)

	f2 := &Feed{}
	n21 := ListNode{Val:2}
	f2.append(&n21)
	n22 := ListNode{Val:3}
	f2.append(&n22)
	n23 := ListNode{Val:4}
	f2.append(&n23)	

	f3 := &Feed{}
	n31 := ListNode{Val:3}
	f3.append(&n31)
	n32 := ListNode{Val:5}
	f3.append(&n32)
	n33 := ListNode{Val:6}
	f3.append(&n33)

	printNode(&n11)
	printNode(&n21)
	printNode(&n31)

	var input []*ListNode
	input = append(input, &n11)
	input = append(input, &n21)
	input = append(input, &n31)

	result := mergeKLists(input)
	fmt.Println("result is : ")
	printNode(result)
	//fmt.Println(input[1].Next.Val)

}