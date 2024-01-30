
 type ListNode struct {
     Val int
     Next *ListNode
 }
 
 func pairSum(head *ListNode) int {
    slow := head
    fast := head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    finReversed := slow
    var prev *ListNode
    for finReversed != nil {
        next := finReversed.Next
        finReversed.Next = prev
        prev = finReversed
        finReversed = next
    }
    finReversed = prev

        first := head
    
    maxTwinSum := first.Val + finReversed.Val
    first = first.Next
    finReversed = finReversed.Next
    for first != nil && finReversed != nil {
        var curSum int 
        curSum = first.Val + finReversed.Val
        if curSum > maxTwinSum {
            maxTwinSum = curSum
        }
        first = first.Next
        finReversed = finReversed.Next
    }

    return maxTwinSum
}
