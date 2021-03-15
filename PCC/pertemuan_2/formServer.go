package main

import (
	"fmt"
	"net/http"
)

func FormData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Halaman Tidak Ditemukan", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")

	case "POST":
		var salahe = r.ParseForm()
		if salahe != nil {
			fmt.Fprintln(w, "Kesalahan:", salahe)
			return
		}

		var nim = r.FormValue("nim")
		var nama = r.FormValue("nama")
		fmt.Fprintln(w, "NIM = ", nim)
		fmt.Fprintln(w, "Nama = ", nama)

	default:
		fmt.Fprint(w, "Maaf. Method yang didukung hanya GET dan POST")
	}
}

func main() {
	http.HandleFunc("/", FormData)
	fmt.Println("Server berjalan di port 8080...")
	http.ListenAndServe(":8080", nil)
}
