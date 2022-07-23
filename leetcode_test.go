package leetcode

import (
    "reflect"
    "testing"
)

func RunBenchs(count int, f func()) {
    for n := 0; n < count; n++ {
        f()
    }
}

func BenchmarkLeetCode(b *testing.B) {
    b.Run("test for two-sum", func(b *testing.B) {
        nums := []int{2, 7, 11, 15}
        RunBenchs(b.N, func() {
            got := twoSum(nums, 9)
            want := []int{0, 1}
            if !reflect.DeepEqual(got, want) {
                b.Error("two-sum test fail")
            }
        })
    })
    b.Run("test for palindrome-number", func(b *testing.B) {
        RunBenchs(b.N, func() {
            gotTrue := isPalindrome(121)
            gotFalse := isPalindrome(10)
            gotMinusFalse := isPalindrome(-121)
            if !gotTrue || gotFalse || gotMinusFalse {
                b.Error("palindrome-number test fail")
            }
        })
    })
    b.Run("test for roman-to-integer", func(b *testing.B) {
        assert := func(got, want int) {
            if got != want {
                b.Error("roman-to-integer test fail")
            }
        }
        RunBenchs(b.N, func() {
            assert(romanToInt("III"), 3)
            assert(romanToInt("LVIII"), 58)
            assert(romanToInt("MCMXCIV"), 1994)
        })
    })
    b.Run("test for longest-common-prefix", func(b *testing.B) {
        RunBenchs(b.N, func() {
            strs := []string{"flower", "flow", "flight"}
            got := longestCommonPrefix(strs)
            want := "fl"
            if want != got {
                b.Error("longest-common-prefix test fail")
            }
        })
    })
    b.Run("test for valid-parentheses", func(b *testing.B) {
        RunBenchs(b.N, func() {
            s := "([{}])"
            if !isValid(s) {
                b.Error("valid-parentheses test fail")
            }
            s2 := "(([{}])"
            if isValid(s2) {
                b.Error("valid-parentheses test fail")
            }
        })
    })
    b.Run("test for merge-two-sorted-lists", func(b *testing.B) {
        RunBenchs(b.N, func() {
            var list1 = &ListNode{
                Val: 1,
                Next: &ListNode{
                    Val: 2,
                    Next: &ListNode{
                        Val: 4,
                    },
                },
            }
            var list2 = &ListNode{
                Val: 1,
                Next: &ListNode{
                    Val: 3,
                    Next: &ListNode{
                        Val: 4,
                    },
                },
            }
            list := mergeTwoLists(list1, list2)
            checkList := []int{}
            for list != nil {
                checkList = append(checkList, list.Val)
                list = list.Next
            }
            want := []int{1, 1, 2, 3, 4, 4}
            if !reflect.DeepEqual(checkList, want) {
                b.Error("test merge-two-sorted-lists fail")
            }
        })
    })
    b.Run("test for remove-duplicates-from-sorted-array", func(b *testing.B) {
        RunBenchs(b.N, func() {
            assertArray := func(count int, num1, num2 []int) {
                for i := 0; i <= count-1; i++ {
                    if num1[i] != num2[i] {
                        b.Error("test remove-duplicates-from-sorted-array fail")
                    }
                }
            }
            nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
            gotCount := removeDuplicates(nums)
            wantArray := []int{0, 1, 2, 3, 4}
            wantCount := len(wantArray)
            if gotCount != wantCount {
                b.Error("test remove-duplicates-from-sorted-array fail")
            }
            assertArray(gotCount, nums, wantArray)
        })
    })
}
