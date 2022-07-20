package leetcode

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start

func twoSum(nums []int, target int) []int {
    if len(nums) <= 1 {
        return []int{}
    }
    m := make(map[int]int)
    for i, num := range nums {
        // j 是历史存入 map 的 i 索引值, 比当前 i 要小
        if j, ok := m[target-num]; ok {
            return []int{j, i}
        }
        m[num] = i
    }
    return []int{}
}

// @lc code=end
