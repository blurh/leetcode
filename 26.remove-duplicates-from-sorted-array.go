package leetcode

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// 统计重复的, 总的减去重复的即去重后个数
// 遍历的时候, 遍历过的元素位置可以被拿来利用, 所以可以直接将需要的元素放到前面的位置(题目要求)
// @lc code=start
func removeDuplicates(nums []int) int {
    var count int
    for i := 1; i <= len(nums)-1; i++ {
        if nums[i-1] == nums[i] {
            count++
        } else {
            nums[i-count] = nums[i]
        }
    }
    return len(nums) - count
}

// @lc code=end
