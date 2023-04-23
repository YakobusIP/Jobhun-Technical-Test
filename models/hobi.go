package models

type Hobi struct {
	Id_hobi int    `json:"id_hobi"`
	Nama    string `json:"nama"`
}

type HobiRequest struct {
	Nama string `json:"nama"`
}