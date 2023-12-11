package factorial

const (
	maxSliceSize = 50000
)

// Recursive ищем факториал с помощью рекурсии.
func Recursive(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * Recursive(n-1)
	}
}

// Dynamic ищем факториал с помощью динамического программирования.
func Dynamic(n int) int {
	memo := make([]int, n+1)
	memo[0] = 1
	for i := 1; i <= n; i++ {
		memo[i] = i * memo[i-1]
	}
	return memo[n]
}

// Calculate функция с расчетом.
func Calculate() [maxSliceSize]int {
	slice := [maxSliceSize]int{}
	for i := 0; i < maxSliceSize; i++ {
		slice[i] = Dynamic(i)
	}
	return slice
}
