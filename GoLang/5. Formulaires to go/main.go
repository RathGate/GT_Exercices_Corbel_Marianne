package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// !Down here is the first method used to treat the information.
		// !As the radios and checkboxes do not appears in the list when not
		// !selected, switched to a more exhaustive way instead.
		// r.ParseForm()
		// for key, value := range r.Form {
		// 	fmt.Printf("%v: %v\n", key, value)
		// }

		// ?Exhaustive way, listing all the input names:
		var inputNames = map[string]string{"First Name": "firstName", "Last Name": "lastName", "Year of Birth": "birthYear", "Favorite Animal": "animal", "Agreed to the TOS": "tos"}
		fmt.Println("\n----------\nReturning the values of the form:\n")
		for key, value := range inputNames {
			fmt.Printf("%v: ", key)
			if r.FormValue(value) == "" {
				fmt.Println("{empty}")
			} else {
				fmt.Println(r.FormValue(value))
			}
		}
		fmt.Println("----------")
	}
	template := template.Must(template.ParseFiles("templates/index.html"))
	template.Execute(w, nil)
}

func main() {
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
