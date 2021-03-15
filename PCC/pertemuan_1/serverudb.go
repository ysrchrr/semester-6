package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type halaman struct {
	Judul     string
	Deskripsi string
	Progdi    []jurusan
}

type jurusan struct {
	Jenjang string
	Nama    string
}

type datamahasiswa struct {
	Judul    string
	NIM      string
	NamaMhs  string
	Fakultas string
}

func profil(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Halaman Profil")
}

func progdi(w http.ResponseWriter, r *http.Request) {
	var tampile, salahe = template.ParseFiles("template.html")
	if salahe != nil {
		fmt.Println(salahe.Error())
		return
	}
	var data = halaman{
		Judul:     "Program studi",
		Deskripsi: "Fakultas Ilmu Komputer",
		Progdi: []jurusan{
			{Jenjang: "D3", Nama: "MI"},
			{Jenjang: "S1", Nama: "SI"},
		},
	}
	tampile.Execute(w, data)
}

func mahasiswa(w http.ResponseWriter, r *http.Request) {
	var tampile, salahe = template.ParseFiles("mahasiswa.html")
	if salahe != nil {
		fmt.Println(salahe.Error())
		return
	}
	var data = datamahasiswa{
		Judul:    "Data Diri Mahasiswa",
		NIM:      "180101143",
		NamaMhs:  "Muhammad Yasir Choiri",
		Fakultas: "S1 - Sistem Informasi",
	}
	tampile.Execute(w, data)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Server UDB")
	})

	http.HandleFunc("/profil", profil)
	http.HandleFunc("/progdi", progdi)
	http.HandleFunc("/mahasiswa", mahasiswa)

	fmt.Println("Server running di: http://localhost:8182")
	http.ListenAndServe(":8182", nil)
}
