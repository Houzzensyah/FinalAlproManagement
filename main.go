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

// Find min/max function
func findMin(arr []Movie, i, n int) int {
    minIdx := i
    for j := i + 1; j < n; j++ {
        if arr[j].Genre < arr[minIdx].Genre {
            minIdx = j
        }
    }
    return minIdx
}

// func findMax(arr []int) int {
func findMax(arr []Movie, i, n int) int {
    maxIdx := i
    for j := i + 1; j < n; j++ {
        if arr[j].Genre > arr[maxIdx].Genre {
            maxIdx = j
        }
    }
    return maxIdx
}

// // Base sorting functions
func selectionSortGenre(arr []Movie, n int, asc bool) {
	for i := 0; i < n-1; i++ {
		var targetIdx int
		if asc {
			targetIdx = findMin(arr, i, n)
		} else {
			targetIdx = findMax(arr, i, n)
		}
		arr[i], arr[targetIdx] = arr[targetIdx], arr[i]
	}
}

func insertionSortJudul(arr []Movie, n int, asc bool) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		
		for j >= 0 {
			if asc {
				if arr[j].Title > key.Title {
					arr[j+1] = arr[j]
					j--
				} else {
					break
				}
			} else {
				if arr[j].Title < key.Title {
					arr[j+1] = arr[j]
					j--
				} else {
					break
				}
			}
		}
		arr[j+1] = key
	}
}



// CRUD Function
func showMovieDetails(movie Movie) {
	fmt.Println()
	fmt.Printf(" Title: %s\n", movie.Title)
	fmt.Printf(" Year: %d\n", movie.Year)
	fmt.Printf(" Description: %s\n", movie.Description)
	fmt.Printf(" Rating: %.1f\n", movie.Rating)
	fmt.Printf(" Genre: %s\n", movie.Genre)
}


func showAllMovies() {
	var i int
	i = 0
	if totalMovies == 0 {
		fmt.Println(" No movies in the collection.")
	} else {
		for i < totalMovies {
			showMovieDetails(movieCollection[i])
			i = i + 1
		}
	}
}

func addMovie() {
	if totalMovies >= MAX_MOVIES {
		fmt.Println(" Movie collection is full. Cannot add more movies.")
	}else{
		fmt.Print(" Enter movie title: ")	
		fmt.Scan(&movieCollection[totalMovies].Title)
		fmt.Print(" Enter release year: ")
		fmt.Scan(&movieCollection[totalMovies].Year)
		fmt.Print(" Enter description: ")
		fmt.Scan(&movieCollection[totalMovies].Description)
		fmt.Print(" Enter rating (0.0 - 10.0): ")
		fmt.Scan(&movieCollection[totalMovies].Rating)
		fmt.Print(" Enter genre: ")
		fmt.Scan(&movieCollection[totalMovies].Genre)
		totalMovies = totalMovies + 1
		fmt.Println(" Movie added successfully!")
	}



}
func editMovie() {
	if totalMovies == 0 {
		fmt.Printf(" No movies to edit.\n")
	}else {
		showAllMovies()
		var index int
		fmt.Print(" Enter the number of the movie to edit (1 - ", totalMovies, "): ")
		fmt.Scan(&index)
		if index < 1 || index > totalMovies {
			fmt.Println(" Invalid movie number.")
		} else {
			index = index - 1 // Make -1 karena base index di golang di mulai dari 0
			fmt.Print(" Enter new title (current: ", movieCollection[index].Title, "): ")
			fmt.Scan(&movieCollection[index].Title)
			fmt.Print(" Enter new release year (current: ", movieCollection[index].Year, "): ")
			fmt.Scan(&movieCollection[index].Year)
			fmt.Print(" Enter new description (current: ", movieCollection[index].Description, "): ")
			fmt.Scan(&movieCollection[index].Description)
			fmt.Print(" Enter new rating (current: ", movieCollection[index].Rating, "): ")
			fmt.Scan(&movieCollection[index].Rating)
			fmt.Print(" Enter new genre (current: ", movieCollection[index].Genre, "): ")
			fmt.Scan(&movieCollection[index].Genre)
			fmt.Println(" Movie updated successfully!")
	}
	}
}

func deleteMovie() {
	if totalMovies == 0 {
		fmt.Printf(" No movies to delete.\n")
	}else {
		showAllMovies()
		var index int
		fmt.Print(" Enter the number of the movie to delete (1 - ", totalMovies, "): ")
		fmt.Scan(&index)
		if index < 1 || index > totalMovies {
			fmt.Println(" Invalid movie number.")
		}else {	
			i := index
			for i < totalMovies-1{
				movieCollection[i] = movieCollection[i+1]
				i = i + 1
			}
			totalMovies = totalMovies - 1
			fmt.Println(" Movie deleted successfully!")
		}
	}
}

// Program functions

func searchMenu() {

}

func statisticsMenu() {

}

func sortMenu() {

}			





// Main function
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






