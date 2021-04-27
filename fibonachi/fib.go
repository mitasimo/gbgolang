package fibonachi

func FibVars(n int) int {
	if n < 2 {
		return n
	}

	pred, cur := 0, 1
	for i := 2; i <= n; i++ {
		cur, pred = pred+cur, cur
	}

	return cur
}

func FibCache(n int) int {
	if n < 2 {
		return n
	}

	cache := make([]int, n+1)
	cache[0], cache[1] = 0, 1
	for i := 2; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}

	return cache[n]
}

func FibRecur(n int) int {
	if n < 2 {
		return n
	}

	return FibRecur(n-2) + FibRecur(n-1)
}
