package main

import (
	"fmt"
	"net/http"
)

func profil(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Halaman Profil")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Server UDB")
	})

	http.HandleFunc("/profil", profil)

	fmt.Println("Server running di: http://localhost:8182")
	http.ListenAndServe(":8182", nil)
}
