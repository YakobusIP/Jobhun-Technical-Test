package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"jobhun-backend.com/database"
	"jobhun-backend.com/models"
	"jobhun-backend.com/utilities"
)

func GetAllMahasiswa(w http.ResponseWriter, r *http.Request) {
	query := "SELECT m.id_mahasiswa, m.nama, m.usia, m.gender, m.tanggal_registrasi, h.id_hobi, h.nama, j.id_jurusan, j.nama FROM mahasiswa m LEFT JOIN mahasiswa_hobi mh ON m.id_mahasiswa = mh.id_mahasiswa LEFT JOIN hobi h ON mh.id_hobi = h.id_hobi LEFT JOIN mahasiswa_jurusan mj ON m.id_mahasiswa = mj.id_mahasiswa LEFT JOIN jurusan j ON mj.id_jurusan = j.id_jurusan"
	rows, err := database.DATABASE.Query(query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var list_mahasiswa []models.Mahasiswa
	var response models.MahasiswaResponse

	for rows.Next() {
		var mahasiswa models.Mahasiswa
		err = rows.Scan(&mahasiswa.Id_mahasiswa, &mahasiswa.Nama, &mahasiswa.Usia, &mahasiswa.Gender, &mahasiswa.Tanggal_registrasi, &mahasiswa.Jurusan.Id_jurusan, &mahasiswa.Jurusan.Nama, &mahasiswa.Hobi.Id_hobi, &mahasiswa.Hobi.Nama)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		list_mahasiswa = append(list_mahasiswa, mahasiswa)
	}

	response = models.MahasiswaResponse{Data: list_mahasiswa}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetMahasiswaOnID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	query := "SELECT m.id_mahasiswa, m.nama, m.usia, m.gender, m.tanggal_registrasi, h.id_hobi, h.nama, j.id_jurusan, j.nama FROM mahasiswa m LEFT JOIN mahasiswa_hobi mh ON m.id_mahasiswa = mh.id_mahasiswa LEFT JOIN hobi h ON mh.id_hobi = h.id_hobi LEFT JOIN mahasiswa_jurusan mj ON m.id_mahasiswa = mj.id_mahasiswa LEFT JOIN jurusan j ON mj.id_jurusan = j.id_jurusan WHERE m.id_mahasiswa=?"
	rows, err := database.DATABASE.Query(query, params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var list_mahasiswa []models.Mahasiswa
	var response models.MahasiswaResponse

	for rows.Next() {
		var mahasiswa models.Mahasiswa
		err = rows.Scan(&mahasiswa.Id_mahasiswa, &mahasiswa.Nama, &mahasiswa.Usia, &mahasiswa.Gender, &mahasiswa.Tanggal_registrasi, &mahasiswa.Jurusan.Id_jurusan, &mahasiswa.Jurusan.Nama, &mahasiswa.Hobi.Id_hobi, &mahasiswa.Hobi.Nama)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		list_mahasiswa = append(list_mahasiswa, mahasiswa)
	}

	response = models.MahasiswaResponse{Data: list_mahasiswa}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertMahasiswa(w http.ResponseWriter, r *http.Request) {
	var request models.MahasiswaRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)	
		return
	}

	var response = models.MahasiswaResponse{}

	query1 := "INSERT INTO mahasiswa(nama, usia, gender, tanggal_registrasi) VALUES(?, ?, ?, NOW())"
	res, err := database.DATABASE.ExecContext(context.Background(), query1, request.Nama, request.Usia, request.Gender)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)	
		return
	}

	mahasiswaId, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve id mahasiswa", http.StatusInternalServerError)	
		return
	}

	var jurusanId int64
	var hobiId int64

	queryCheckJurusan := "SELECT * FROM jurusan WHERE nama=?"
	rows, err := database.DATABASE.Query(queryCheckJurusan, request.Jurusan.Nama)

	if (utilities.CountRows(rows) == 0) {
		query2 := "INSERT INTO jurusan(nama) VALUES(?)"
		res, err = database.DATABASE.ExecContext(context.Background(), query2, request.Jurusan.Nama)
	
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)	
			return
		}
	
		jurusanId, err = res.LastInsertId()
		if err != nil {
			http.Error(w, "Failed to retrieve id jurusan", http.StatusInternalServerError)	
			return
		}
	} else {
		queryGetJurusan := "SELECT id_jurusan FROM jurusan WHERE nama=?"
		rows, err = database.DATABASE.Query(queryGetJurusan, request.Jurusan.Nama)

		if err != nil {
			http.Error(w, "Failed to retrieve rows on jurusan", http.StatusInternalServerError)	
			return
		}

		for rows.Next() {
			err = rows.Scan(&jurusanId)

			if err != nil {
				http.Error(w, "Failed to retrieve id jurusan", http.StatusInternalServerError)	
				return
			}
		}
	}

	query3 := "INSERT INTO hobi(nama) VALUES(?)"
	res, err = database.DATABASE.ExecContext(context.Background(), query3, request.Hobi.Nama)

	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		hobiId, err = res.LastInsertId()
		if err != nil {
			http.Error(w, "Failed to retrieve id hobi", http.StatusInternalServerError)	
			return
		}
	}

		
	// } else {
	// 	queryGetHobi := "SELECT id_hobi FROM hobi WHERE nama=?"
	// 	rows, err = database.DATABASE.Query(queryGetHobi, request.Hobi.Nama)

	// 	if err != nil {
	// 		http.Error(w, "Failed to retrieve rows on hobi", http.StatusInternalServerError)	
	// 		return
	// 	}

	// 	for rows.Next() {
	// 		err = rows.Scan(&hobiId)

	// 		if err != nil {
	// 			http.Error(w, "Failed to retrieve id hobi", http.StatusInternalServerError)	
	// 			return
	// 		}
	// 	}
	// }

	query4 := "INSERT INTO mahasiswa_hobi(id_mahasiswa, id_hobi) VALUES(?, ?)"
	_, err = database.DATABASE.ExecContext(context.Background(), query4, mahasiswaId, hobiId)

	query5 := "INSERT INTO mahasiswa_jurusan(id_mahasiswa, id_jurusan) VALUES(?, ?)"
	_, err = database.DATABASE.ExecContext(context.Background(), query5, mahasiswaId, jurusanId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)	
		return
	}
	response = models.MahasiswaResponse{Message: "Successfully insert mahasiswa to the database"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateMahasiswa(w http.ResponseWriter, r *http.Request) {
	
}

func DeleteMahasiswa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	query1 := "DELETE FROM mahasiswa_jurusan WHERE id_mahasiswa=?"
	_, err := database.DATABASE.ExecContext(context.Background(), query1, params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)	
		return
	}

	query2 := "DELETE FROM mahasiswa_hobi WHERE id_mahasiswa=?"
	_, err = database.DATABASE.ExecContext(context.Background(), query2, params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)	
		return
	}

	query3 := "DELETE FROM mahasiswa WHERE id_mahasiswa=?"
	_, err = database.DATABASE.ExecContext(context.Background(), query3, params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)	
		return
	}

	var response models.MahasiswaResponse

	response = models.MahasiswaResponse{Message: "Successfully delete mahasiswa"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}