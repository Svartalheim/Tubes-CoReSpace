package main

import (
	"fmt"
	"strings"
)

const maximalSpaces int = 100

var spaces [maximalSpaces]CoworkingSpace
var lastUsedID int = 0 // Untuk melacak ID tertinggi yang sudah digunakan

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
	Facilities   []string
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
	// Menggunakan ID yang lebih tinggi dari ID tertinggi yang pernah digunakan
	lastUsedID++
	space.ID = lastUsedID

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

	// Untuk input fasilitas (dipisahkan dengan koma)
	fmt.Print("Facilities: ")
	var facilitiesInput string
	fmt.Scanln(&facilitiesInput)

	// Pisahkan berdasarkan koma dan tambahkan fasilitas
	if facilitiesInput != "" {
		facilities := strings.Split(facilitiesInput, ",")
		for _, facility := range facilities {
			space.Facilities = append(space.Facilities, strings.TrimSpace(facility))
		}
	}

	spaces[spaceCount] = space
	fmt.Println("Co-working space successfully added!")
	spaceCount++
}

// Fungsi untuk menampilkan list coworking space yang ada
func viewAllSpaces() {
	if spaceCount == 0 {
		fmt.Println("No co-working space available.")
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

	sortChoice := readInt() // Gunakan readInt() yang sudah ada

	// Salin array baru agar tidak memodifikasi array utama
	var sortedSpaces [maximalSpaces]CoworkingSpace
	for i := 0; i < spaceCount; i++ {
		sortedSpaces[i] = spaces[i]
	}

	switch sortChoice {
	case 1:
		sortByNameAscending(&sortedSpaces, spaceCount) // Tambah &
	case 2:
		sortByNameDescending(&sortedSpaces, spaceCount) // Tambah &
	case 3:
		sortByPriceAscending(&sortedSpaces, spaceCount) // Tambah &
	case 4:
		sortByPriceDescending(&sortedSpaces, spaceCount) // Tambah &
	case 5:
		sortByRatingAscending(&sortedSpaces, spaceCount) // Tambah &
	case 6:
		sortByRatingDescending(&sortedSpaces, spaceCount) // Tambah &
	default:
		fmt.Println("Invalid choice. Displaying list without sorting.")
	}

	displaySpaces(sortedSpaces, spaceCount)
}

// Mencari coworking space berdasarkan nama
func searchSpace() {
	if spaceCount == 0 {
		fmt.Println("No co-working space available for searching.")
		return
	}

	fmt.Println("\nSearch by:")
	fmt.Println("1. Name")
	fmt.Println("2. Location")
	fmt.Println("3. Facilities")
	fmt.Print("Enter your choice: ")

	var searchChoice string
	fmt.Scanln(&searchChoice)

	var searchKey string

	if !strings.Contains(searchChoice, "1") && !strings.Contains(searchChoice, "2") && !strings.Contains(searchChoice, "3") {
		fmt.Println("Invalid choice.")
		return
	}

	if searchChoice == "1" {
		fmt.Print("Enter the name to search for: ")
		fmt.Scanln(&searchKey)
		searchSpacesByField("1", searchKey)
	} else if searchChoice == "2" {
		fmt.Print("Enter the location to search for: ")
		fmt.Scanln(&searchKey)
		searchSpacesByField("2", searchKey)
	} else if searchChoice == "3" {
		fmt.Print("Enter the facility to search for: ")
		fmt.Scanln(&searchKey)
		searchByFacility(searchKey)
	}
}

// Mengubah isi data dari coworking space berdasarkan ID
func editSpace() {
	if spaceCount == 0 {
		fmt.Println("No co-working space available for editing.")
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
	fmt.Println("5. Facilities")
	fmt.Print("Enter your choice: ")

	var editChoice int
	fmt.Scanln(&editChoice)

	switch editChoice {
	case 1:
		fmt.Print("Enter the new name: ")
		fmt.Scanln(&spaces[index].Name)
	case 2:
		fmt.Print("Enter the new location: ")
		fmt.Scanln(&spaces[index].Location)
	case 3:
		fmt.Print("Enter the new capacity: ")
		fmt.Scanln(&spaces[index].Capacity)
	case 4:
		fmt.Print("Enter the new price per hour: ")
		fmt.Scanln(&spaces[index].PricePerHour)
	case 5:
		fmt.Print("Enter the new facilities: ")
		var facilitiesInput string
		fmt.Scanln(&facilitiesInput)

		// Reset fasilitas
		spaces[index].Facilities = nil

		// Pisahkan berdasarkan koma dan tambahkan fasilitas
		if facilitiesInput != "" {
			facilities := strings.Split(facilitiesInput, ",")
			for _, facility := range facilities {
				spaces[index].Facilities = append(spaces[index].Facilities, strings.TrimSpace(facility))
			}
		}
	default:
		fmt.Println("Invalid choice.")
		return
	}

	fmt.Println("Co-working space successfully updated!")
}

// Menghapus item dari coworking space berdasarkan ID
func deleteSpace() {
	if spaceCount == 0 {
		fmt.Println("No co-working space available for deletion.")
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
		fmt.Println("Co-working space successfully deleted!")
	} else {
		fmt.Println("Deletion cancelled.")
	}
}

// Fungsi untuk mereview sebuah coworking space
func addReview() {
	if spaceCount == 0 {
		fmt.Println("No co-working space available for review.")
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
		fmt.Println("Maximum review count reached for this space!")
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
	var sum float64
	for i := 0; i < spaces[index].ReviewCount; i++ {
		sum += spaces[index].Reviews[i].Rating
	}
	spaces[index].Rating = sum / float64(spaces[index].ReviewCount)

	fmt.Println("Review successfully added!")
}

// Lihat review
func viewReviews() {
	if spaceCount == 0 {
		fmt.Println("No co-working space available.")
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

// Menampilkan detail sebuah corowking space
func displaySpace(space CoworkingSpace) {
	fmt.Printf("ID: %d\n", space.ID)
	fmt.Printf("Name: %s\n", space.Name)
	fmt.Printf("Location: %s\n", space.Location)
	fmt.Printf("Capacity: %d people\n", space.Capacity)
	fmt.Printf("Price: $%.2f per hour\n", space.PricePerHour)
	fmt.Printf("Rating: %.1f/5.0 (%d review)\n", space.Rating, space.ReviewCount)

	if len(space.Facilities) > 0 {
		fmt.Print("Facilities: ")
		for i, facility := range space.Facilities {
			fmt.Print(facility)
			if i < len(space.Facilities)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}

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
		fmt.Println("No co-working space found that matches your search.")
	}
}

// Implementasi Binary search by ID
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

// Mencari coworking space berdasarkan fasilitas
func searchByFacility(searchKey string) {
	found := false

	fmt.Println("\nCo-working space with facilities:", searchKey)
	fmt.Println("---------------------------------------")

	for i := 0; i < spaceCount; i++ {
		hasFacility := false
		facilityIndex := 0
		for facilityIndex < len(spaces[i].Facilities) && !hasFacility {
			facility := spaces[i].Facilities[facilityIndex]
			if strings.Contains(strings.ToLower(facility), strings.ToLower(searchKey)) {
				hasFacility = true
			}
			facilityIndex++
		}

		if hasFacility {
			displaySpace(spaces[i])
			found = true
		}
	}

	if !found {
		fmt.Println("No co-working space found with the specified facility.")
	}
}

// ALGORITMA PENGURUTAN
// Selection Sort untuk pengurutan berdasarkan nama (ascending)
func sortByNameAscending(spaces *[maximalSpaces]CoworkingSpace, count int) {
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

// Selection Sort untuk pengurutan berdasarkan nama (descending)
func sortByNameDescending(spaces *[maximalSpaces]CoworkingSpace, count int) {
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

// Insertion Sort untuk pengurutan berdasarkan harga (ascending)
func sortByPriceAscending(spaces *[maximalSpaces]CoworkingSpace, count int) {
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

// Insertion Sort untuk pengurutan berdasarkan harga (descending)
func sortByPriceDescending(spaces *[maximalSpaces]CoworkingSpace, count int) {
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

// Selection Sort untuk pengurutan berdasarkan rating (ascending)
func sortByRatingAscending(spaces *[maximalSpaces]CoworkingSpace, count int) {
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

// Insertion Sort untuk pengurutan berdasarkan rating (descending)
func sortByRatingDescending(spaces *[maximalSpaces]CoworkingSpace, count int) {
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

func readInt() int {
	var value int
	fmt.Scanln(&value)
	return value
}

// Untuk membaca seluruh baris termasuk spasi
func readLineWithSpaces() string {
	var input string
	fmt.Scanf("%[^\n]", &input)
	return input
}
