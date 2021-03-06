// if word1[i] != word2[j]
// dp[i-1][j-1] 代表去掉最后一个字符后，前面的字符都一样了，然后在此基础上 +1，所以是替换操作
// dp[i-1][j] 代表把 word1 的最后一个字符删掉，然后在此基础上 +1，所以是增加操作
// dp[i][j-1] 代表把 word2 的最后一个字符删掉，所以是删减操作
func minDistance(word1 string, word2 string) int {
    m, n := len(word1)+1, len(word2)+1
    dp := make([][]int, m, m)
    for i := range dp {
        dp[i] = make([]int, n, n)
    }
    dp[0][0] = 0
    for i := 1; i < m; i++ {
        dp[i][0] = i
    }
    for i := 1; i < n; i++ {
        dp[0][i] = i
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
            }
        }
    }

    return dp[m-1][n-1]
}

func min(x, y, z int) int {
    m := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }

    return m(m(x, y), z)
}