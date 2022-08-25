package controllers

import (
	"encoding/json"
	"net/http"
	. "wilayahIndonesia/app/config"
	. "wilayahIndonesia/app/models"
)

/**
 * Fungsi untuk menampilkan semua data provinsi
 */
func Provinsi_ListAll(w http.ResponseWriter, r *http.Request) {
	var provinsi []Provinces
	var results = make(map[string]interface{})

	query := Db().Find(&provinsi)

	// cek apakah data ditemukan atau tidak
	if query.RowsAffected == 0 {
		results["status"] = "data_tidak_ditemukan"
		results["data"] = provinsi
	} else {
		// Menyusun hasil respon
		results["status"] = "success"
		results["data"] = provinsi
	}
	
	// Menampilkan hasil respon menjadi JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func Provinsi_ById(w http.ResponseWriter, r *http.Request) {
	var results = make(map[string]interface{})

	// Mendapatkan parameter id dari URL
	id := r.URL.Query().Get("id")
	
	// Mencari data provinsi berdasarkan id
	var provinsi Provinces
	query := Db().Where("id = ?", id).First(&provinsi)

	// cek apakah data ditemukan atau tidak
	if query.RowsAffected == 0 {
		// data tidak ditemukan
		// Menyusun hasil respon
		results["status"] = "data_tidak_ditemukan"
		results["data"] = provinsi

	} else {
		// Menyusun hasil respon
		results["status"] = "success"
		results["data"] = provinsi
	}

	// Menampilkan hasil respon menjadi JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}