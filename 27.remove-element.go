package leetcode

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
    count := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == val {
            count++
        } else {
            nums[i-count] = nums[i]
        }
    }
    return len(nums) - count
}

// @lc code=end
