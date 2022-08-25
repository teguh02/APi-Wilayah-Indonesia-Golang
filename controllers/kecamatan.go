package controllers

import (
	"encoding/json"
	"net/http"
	. "wilayahIndonesia/app/config"
	. "wilayahIndonesia/app/models"
)

func Kecamatan_By_Regency_id(w http.ResponseWriter, r *http.Request) {
	var results = make(map[string]interface{})

	// Mendapatkan parameter id dari URL
	id := r.URL.Query().Get("id_kabupaten")

	// cek apakah id ada atau tidak
	if id == "" {
		results["status"] = "id_kabupaten_tidak_boleh_kosong"
		results["data"] = []Districts{}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)

		return
	}

	// Mencari data kecamatan berdasarkan id
	var kecamatan []Districts
	query := Db().Where("regency_id = ?", id).Find(&kecamatan)

	// cek apakah data ditemukan atau tidak
	if query.RowsAffected == 0 {
		// data tidak ditemukan
		// Menyusun hasil respon
		results["status"] = "data_tidak_ditemukan"
		results["data"] = kecamatan

	} else {
		// Menyusun hasil respon
		results["status"] = "success"
		results["data"] = kecamatan
	}

	// Menampilkan hasil respon menjadi JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func Kecamatan_By_id(w http.ResponseWriter, r *http.Request) {
	var results = make(map[string]interface{})
	var kecamatan = []Districts{}

	// Mendapatkan parameter id dari URL
	id := r.URL.Query().Get("id")

	// cek apakah id ada atau tidak
	if id == "" {
		results["status"] = "id_tidak_boleh_kosong"
		results["data"] = []Districts{}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)

		return
	}
	
	// Mencari data kecamatan berdasarkan id
	query := Db().Where("id = ?", id).First(&kecamatan)

	// cek apakah data ditemukan atau tidak
	if query.RowsAffected == 0 {
		// data tidak ditemukan
		// Menyusun hasil respon
		results["status"] = "data_tidak_ditemukan"
		results["data"] = kecamatan

	} else {
		// Menyusun hasil respon
		results["status"] = "success"
		results["data"] = kecamatan
	}

	// Menampilkan hasil respon menjadi JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}