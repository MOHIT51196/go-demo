package search

func BinarySearch(arr []int, num int) int {
	var left, right = 0, len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == num {
			return mid
		} else if num < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// not found
	return -1
}
