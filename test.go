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
"container/heap"
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

type minHeap []*ListNode

func (h minHeap) Less(i, j int) bool {
    if h[i] == nil {
        return false
    }
    
    if h[j] == nil {
        return true
    }
    
    return h[i].Val < h[j].Val
}

func (h minHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h minHeap) Len() int {
    return len(h)
}

func (h *minHeap) Push(v interface{}) {
    *h = append(*h, v.(*ListNode))
}

func (h *minHeap) Pop() interface{} {
    old := *h
    l := len(old)
    r := old[l-1]
    *h = old[:l-1]
    return r
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    tmp := minHeap(lists)
    h := &tmp
    heap.Init(h)
    fmt.Println(" h is ------- ",h)
    var head, last *ListNode
    
    for h.Len() > 0 {
    	fmt.Println(" h length ",h.Len())
        v := heap.Pop(h).(*ListNode)
        fmt.Println("v is :", v)
        fmt.Println("v next is :", v.Next)
        if v == nil {
            continue
        }
        
        if last != nil {
            last.Next = v
        }
        
        if head == nil {
            head = v
        }
        
        last = v
        if v.Next != nil {
            heap.Push(h, v.Next)
        }
    }
    return head
}

func main() {
	
	f1 := &Feed{}
	n11 := ListNode{Val:3}
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