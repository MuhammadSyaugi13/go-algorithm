package main

import "fmt"

func BinarySearch(arr []int, target int) bool {
	low := 0
	high := len(arr) - 1
	mid := (low + high) / 2

	for low <= mid && mid <= high {

		// jika target sudah sesuai dengan nilai pada array dengan index mid, maka selesai
		if target == arr[mid] {
			return true
		}

		if target > arr[mid] {
			low = mid + 1
		}

		if target < arr[mid] {
			high = mid - 1
		}

		mid = (low + high) / 2

	}

	// data tidak ditemukan
	return false
}

func main() {

	arr := []int{1, 4, 5, 6, 7, 8, 9, 15, 20, 30}

	fmt.Println(BinarySearch(arr, 23))

}
