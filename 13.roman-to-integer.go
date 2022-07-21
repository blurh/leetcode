package leetcode

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start

import (
    "bufio"
    "strings"
)

func romanToInt(s string) int {
    var m = make(map[string]int)
    m["I"] = 1
    m["V"] = 5
    m["X"] = 10
    m["L"] = 50
    m["C"] = 100
    m["D"] = 500
    m["M"] = 1000

    r := strings.NewReader(s)
    b := bufio.NewReaderSize(r, len(s))
    var res int

    for i := 0; i <= len(s)-1; i++ {
        pos, _ := b.ReadByte()
        peek, _ := b.Peek(1)
        if m[string(peek)] <= m[string(pos)] {
            res += m[string(pos)]
        } else {
            res -= m[string(pos)]
        }
    }

    return res
}

// @lc code=end
