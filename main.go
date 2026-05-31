package main 
import "fmt"

const MAX_MOVIES = 100

var movieCollection [MAX_MOVIES]Movie
var totalMovies int

type Movie struct {
	Title       string
	Year        int
	Description string
	Rating      float64
	Genre       string
}

func main() {
	totalMovies = 0

	fmt.Println("==========================================")
	fmt.Println("        Welcome to CineReview!        ")
	fmt.Println("==========================================")

	choice := -1
	for choice != 0 {
		fmt.Println("\n==========================================")
		fmt.Printf("  Current collection: %d movies\n", totalMovies)
		fmt.Println("==========================================")
		fmt.Println("  1. View all movies")
		fmt.Println("  2. Add movie")
		fmt.Println("  3. Edit movie")
		fmt.Println("  4. Delete movie")
		fmt.Println("  5. Search movie")
		fmt.Println("  6. Collection statistics")
		fmt.Println("  0. Exit")
		fmt.Println("==========================================")
		fmt.Print("  Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			showAllMovies()
			if totalMovies > 0 {
				var sortChoice int
				fmt.Println("\n  Do you want to sort?")
				fmt.Println("  1. Yes")
				fmt.Println("  2. No")
				fmt.Print("  Choose: ")
				fmt.Scan(&sortChoice)
				if sortChoice == 1 {
					sortMenu()
				}
			}
		} else if choice == 2 {
			addMovie()
		} else if choice == 3 {
			editMovie()
		} else if choice == 4 {
			deleteMovie()
		} else if choice == 5 {
			searchMenu()
		} else if choice == 6 {
			statisticsMenu()
		} else if choice != 0 {
			fmt.Println("  Invalid choice, please try again.")
		}
	}

	fmt.Println("\n  Thank you for using CineReview.")
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

