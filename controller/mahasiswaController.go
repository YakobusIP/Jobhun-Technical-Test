package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"jobhun-backend.com/database"
	"jobhun-backend.com/models"
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

	var mahasiswa models.Mahasiswa
	var list_mahasiswa []models.Mahasiswa
	var response models.MahasiswaResponse

	query := "SELECT m.id_mahasiswa, m.nama, m.usia, m.gender, m.tanggal_registrasi, h.id_hobi, h.nama, j.id_jurusan, j.nama FROM mahasiswa m LEFT JOIN mahasiswa_hobi mh ON m.id_mahasiswa = mh.id_mahasiswa LEFT JOIN hobi h ON mh.id_hobi = h.id_hobi LEFT JOIN mahasiswa_jurusan mj ON m.id_mahasiswa = mj.id_mahasiswa LEFT JOIN jurusan j ON mj.id_jurusan = j.id_jurusan WHERE m.id_mahasiswa=?"
	err := database.DATABASE.QueryRow(query, params["id"]).Scan(&mahasiswa.Id_mahasiswa, &mahasiswa.Nama, &mahasiswa.Usia, &mahasiswa.Gender, &mahasiswa.Tanggal_registrasi, &mahasiswa.Jurusan.Id_jurusan, &mahasiswa.Jurusan.Nama, &mahasiswa.Hobi.Id_hobi, &mahasiswa.Hobi.Nama)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No data found on id", http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	list_mahasiswa = append(list_mahasiswa, mahasiswa)

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
	var jurusanId int64
	var hobiId int64
	
	// Insert new jurusan, if jurusan already exist within the database, 
	// get the id and insert into mahasiswa_jurusan. If jurusan does not exist in the database,
	// create a new jurusan and use that id
	query1 := "INSERT INTO jurusan(nama) VALUES(?)"
	res, err := database.DATABASE.ExecContext(context.Background(), query1, request.Jurusan.Nama)

	if err != nil {
		// Found duplicate entry (violating unique trait of the attribute)
		if strings.Contains(err.Error(), "Duplicate entry") {
			queryGetJurusan := "SELECT id_jurusan FROM jurusan WHERE nama=?"
			err := database.DATABASE.QueryRow(queryGetJurusan, request.Jurusan.Nama).Scan(&jurusanId)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)	
				return
			}
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)	
			return
		}
	} else {
		// Insertion successful, fetch the id
		jurusanId, err = res.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)	
			return
		}
	}

	// Insert new hobi, if hobi already exist within the database, 
	// get the id and insert into mahasiswa_hobi. If hobi does not exist in the database,
	// create a new hobi and use that id
	query2 := "INSERT INTO hobi(nama) VALUES(?)"
	res, err = database.DATABASE.ExecContext(context.Background(), query2, request.Hobi.Nama)

	if err != nil {
		// Found duplicate entry (violating unique trait of the attribute)
		if strings.Contains(err.Error(), "Duplicate entry") {
			queryGetHobi := "SELECT id_hobi FROM hobi WHERE nama=?"
			err := database.DATABASE.QueryRow(queryGetHobi, request.Hobi.Nama).Scan(&hobiId)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)	
				return
			}
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)	
			return
		}
	} else {
		// Insertion successful, fetch the id
		hobiId, err = res.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)	
			return
		}
	}

	query3 := "INSERT INTO mahasiswa(nama, usia, gender, tanggal_registrasi) VALUES(?, ?, ?, NOW())"
	res, err = database.DATABASE.ExecContext(context.Background(), query3, request.Nama, request.Usia, request.Gender)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)	
		return
	}

	mahasiswaId, err := res.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)	
		return
	}

	query4 := "INSERT INTO mahasiswa_hobi(id_mahasiswa, id_hobi) VALUES(?, ?)"
	_, err = database.DATABASE.ExecContext(context.Background(), query4, mahasiswaId, hobiId)

	query5 := "INSERT INTO mahasiswa_jurusan(id_mahasiswa, id_jurusan) VALUES(?, ?)"
	_, err = database.DATABASE.ExecContext(context.Background(), query5, mahasiswaId, jurusanId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)	
		return
	}
	response = models.MahasiswaResponse{Message: "Successfully inserted mahasiswa to the database"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateMahasiswa(w http.ResponseWriter, r *http.Request) {
	var request models.MahasiswaRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var id_jurusan int64
	var id_hobi int64

	// Check if jurusan already exists in the database, if not then create a new jurusan and get the id.
	// If jurusan exists, then fetch the id and update mahasiswa_jurusan
	query1 := "SELECT id_jurusan FROM jurusan WHERE nama=?"
	err = database.DATABASE.QueryRow(query1, request.Jurusan.Nama).Scan(&id_jurusan)

	if err != nil {
		if err == sql.ErrNoRows {
			query := "INSERT INTO jurusan(nama) VALUES(?)"
			res, err := database.DATABASE.ExecContext(context.Background(), query, request.Jurusan.Nama)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)	
				return
			}

			id_jurusan, err = res.LastInsertId()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)	
				return
			}
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	
	queryUpdate := "UPDATE mahasiswa_jurusan SET id_jurusan=? WHERE id_mahasiswa=?"
	_, err = database.DATABASE.ExecContext(context.Background(), queryUpdate, id_jurusan, request.Id_mahasiswa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Check if hobi already exists in the database, if not then create a new hobi and get the id.
	// If hobi exists, then fetch the id and update mahasiswa_hobi
	query2 := "SELECT id_hobi FROM hobi WHERE nama=?"
	err = database.DATABASE.QueryRow(query2, request.Hobi.Nama).Scan(&id_hobi)

	if err != nil {
		if err == sql.ErrNoRows {
			query := "INSERT INTO hobi(nama) VALUES(?)"
			res, err := database.DATABASE.ExecContext(context.Background(), query, request.Hobi.Nama)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)	
				return
			}

			id_hobi, err = res.LastInsertId()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)	
				return
			}
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} 
	
	queryUpdate = "UPDATE mahasiswa_hobi SET id_hobi=? WHERE id_mahasiswa=?"
	_, err = database.DATABASE.ExecContext(context.Background(), queryUpdate, id_hobi, request.Id_mahasiswa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query := "UPDATE mahasiswa SET nama=?, usia=?, gender=? WHERE id_mahasiswa=?"
	_, err = database.DATABASE.ExecContext(context.Background(), query, request.Nama, request.Usia, request.Gender, request.Id_mahasiswa)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response models.MahasiswaResponse

	response = models.MahasiswaResponse{Message: "Successfully updated mahasiswa"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteMahasiswa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var id_mahasiswa int64
	queryCheck := "SELECT id_mahasiswa FROM mahasiswa WHERE id_mahasiswa=?"
	err := database.DATABASE.QueryRow(queryCheck, params["id"]).Scan(&id_mahasiswa)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No data found on id", http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	query1 := "DELETE FROM mahasiswa_jurusan WHERE id_mahasiswa=?"
	_, err = database.DATABASE.ExecContext(context.Background(), query1, params["id"])

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