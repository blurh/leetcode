package leetcode

type ListNode struct {
    Val  int
    Next *ListNode
}

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    next := &ListNode{}
    head := next
    for list1 != nil || list2 != nil {
        if list1 == nil {
            next.Next = list2
            break
        } else if list2 == nil {
            next.Next = list1
            break
        }
        if list1.Val < list2.Val {
            next.Next = list1
            list1 = list1.Next
        } else {
            next.Next = list2
            list2 = list2.Next
        }
        next = next.Next
    }
    return head.Next
}

// @lc code=end
