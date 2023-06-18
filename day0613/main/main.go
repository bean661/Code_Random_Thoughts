package main

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		tmp := make([]int, len(word2)+1)
		dp[i] = tmp
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func findMinimum(numbers ...int) int {
	if len(numbers) == 0 {
		panic("At least one number must be provided")
	}

	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}

	return min
}

func minDistanceII(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		tmp := make([]int, len(word2)+1)
		dp[i] = tmp
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = findMinimum(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
