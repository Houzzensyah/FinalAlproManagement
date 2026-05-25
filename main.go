package main 
import "fmt"

func main() {
}

func findMin(arr []int) int {
	min = arr[0]
	for i = 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func findMax(arr []int) int {
	max = arr[0]
	for i = 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func selectionSort(arr []int) {
	n := len(arr)
	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func insertionSort(arr []int) {
	n = len(arr)
	for i = 1; i < n; i++ {
		key := arr[i]
		j = i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

