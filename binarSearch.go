package main

import "fmt"

func binarSearch(a []int, b int) int {
	first := 0
	last := len(a) - 1

	for first <= last {
		mid := (first + last) / 2
		if a[mid] == b {
			return mid
		} else if a[mid] < b {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}

	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := 5
	index := binarSearch(a, b)

	if index != -1 {
		fmt.Printf("Элемент найден на индексе: %d\n", index)
	} else {
		fmt.Printf("Элемент не найден")
	}
}
