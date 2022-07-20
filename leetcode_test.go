package leetcode

import (
    "reflect"
    "testing"
)

func RunBenchCount(count int, f func()) {
    for n := 0; n < count; n++ {
        f()
    }
}

func BenchmarkLeetCode(b *testing.B) {
    b.Run("test for two-sum", func(b *testing.B) {
        nums := []int{2, 7, 11, 15}
        RunBenchCount(b.N, func() {
            got := twoSum(nums, 9)
            want := []int{0, 1}
            if !reflect.DeepEqual(got, want) {
                b.Error("two-sum test fail")
            }
        })
    })
    b.Run("test for palindrome-number", func(b *testing.B) {
        RunBenchCount(b.N, func() {
            gotTrue := isPalindrome(121)
            gotFalse := isPalindrome(10)
            gotMinusFalse := isPalindrome(-121)
            if !gotTrue || gotFalse || gotMinusFalse {
                b.Error("palindrome-number test fail")
            }
        })
    })
}
