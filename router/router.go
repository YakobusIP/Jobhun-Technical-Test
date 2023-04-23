package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"jobhun-backend.com/controller"
)

func Router() {
	fmt.Println("Initializing router...")
	router := mux.NewRouter()

	router.HandleFunc("/get-all-mahasiswa", controller.GetAllMahasiswa).Methods("GET")

	router.HandleFunc("/get-mahasiswa-on-id/{id}", controller.GetMahasiswaOnID).Methods("GET")

	router.HandleFunc("/insert-mahasiswa", controller.InsertMahasiswa).Methods("POST")

	router.HandleFunc("/update-mahasiswa", controller.UpdateMahasiswa).Methods("PUT")

	router.HandleFunc("/delete-mahasiswa/{id}", controller.DeleteMahasiswa).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8080", router))
}