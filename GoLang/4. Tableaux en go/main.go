package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Generates an Int between `min` and `max` both included:
func generateIntFromRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// Returns a 2D array of size x * y, filled with random values:
func generateRandomArr(x, y int) [][]int {
	arr := make([][]int, y)
	for i := range arr {
		arr[i] = make([]int, x)
		for j := range arr[i] {
			arr[i][j] = generateIntFromRange(0, 100)
		}
	}
	return arr
}

// Struct which will carry the values to the template:
type Data struct {
	Array [][]int
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Generates table size values and table content values:
	x := generateIntFromRange(3, 15)
	y := generateIntFromRange(3, 20)
	var data = Data{Array: generateRandomArr(x, y)}

	// Loads the template with the generated values:
	template := template.Must(template.ParseFiles("templates/index.html"))
	template.Execute(w, data)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	// Handles routing:
	http.HandleFunc("/", indexHandler)

	// Launches the server:
	preferredPort := ":8080"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, nil); err != nil {
		log.Fatal(err)
	}
}
