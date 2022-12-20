package main

func sum(n int) int {
	return (n * (n + 1)) / 2
}

func sum_in_range(start, end int) int {
	return sum(end) - sum(start-1)
}
