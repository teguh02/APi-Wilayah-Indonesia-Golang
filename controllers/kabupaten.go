package controllers

import (
	"encoding/json"
	"net/http"
	. "wilayahIndonesia/app/config"
	. "wilayahIndonesia/app/models"
)

func Kabupaten_ByProvinceId(w http.ResponseWriter, r *http.Request)  {
	var kabupaten []Regencies
	var results = make(map[string]interface{})
	id_provinsi := r.URL.Query().Get("id_provinsi")

	// Cek apakah id_provinsi ada atau tidak
	if id_provinsi == "" {
		results["status"] = "id_provinsi_tidak_boleh_kosong"
		results["data"] = []Regencies{}
	} else {
		query := Db().Where("province_id = ?", id_provinsi).Find(&kabupaten)

		// cek apakah data ditemukan atau tidak
		if query.RowsAffected == 0 {
			results["status"] = "data_tidak_ditemukan"
			results["data"] = kabupaten
		} else {
			// Menyusun hasil respon
			results["status"] = "success"
			results["data"] = kabupaten
		}
	}

	// Menampilkan hasil respon menjadi JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func Kabupaten_ById(w http.ResponseWriter, r *http.Request)  {
	var results = make(map[string]interface{})
	var kabupaten = []Regencies{}

	// Mendapatkan parameter id dari URL
	id := r.URL.Query().Get("id")

	// cek apakah id ada atau tidak
	if id == "" {
		results["status"] = "id_tidak_boleh_kosong"
		results["data"] = []Regencies{}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)

		return 
	}
	
	// Mencari data kabupaten berdasarkan id
	query := Db().Where("id = ?", id).First(&kabupaten)

	// cek apakah data ditemukan atau tidak
	if query.RowsAffected == 0 {
		// data tidak ditemukan
		// Menyusun hasil respon
		results["status"] = "data_tidak_ditemukan"
		results["data"] = kabupaten

	} else {
		// Menyusun hasil respon
		results["status"] = "success"
		results["data"] = kabupaten

	}

	// Menampilkan hasil respon menjadi JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}