package leetcode

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start

import "strconv"

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    if x < 10 {
        return true
    }
    xStr := strconv.Itoa(x)
    xLen := len(xStr)
    for i := 0; i <= xLen/2; i++ {
        if xStr[i] != xStr[xLen-1-i] {
            return false
        }
    }
    return true
}

// @lc code=end
