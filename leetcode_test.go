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
}
