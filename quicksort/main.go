package main

import (
	"fmt"
)

//QuickSort is struct for implementation algoritm quicksort.
type QuickSort struct{}

//partition is responsable by create pivot e move positions.
func partition(a []int, lo, hi int) int {
	p := a[hi]
	for j := lo; j < hi; j++ {
		if a[j] < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	return lo
}

//quickSort return slice sorted
func quickSort(a []int, lo, hi int) []int {
	if lo > hi {
		return nil
	}

	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
	return a
}

//Sort is responsable by sorted slice the int.
func (q *QuickSort) Sort(a []int) []int {
	arr := quickSort(a, 0, len(a)-1)
	return arr
}

func main() {
	q := QuickSort{}
	list := []int{10, 5, 3, 23, 56, 42, 2, 1, 4, 6, 5}
	fmt.Println(q.Sort(list))
}
