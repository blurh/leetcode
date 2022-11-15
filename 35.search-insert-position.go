package leetcode

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
    // 用二分法解
    left, right := 0, len(nums)-1
    var middle, middleValue int
    for right > left {
        middle = (left + right) / 2
        middleValue = nums[middle]

        if target > middleValue {
            left = middle + 1
        } else if target < middleValue {
            right = middle - 1
        } else {
            return middle
        }
    }
    if target <= nums[left] {
        return left
    }
    return left + 1
}

// @lc code=end
