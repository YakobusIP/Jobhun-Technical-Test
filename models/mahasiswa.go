package models

import "time"

type Mahasiswa struct {
	Id_mahasiswa int `json:"id_mahasiswa"`
	Nama string `json:"nama"`
	Usia int `json:"usia"`
	Gender int `json:"gender"`
	Tanggal_registrasi time.Time `json:"tanggal_registrasi"`
	Jurusan Jurusan `json:"jurusan"`
	Hobi Hobi `json:"hobi"`
}

type MahasiswaRequest struct {
	Nama string `json:"nama"`
	Usia int `json:"usia"`
	Gender int `json:"gender"`
	Jurusan JurusanRequest `json:"jurusan"`
	Hobi HobiRequest `json:"hobi"`
}

type MahasiswaResponse struct {
	Data []Mahasiswa `json:"data"`
	Message string `json:"message"`
}