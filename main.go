package main

import (
	"fmt"
	"strings"
)

const maximalSpaces int = 100

var spaces [maximalSpaces]CoworkingSpace

// untuk meng track jumlah coworking space yang ada, jika langsung didefinisikan maka akan statis tergantung maximalSpacesnya
var spaceCount int

// Struktur objek untuk setiap item dari array coworking space
type CoworkingSpace struct {
	ID           int
	Name         string
	Location     string
	Capacity     int
	PricePerHour float64
	Rating       float64
	Reviews      [50]Review
	ReviewCount  int
}

// Struktur objek pada review setiap corowking space
type Review struct {
	UserName string
	Rating   float64
	Comment  string
}

func main() {
	fmt.Println("===================================")
	fmt.Println("     CO-WORKING SPACE MANAGER     ")
	fmt.Println("===================================")

	var running bool = true
	for running {
		displayMainMenu()
		choice := getUserChoice()
		switch choice {
		case 1:
			addCoworkingSpace()
		case 2:
			viewAllSpaces()
		case 3:
			searchSpace()
		case 4:
			editSpace()
		case 5:
			deleteSpace()
		case 6:
			addReview()
		case 7:
			viewReviews()
		case 8:
			running = false
			fmt.Println("Thank you for using CO-WORKING SPACE MANAGER!")
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// Fungsi untuk menampilkan menu awal
func displayMainMenu() {
	fmt.Println("\nMain Menu:")
	fmt.Println("1. Add a co-working space")
	fmt.Println("2. View all co-working spaces")
	fmt.Println("3. Search a co-working space")
	fmt.Println("4. Edit a co-working space")
	fmt.Println("5. Delete a co-working space")
	fmt.Println("6. Add a review for a space")
	fmt.Println("7. View reviews for a space")
	fmt.Println("8. Exit")
	fmt.Print("Enter your choice: ")
}

// Fungsi untuk mendapatkan pilihan user
func getUserChoice() int {
	var choice int
	fmt.Scanln(&choice)
	return choice
}

// Fungsi untuk menambahkan item dari coworking space
func addCoworkingSpace() {
	// fmt.Println(spaceCount, spaces)
	if spaceCount >= maximalSpaces {
		fmt.Println("Maximum number of co-working spaces reached!")
		return
	}
	var space CoworkingSpace
	// TODO
	// disini masih perlu perubahan karena ID setelah di delete akan menjadi ID yang sama dengan yang sudah ada
	space.ID = spaceCount + 1

	fmt.Println("\nAdding a new co-working space:")
	fmt.Print("Name: ")
	fmt.Scanln(&space.Name)

	fmt.Print("Location: ")
	var location string
	fmt.Scanln(&location)
	space.Location = location

	fmt.Print("Capacity: ")
	fmt.Scanln(&space.Capacity)

	fmt.Print("Price per hour: ")
	fmt.Scanln(&space.PricePerHour)
	spaces[spaceCount] = space
	fmt.Println("Co-working space added successfully!")
	spaceCount++

}

// Fungsi untuk menampilkan list coworking space yang ada
func viewAllSpaces() {
	if spaceCount == 0 {
		fmt.Println("No co-working spaces available.")
		return
	}

	fmt.Println("\nView co-working spaces sorted by:")
	fmt.Println("1. Name (Ascending)")
	fmt.Println("2. Name (Descending)")
	fmt.Println("3. Price (Ascending)")
	fmt.Println("4. Price (Descending)")
	fmt.Println("5. Rating (Ascending)")
	fmt.Println("6. Rating (Descending)")
	fmt.Print("Enter your choice: ")

	var sortChoice int
	fmt.Scanln(&sortChoice)

	// Copy array baru agar tidak memodifikasi array utama
	var sortedSpaces [maximalSpaces]CoworkingSpace
	for i := 0; i < spaceCount; i++ {
		sortedSpaces[i] = spaces[i]
	}

	switch sortChoice {
	case 1:
		sortByNameAscending(sortedSpaces, spaceCount)
	case 2:
		sortByNameDescending(sortedSpaces, spaceCount)
	case 3:
		sortByPriceAscending(sortedSpaces, spaceCount)
	case 4:
		sortByPriceDescending(sortedSpaces, spaceCount)
	case 5:
		sortByRatingAscending(sortedSpaces, spaceCount)
	case 6:
		sortByRatingDescending(sortedSpaces, spaceCount)
	default:
		fmt.Println("Invalid choice. Showing unsorted list.")
	}

	displaySpaces(sortedSpaces, spaceCount)
}

// mencari coworking space berdasarkan nama
func searchSpace() {
	if spaceCount == 0 {
		fmt.Println("No co-working spaces available to search.")
		return
	}

	fmt.Println("\nSearch by:")
	fmt.Println("1. Name")
	fmt.Println("2. Location")
	fmt.Print("Enter your choice: ")

	var searchChoice string
	fmt.Scanln(&searchChoice)

	var searchKey string

	if !strings.Contains(searchChoice, "1") && !strings.Contains(searchChoice, "2") {
		fmt.Println("Invalid choice.")
		return
	}

	if searchChoice == "1" {
		fmt.Print("Enter name to search: ")
	} else if searchChoice == "2" {
		fmt.Print("Enter location to search: ")

	}

	fmt.Scanln(&searchKey)
	searchSpacesByField(searchChoice, searchKey)

}

// mengubah isi data dari coworking space berdasarkan ID
func editSpace() {
	if spaceCount == 0 {
		fmt.Println("No co-working spaces available to edit.")
		return
	}

	fmt.Print("Enter the ID of the co-working space to edit: ")
	var id int
	fmt.Scanln(&id)

	// Binary search untuk ID (ada di spesifikasi umum tugas) sekaligus lebih cepat dari sequential search
	index := binarySearchById(id)
	if index == -1 {
		fmt.Println("Co-working space not found.")
		return
	}

	fmt.Println("\nEditing co-working space:", spaces[index].Name)
	fmt.Println("What do you want to edit?")
	fmt.Println("1. Name")
	fmt.Println("2. Location")
	fmt.Println("3. Capacity")
	fmt.Println("4. Price per hour")
	fmt.Print("Enter your choice: ")

	var editChoice int
	fmt.Scanln(&editChoice)

	switch editChoice {
	case 1:
		fmt.Print("Enter new name: ")
		fmt.Scanln(&spaces[index].Name)
	case 2:
		fmt.Print("Enter new location: ")
		fmt.Scanln(&spaces[index].Location)
	case 3:
		fmt.Print("Enter new capacity: ")
		fmt.Scanln(&spaces[index].Capacity)
	case 4:
		fmt.Print("Enter new price per hour: ")
		fmt.Scanln(&spaces[index].PricePerHour)
	default:
		fmt.Println("Invalid choice.")
		return
	}

	fmt.Println("Co-working space updated successfully!")
}

// menghapus item dari coworking space berdasarkan ID
func deleteSpace() {
	if spaceCount == 0 {
		fmt.Println("No co-working spaces available to delete.")
		return
	}

	fmt.Print("Enter the ID of the co-working space to delete: ")
	var id int
	fmt.Scanln(&id)

	// Binary search untuk ID (ada di spesifikasi umum tugas) sekaligus lebih cepat dari sequential search
	index := binarySearchById(id)
	if index == -1 {
		fmt.Println("Co-working space not found.")
		return
	}

	fmt.Printf("Are you sure you want to delete '%s'? (y/n): ", spaces[index].Name)
	var confirm string
	fmt.Scanln(&confirm)

	if strings.ToLower(confirm) == "y" {
		// Indeks awal berawal dari hasil ID, kemudian digeser + 1 element agar indeks yang ter rewrite menjadi indeks + 1 (otomatis indeks sekarang sudah terubah dengan indeks+1)
		for i := index; i < spaceCount-1; i++ {
			spaces[i] = spaces[i+1]
		}
		spaceCount--
		fmt.Println("Co-working space deleted successfully!")
	} else {
		fmt.Println("Deletion cancelled.")
	}
}

// Fungsi untuk mereview sebuah coworking space
func addReview() {
	if spaceCount == 0 {
		fmt.Println("No co-working spaces available to review.")
		return
	}

	fmt.Print("Enter the ID of the co-working space to review: ")
	var id int
	fmt.Scanln(&id)

	// Binary search untuk ID (ada di spesifikasi umum tugas) sekaligus lebih cepat dari sequential search dan juga sudah sorted
	index := binarySearchById(id)
	if index == -1 {
		fmt.Println("Co-working space not found.")
		return
	}

	if spaces[index].ReviewCount >= 50 {
		fmt.Println("Maximum number of reviews reached for this space!")
		return
	}

	var review Review

	fmt.Print("Your name: ")
	fmt.Scanln(&review.UserName)

	fmt.Print("Rating (0.0-5.0): ")
	fmt.Scanln(&review.Rating)
	if review.Rating < 0 {
		review.Rating = 0
	} else if review.Rating > 5 {
		review.Rating = 5
	}

	fmt.Print("Comment: ")
	var comment string
	fmt.Scanln(&comment)
	review.Comment = comment

	// Penambahan review ke dalam corowkingSpace
	spaces[index].Reviews[spaces[index].ReviewCount] = review
	spaces[index].ReviewCount++

	// Fungsi untuk menghitung rating
	// TODO: Dibikin fungsi terpisah
	var sum float64
	for i := 0; i < spaces[index].ReviewCount; i++ {
		sum += spaces[index].Reviews[i].Rating
	}
	spaces[index].Rating = sum / float64(spaces[index].ReviewCount)

	fmt.Println("Review added successfully!")
}

// Lihat review
func viewReviews() {
	if spaceCount == 0 {
		fmt.Println("No co-working spaces available.")
		return
	}

	fmt.Print("Enter the ID of the co-working space to view reviews: ")
	var id int
	fmt.Scanln(&id)

	// Binary search sudah dijelaskan di komen serupa
	index := binarySearchById(id)
	if index == -1 {
		fmt.Println("Co-working space not found.")
		return
	}

	if spaces[index].ReviewCount == 0 {
		fmt.Println("No reviews available for this space.")
		return
	}

	fmt.Printf("\nReviews for %s (Average Rating: %.1f/5.0):\n", spaces[index].Name, spaces[index].Rating)
	fmt.Println("---------------------------------------")

	for i := 0; i < spaces[index].ReviewCount; i++ {
		review := spaces[index].Reviews[i]
		fmt.Printf("User: %s\n", review.UserName)
		fmt.Printf("Rating: %.1f/5.0\n", review.Rating)
		fmt.Printf("Comment: %s\n", review.Comment)
		fmt.Println("---------------------------------------")
	}
}

// Pengimplementasian sequential search, sesuai spesifikasi umum, dengan dinamis parameter sesuai field yang diinginkan
func searchSpacesByField(field string, query string) {
	found := false

	for i := 0; i < spaceCount; i++ {
		match := false

		switch strings.ToLower(field) {
		case "1":
			match = strings.Contains(strings.ToLower(spaces[i].Name), strings.ToLower(query))
		case "2":
			match = strings.Contains(strings.ToLower(spaces[i].Location), strings.ToLower(query))
		case "id":
			match = fmt.Sprintf("%d", spaces[i].ID) == query
		case "capacity":
			match = fmt.Sprintf("%d", spaces[i].Capacity) == query
		// Tinggal tambahin disini kalau ada field baru
		default:
			fmt.Printf("Unknown search field: %s\n", field)
			return
		}

		if match {
			if !found {
				fmt.Println("\nSearch Results:")
				fmt.Println("---------------------------------------")
				found = true
			}
			displaySpace(spaces[i])
		}
	}

	if !found {
		fmt.Println("No co-working spaces found matching your query.")
	}
}

// Classic Binary search by ID implementation
func binarySearchById(id int) int {
	left := 0
	right := spaceCount - 1

	for left <= right {
		mid := (left + right) / 2

		if spaces[mid].ID == id {
			return mid
		} else if spaces[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// Menampilkan detail sebuah corowking space
func displaySpace(space CoworkingSpace) {
	fmt.Printf("ID: %d\n", space.ID)
	fmt.Printf("Name: %s\n", space.Name)
	fmt.Printf("Location: %s\n", space.Location)
	fmt.Printf("Capacity: %d people\n", space.Capacity)
	fmt.Printf("Price: $%.2f per hour\n", space.PricePerHour)
	fmt.Printf("Rating: %.1f/5.0 (%d reviews)\n", space.Rating, space.ReviewCount)
	fmt.Println("---------------------------------------")
}

// Menampilkan array of detail seluruh corowking space
func displaySpaces(spaces [maximalSpaces]CoworkingSpace, count int) {
	fmt.Println("\nAll Co-working Spaces:")
	fmt.Println("---------------------------------------")

	for i := 0; i < count; i++ {
		displaySpace(spaces[i])
	}
}

// SORTING ALGORITHMS
// TODO: FUNGSI KEBAWAH SEMUANYA JADI SATU AJA, berdasarkan field yang ingin di sort eyyyeyeye

// Selection Sort for sorting by name (ascending)
func sortByNameAscending(spaces [maximalSpaces]CoworkingSpace, count int) {
	for i := 0; i < count-1; i++ {
		minIndex := i
		for j := i + 1; j < count; j++ {
			if spaces[j].Name < spaces[minIndex].Name {
				minIndex = j
			}
		}
		if minIndex != i {
			spaces[i], spaces[minIndex] = spaces[minIndex], spaces[i]
		}
	}
}

// Selection Sort for sorting by name (descending)
func sortByNameDescending(spaces [maximalSpaces]CoworkingSpace, count int) {
	for i := 0; i < count-1; i++ {
		maxIndex := i
		for j := i + 1; j < count; j++ {
			if spaces[j].Name > spaces[maxIndex].Name {
				maxIndex = j
			}
		}
		if maxIndex != i {
			spaces[i], spaces[maxIndex] = spaces[maxIndex], spaces[i]
		}
	}
}

// Insertion Sort for sorting by price (ascending)
func sortByPriceAscending(spaces [maximalSpaces]CoworkingSpace, count int) {
	for i := 1; i < count; i++ {
		key := spaces[i]
		j := i - 1

		for j >= 0 && spaces[j].PricePerHour > key.PricePerHour {
			spaces[j+1] = spaces[j]
			j--
		}

		spaces[j+1] = key
	}
}

// Insertion Sort for sorting by price (descending)
func sortByPriceDescending(spaces [maximalSpaces]CoworkingSpace, count int) {
	for i := 1; i < count; i++ {
		key := spaces[i]
		j := i - 1

		for j >= 0 && spaces[j].PricePerHour < key.PricePerHour {
			spaces[j+1] = spaces[j]
			j--
		}

		spaces[j+1] = key
	}
}

// Selection Sort for sorting by rating (ascending)
func sortByRatingAscending(spaces [maximalSpaces]CoworkingSpace, count int) {
	for i := 0; i < count-1; i++ {
		minIndex := i
		for j := i + 1; j < count; j++ {
			if spaces[j].Rating < spaces[minIndex].Rating {
				minIndex = j
			}
		}
		if minIndex != i {
			spaces[i], spaces[minIndex] = spaces[minIndex], spaces[i]
		}
	}
}

// Insertion Sort for sorting by rating (descending)
func sortByRatingDescending(spaces [maximalSpaces]CoworkingSpace, count int) {
	for i := 1; i < count; i++ {
		key := spaces[i]
		j := i - 1

		for j >= 0 && spaces[j].Rating < key.Rating {
			spaces[j+1] = spaces[j]
			j--
		}

		spaces[j+1] = key
	}
}
