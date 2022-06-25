package main

import "fmt"

func binarySearch(fakeArr [7]int, target int) int {
	r := -1
	low := 0
	high := len(fakeArr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := fakeArr[mid]

		if guess == target {
			r = guess
			break
		}
		if guess < target {
			low = mid + 1
			mid = (low + high) / 2
		}
		if guess > target {
			high = mid - 1
			mid = (low + high) / 2
		}
	}
	return r
}

func main() {
	arr := [7]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(binarySearch(arr, 2))
}
