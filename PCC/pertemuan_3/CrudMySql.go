package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

type mahasiswa struct {
	Nim    string
	Nama   string
	Progdi string
	Smt    int
}

type response struct {
	Status bool
	Pesan  string
	Data   []mahasiswa
}

func koneksi() (*sql.DB, error) {
	db, salahe := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/cloud_udb")
	if salahe != nil {
		return nil, salahe
	}

	return db, nil
}

func tampil(pesane string) response {
	db, salahe := koneksi()
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal query: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer db.Close()

	dataMhs, salahe := db.Query("SELECT * FROM mahasiswa")
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal query: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer dataMhs.Close()

	var hasil []mahasiswa

	for dataMhs.Next() {
		var mhs = mahasiswa{}
		var salahe = dataMhs.Scan(&mhs.Nim, &mhs.Nama, &mhs.Progdi, &mhs.Smt)

		if salahe != nil {
			return response{
				Status: false,
				Pesan:  "Gagal baca: " + salahe.Error(),
				Data:   []mahasiswa{},
			}
		}

		hasil = append(hasil, mhs)
	}

	salahe = dataMhs.Err()

	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Kesalahan: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}

	return response{
		Status: true,
		Pesan:  pesane,
		Data:   hasil,
	}
}

func getMhs(nim string) response {
	db, salahe := koneksi()
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal koneksi: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer db.Close()

	dataMhs, salahe := db.Query("SELECT * FROM mahasiswa WHERE nim=?", nim)
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal query: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer dataMhs.Close()

	var hasil []mahasiswa
	for dataMhs.Next() {
		var mhs = mahasiswa{}
		var salahe = dataMhs.Scan(&mhs.Nim, &mhs.Nama, &mhs.Progdi, &mhs.Smt)

		if salahe != nil {
			return response{
				Status: false,
				Pesan:  "Gagal baca: " + salahe.Error(),
				Data:   []mahasiswa{},
			}
		}

		hasil = append(hasil, mhs)
	}

	salahe = dataMhs.Err()
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Kesalahan: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}

	return response{
		Status: true,
		Pesan:  "Berhasil tampil",
		Data:   hasil,
	}
}

func tambah(nim string, nama string, progdi string, smt string) response {
	db, salahe := koneksi()
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal Koneksi: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer db.Close()

	_, salahe = db.Exec("INSERT INTO mahasiswa VALUES (?, ?, ?, ?)", nim, nama, progdi, smt)
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal query: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}

	return response{
		Status: true,
		Pesan:  "Berhasil ubah",
		Data:   []mahasiswa{},
	}
}

func ubah(nim string, nama string, progdi string, smt string) response {
	db, salahe := koneksi()
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal Koneksi: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer db.Close()

	_, salahe = db.Exec("UPDATE mahasiswa SET nama=?, progdi=?, nim=?", nim, nama, progdi, smt)
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal query: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}

	return response{
		Status: true,
		Pesan:  "Berhasil tambah",
		Data:   []mahasiswa{},
	}
}

func hapus(nim string) response {
	db, salahe := koneksi()
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal koneksi: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}
	defer db.Close()
	_, salahe = db.Exec("DELETE FROM mahasiswa WHERE nim=?", nim)
	if salahe != nil {
		return response{
			Status: false,
			Pesan:  "Gagal hapus: " + salahe.Error(),
			Data:   []mahasiswa{},
		}
	}

	return response{
		Status: true,
		Pesan:  "Berhasil hapus",
		Data:   []mahasiswa{},
	}
}

func kontroller(w http.ResponseWriter, r *http.Request) {
	var tampilHtml, salaheTampil = template.ParseFiles("template/tampil.html")
	if salaheTampil != nil {
		fmt.Println(salaheTampil.Error())
		return
	}

	var tambahHtml, salaheTambah = template.ParseFiles("template/tambah.html")
	if salaheTambah != nil {
		fmt.Println(salaheTambah.Error())
		return
	}

	var ubahHtml, salaheUbah = template.ParseFiles("template/ubah.html")
	if salaheUbah != nil {
		fmt.Println(salaheUbah.Error())
		return
	}

	var hapusHtml, salaheHapus = template.ParseFiles("template/hapus.html")
	if salaheHapus != nil {
		fmt.Println(salaheHapus.Error())
		return
	}

	switch r.Method {
	case "GET":
		aksi := r.URL.Query()["aksi"]
		if len(aksi) == 0 {
			tampilHtml.Execute(w, tampil("Berhasil tampil"))
		} else if aksi[0] == "tambah" {
			tambahHtml.Execute(w, nil)
		} else if aksi[0] == "ubah" {
			nim := r.URL.Query()["nim"]
			ubahHtml.Execute(w, getMhs(nim[0]))
		} else if aksi[0] == "hapus" {
			nim := r.URL.Query()["nim"]
			hapusHtml.Execute(w, getMhs(nim[0]))
		} else {
			tampilHtml.Execute(w, tampil("Berhasil tampil"))
		}
	case "POST":
		var salahe = r.ParseForm()
		if salahe != nil {
			fmt.Println(w, "Kesalahan: ", salahe)
			return
		}

		var nim = r.FormValue("nim")
		var nama = r.FormValue("nama")
		var progdi = r.FormValue("progdi")
		var smt = r.FormValue("smt")

		var aksi = r.URL.Path

		if aksi == "/tambah" {
			var hasil = tambah(nim, nama, progdi, smt)
			tampilHtml.Execute(w, tampil(hasil.Pesan))
		} else if aksi == "/ubah" {
			var hasil = ubah(nim, nama, progdi, smt)
			tampilHtml.Execute(w, tampil(hasil.Pesan))
		} else if aksi == "/hapus" {
			var hasil = hapus(nim)
			tampilHtml.Execute(w, tampil(hasil.Pesan))
		} else {
			tampilHtml.Execute(w, tampil("Berhasil tampil"))
		}
	default:
		fmt.Fprint(w, "Maaf, method cuma get n post")
	}
}

func main() {
	http.HandleFunc("/", kontroller)
	fmt.Println("Server running on 8080...")
	http.ListenAndServe(":8080", nil)
}
