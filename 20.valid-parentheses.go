package leetcode

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
    if len(s) == 1 {
        return false
    }
    stack := []rune{}
    for _, w := range s {
        switch w {
        case '{', '[', '(':
            stack = append(stack, w)
        case '}':
            if len(stack) <= 0 || stack[len(stack)-1] != '{' {
                return false
            } else {
                stack = stack[:len(stack)-1]
            }
        case ')':
            if len(stack) <= 0 || stack[len(stack)-1] != '(' {
                return false
            } else {
                stack = stack[:len(stack)-1]
            }
        case ']':
            if len(stack) <= 0 || stack[len(stack)-1] != '[' {
                return false
            } else {
                stack = stack[:len(stack)-1]
            }
        default:
            return false
        }
    }
    if len(stack) > 0 {
        return false
    }
    return true
}

// @lc code=end
