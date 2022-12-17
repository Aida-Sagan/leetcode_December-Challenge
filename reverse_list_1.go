package main

import (
    "fmt"
)

type ListNode struct {
    Val int 
    Next *ListNode
    //key 
}


/*func (L *ListNode) Insert(key interface{}) {
    list := &ListNode{
        Next: L.head,
        key:  key,
    }
    if L.head != nil {
        L.head.prev = list
    }
    L.head = list

    l := L.head
    for l.next != nil {
        l = l.next
    }
    L.tail = l
}*/

func reverseList(head *ListNode) *ListNode {
    var prevNode * ListNode = nil
    var current *ListNode = head
    var nextNode *ListNode

    for current != nil{
        nextNode = current.Next
        current.Next = prevNode
        prevNode = current
        current = nextNode
    }
    return prevNode
}

/*func main() {
    struct head *ListNode = nil
    head.Val = 1
    head.Next = 

    fmt.Println(reverseList(head))

}*/