package leetcode

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return strs[0]
    }
    longestPre := strs[0]
    for i := 1; i <= len(strs)-1; i++ {
        for j := 0; j <= len(longestPre)-1; j++ {
            if len(strs[i]) <= j || longestPre[j] != strs[i][j] {
                longestPre = longestPre[:j]
                break
            }
            if longestPre == "" {
                return ""
            }
        }
    }
    return longestPre
}

// @lc code=end
