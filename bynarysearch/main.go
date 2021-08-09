package main

import "fmt"

//BynarySearch is struct base for implements methods to bynary search
//algoritm.
type BynarySearch struct{}

//bynarySearch is responsable by search element in array.
func bynarySearch(arr []int, elem, start, end int) int {
	index := (start + end) / 2
	if index < len(arr) && arr[index] == elem {
		return index
	}
	if start != end {
		if arr[index] < elem {
			return bynarySearch(arr, elem, index+1, end)
		}
		return bynarySearch(arr, elem, start, index-1)
	}
	return -1
}

//Search is return of index the element.
func (b *BynarySearch) Search(arr []int, number int) int {
	return bynarySearch(arr, number, 0, len(arr))
}

func main() {
	arr := []int{1, 2, 5, 6, 7, 9, 12, 14}
	b := BynarySearch{}
	index := b.Search(arr, 12)
	fmt.Println(index)
}
