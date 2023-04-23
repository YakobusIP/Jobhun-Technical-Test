package models

type Jurusan struct {
	Id_jurusan int    `json:"id_jurusan"`
	Nama       string `json:"nama"`
}

type JurusanRequest struct {
	Nama string `json:"nama"`
}